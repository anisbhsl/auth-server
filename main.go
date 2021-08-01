package main

import (
	"flag"
	"github.com/anisbhsl/auth-server/executor"
	"github.com/anisbhsl/auth-server/logger"
	"github.com/anisbhsl/auth-server/utils"
	"os"
)

func init(){
	logger.NewLogger(logger.Config{
		Service:  "auth-server",
	})
}

func main(){
	host:=flag.String("host","127.0.0.1","host address")
	port:=flag.String("port","5000","port")
	apiBase:=flag.String("api-base","/api/v1","base api version")
	flag.Parse()

	utils.AppParams=&utils.AppConfig{
		HostAddr:  *host,
		Port:      *port,
		SecretKey: os.Getenv("APP_SECRET"),
		ApiBase:   *apiBase,
	}

	executor.NewExecutor(utils.AppParams).Execute()
}
