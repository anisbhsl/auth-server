package executor

import (
	"fmt"
	"github.com/anisbhsl/auth-server/api"
	"github.com/anisbhsl/auth-server/auth"
	"github.com/anisbhsl/auth-server/logger"
	"github.com/anisbhsl/auth-server/routes"
	"github.com/anisbhsl/auth-server/store"
	"github.com/anisbhsl/auth-server/utils"
	"net/http"
	"time"
)

var (
	timeout = 15 * time.Second
)

type Executor struct {
	Config *utils.AppConfig
	Api    api.Service
}

func NewExecutor(config *utils.AppConfig) *Executor {
	return &Executor{
		Config: config,
		Api:api.New(auth.New(config.SecretKey,config.PrivateKeyPath,config.PublicKeyPath),store.New()),
	}
}

func (ex *Executor) Execute() {
	//spin up server
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", ex.Config.HostAddr, ex.Config.Port),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  timeout,
		Handler:      routes.RegisterRoutes(ex.Api),
	}

	logger.Debug("Starting API Server")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
