package executor

import (
	"github.com/anisbhsl/auth-server/routes"
	"github.com/anisbhsl/auth-server/utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	timeout=15*time.Second
)

type Executor struct{
	config *utils.AppConfig
}

func NewExecutor(config *utils.AppConfig) *Executor{
	return &Executor{
		config: config,
	}
}

func (e *Executor) Execute(){
	//initialize store connection

	//spin up server
	srv:=&http.Server{
		Addr:              fmt.Sprintf("%s:%s",e.config.HostAddr,e.config.Port),
		ReadTimeout:       timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
		Handler: routes.RegisterRoutes(),
	}
	log.Print("Starting API Server")
	if err:=srv.ListenAndServe();err!=nil{
		log.Fatal(err)
	}
}
