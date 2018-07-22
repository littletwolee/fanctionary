package main

import (
	"context"
	"fanctionary/routers"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/littletwolee/commons"
)

var log = commons.GetLogger()

func main() {
	router := &routers.Router{}
	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%d", commons.Config.GetString("server.host"), commons.Config.GetInt("server.port")),
		Handler:      router.Handler(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := gracefulRun(&srv); err != nil {
		log.LogErr(err)
	}
}

func gracefulRun(srv *http.Server) error {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-stopChan
		srv.Shutdown(context.Background())
	}()

	log.OutMsg(fmt.Sprintf("listening to %s", srv.Addr))
	switch err := srv.ListenAndServe(); err {
	case http.ErrServerClosed, nil:
		return nil
	default:
		return err
	}
}
