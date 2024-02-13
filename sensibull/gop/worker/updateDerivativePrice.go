package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/sheelendar/src/go_projects/sensibull/gop/sensibull/consts"
	"github.com/sheelendar/src/go_projects/sensibull/gop/sensibull/dao"
	"github.com/sheelendar/src/go_projects/sensibull/gop/sensibull/logger"
	UnderlyingAssetHandler "github.com/sheelendar/src/go_projects/sensibull/gop/sensibull/underlyingAssetHandler"
	"github.com/sheelendar/src/go_projects/sensibull/gop/sensibull/utils"

	"github.com/gorilla/websocket"
)

var webSocketConn *websocket.Conn

// InitWebSocket init websocket connection and start subscribing channel.
func InitWebSocket() *websocket.Conn {
	var err error
	webSocketConn, _, err = websocket.DefaultDialer.Dial(consts.WebSocketServerURL, nil)
	if err != nil {
		fmt.Println("Error connecting to WebSocket server:", err)
		panic("error into socket connection please check on priority")
	}
	startSubscribingDerivativeQuote()
	return webSocketConn
}

// startSubscribingDerivativeQuote start subscribing channel async way.
func startSubscribingDerivativeQuote() {
	var wg sync.WaitGroup
	wg.Add(1)
	go refreshDerivativeCacheAndSubscribeNewItem(wg)
	wg.Add(1)
	go readQuotesFromServer(wg)
	wg.Wait()
}

// readQuotesFromServer read msg from server call update db with latest derivative price.
func readQuotesFromServer(wg sync.WaitGroup) {
	defer wg.Done()
	for {
		_, message, err := webSocketConn.ReadMessage()
		if err != nil {
			logger.SensibullError{Message: "Error reading message check on priority" + err.Error()}.Err()
			continue
		}
		logger.SensibullError{Message: "Received message from server:" + string(message)}.Info()
		psm := dao.PingServerMessage{}
		err = json.Unmarshal(message, &psm)
		if err != nil {
			logger.SensibullError{Message: "unmarshalling error:" + err.Error()}.Info()
			continue
		}
		if psm.DataType == consts.DataTypePing || psm.DataType == consts.DataTypeError {
			fmt.Println("error: ", err)
		} else {
			qsm := dao.QuoteServerMessage{}
			err = json.Unmarshal(message, &qsm)
			updateQuotePriceInDB(qsm)
		}

	}
}

// updateQuotePriceInDB update latest derivative price in db.
func updateQuotePriceInDB(qsm dao.QuoteServerMessage) {
	var tokenMap map[int]dao.SubscribedDetails
	err := utils.GetObjectFromRedis(consts.DTUTKM, tokenMap)

	if qsm.DataType != consts.DataTypeQuote || tokenMap == nil || err != nil {
		return
	}
	redisDevKey := fmt.Sprintf(consts.DerivativeKey, tokenMap[qsm.Payload.Token])
	dbuao := dao.DBUnderlyingAssetObject{}
	err = utils.GetObjectFromRedis(redisDevKey, &dbuao)
	if err != nil || dbuao.DerivativeToDataMap == nil {
		logger.SensibullError{Message: "derivative not found into db" + err.Error()}.Err()
		return
	}
	quote := dbuao.DerivativeToDataMap[qsm.Payload.Token]
	quote.Strike = qsm.Payload.Price
	dbuao.DerivativeToDataMap[qsm.Payload.Token] = quote
	err = utils.SaveObjectInRedis(redisDevKey, dbuao, time.Minute)
	if err != nil {
		logger.SensibullError{Message: "error while updating price of quote" + err.Error()}.Err()
	}
}

// refreshDerivativeCacheAndSubscribeNewItem subscribes to channels or derivatives.
func refreshDerivativeCacheAndSubscribeNewItem(wg sync.WaitGroup) {
	defer wg.Done()

	var quotes []int
	for true {
		tokenMap := UnderlyingAssetHandler.InitRefreshRequiredCacheAndGetALLTokenMap()
		isNotSubscribed := false
		for derivativeToken, v := range tokenMap {
			if !v.IsSubscribed {
				isNotSubscribed = true
				v.IsSubscribed = true
				quotes = append(quotes, derivativeToken)
			}
		}
		if isNotSubscribed {
			sq := dao.SubscribeQuote{MsgCommand: "subscribe",
				DataType: "quote",
				Quote:    quotes}
			quoteBytes, err := json.Marshal(sq)
			if err != nil {
				logger.SensibullError{Message: "error while marshaling quotes" + err.Error()}.Err()
			}
			err = webSocketConn.WriteMessage(websocket.TextMessage, quoteBytes)
			if err != nil {
				logger.SensibullError{Message: "error while subscribing to channel" + err.Error()}.Err()
			} else {
				logger.SensibullError{Message: "channel subscribed successfully"}.Info()
			}
		} else {
			logger.SensibullError{Message: "no derivative found to subscribed"}.Info()
		}
		time.Sleep(time.Minute)
	}
}

// UnSubscribingDerivativeQuote need to implement to unsubscribe derivative price.
func UnSubscribingDerivativeQuote() {

}
