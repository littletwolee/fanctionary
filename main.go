package main

import (
	"context"
	"fanctionary/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/littletwolee/commons"
)

func main() {
	srv := http.Server{
		Addr:         commons.GetConfig().GetString("sys.host"),
		Handler:      getRouter().Handler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := gracefulRun(&srv); err != nil {
		commons.Console().Error(err)
	}
}

func getRouter() *routers.Router {
	mongo := commons.NewMongo(
		commons.GetConfig().GetString("mongo.host"),
		commons.GetConfig().GetString("mongo.port"),
		commons.GetConfig().GetString("mongo.database"),
		commons.GetConfig().GetString("mongo.user"),
		commons.GetConfig().GetString("mongo.pwd"),
		commons.GetConfig().GetInt("mongo.pool_limit"),
	)
	return &routers.Router{
		Mongo: mongo,
	}
}
func gracefulRun(srv *http.Server) error {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-stopChan
		srv.Shutdown(context.Background())
	}()
	commons.Console().InfoF("listening to %s", srv.Addr)
	switch err := srv.ListenAndServe(); err {
	case http.ErrServerClosed, nil:
		return nil
	default:
		return err
	}
}
