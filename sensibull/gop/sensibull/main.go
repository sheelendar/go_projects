package main

import (
	"fmt"

	"github.com/sheelendar/src/go_projects/sensibull/gop/sensibull/consts"
	"github.com/sheelendar/src/go_projects/sensibull/gop/sensibull/internal"
	"github.com/sheelendar/src/go_projects/sensibull/gop/sensibull/logger"
	"github.com/sheelendar/src/go_projects/sensibull/gop/sensibull/utils"

	"net/http"
)

func main() {
	redisCli := utils.InitRedis(consts.RedisHostAndPort, "", 0)
	internal.InitHttpClient()
	fmt.Println("Server listening on :", consts.HostAndPort)

	if err := http.ListenAndServe(consts.HostAndPort, nil); err != nil {
		fmt.Println(err)
		logger.SensibullError{Message: "Not able to start server port check on priority", ErrorCode: http.StatusInternalServerError}.Err()
	}

	defer func() {
		if redisCli != nil {
			redisCli.Close()
		}
	}()
}
