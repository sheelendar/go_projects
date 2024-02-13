package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// updateOutputStream aggregate new input data coming from input file in aggregation map(streamOutput).
func updateOutputStream(input *Input) {
	milliTime := time.Unix(input.Timestamp, int64(time.Millisecond))
	dayTime := milliTime.Format(time.DateOnly)

	userKey := getUserKey(input.UserID, dayTime)
	output, ok := streamOutput[userKey]

	if !ok {
		output = Output{UserID: input.UserID}
		userKeyOrder = append(userKeyOrder, userKey)
	}

	switch input.EventType {
	case LikeReceived:
		output.LikeReceived++
	case Comment:
		output.Comment++
		break
	case Post:
		output.Post++
	default:
		fmt.Println("no input event recored for user ", input.UserID)
	}

	output.Date = dayTime
	streamOutput[userKey] = output
}

// validateInputParseInputData validate input data read from input file.
// this function can be changed based out requirement or how we read data from file
// for simplicity we just read line by line and validate json.
func validateInputParseInputData(text string) *Input {
	s := strings.IndexAny(text, "{")
	e := strings.IndexAny(text, "}")
	if s == -1 && e == -1 {
		return nil
	}
	text = text[s : e+1]
	input := Input{}

	err := json.Unmarshal([]byte(text), &input)
	if err != nil {
		fmt.Println("error while parsing read data err: ", err)
	}
	return &input

}

// writeDataInOutputFile read data from streamOutput and writes into output.json file.
func writeDataInOutputFile(outputFileName string) {

	f, err := os.OpenFile("../"+outputFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	outputArray := make([]Output, len(userKeyOrder))
	for index, out := range userKeyOrder {
		outputArray[index] = streamOutput[out]
		index++
	}
	byteArray, err := json.MarshalIndent(outputArray, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	n, err := f.Write(byteArray)
	if err != nil {
		fmt.Println(n, err)
	}
}

// getUserKey return key for each user based on dayily basis.
func getUserKey(userID int64, date string) string {
	return fmt.Sprintf(USER_KEY, userID, date)
}
