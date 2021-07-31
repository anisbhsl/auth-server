package main

import (
	"github.com/anisbhsl/auth-server/executor"
	"github.com/anisbhsl/auth-server/utils"
	"flag"
)

func main(){
	host:=flag.String("host","127.0.0.1","host address")
	port:=flag.String("port","5000","port")
	apiBase:=flag.String("api-base","/api/v1","base api version")
	flag.Parse()

	utils.AppParams=&utils.AppConfig{
		HostAddr:  *host,
		Port:      *port,
		StoreName: "",
		StoreAddr: "",
		SecretKey: "",
		ApiBase:   *apiBase,
	}

	executor.NewExecutor(utils.AppParams).Execute()
}
