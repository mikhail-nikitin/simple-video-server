package main

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/mikhail.nikitin/simple-video-server/handlers"
	"os"
	"context"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	}
	defer file.Close()

	const serverUrl = ":8000"
	srv := startServer(serverUrl)

	killChannel := getSignalChannel()
	waitForKillSignal(killChannel)
	srv.Shutdown(context.Background())
}

func startServer(serverUrl string) *http.Server {
	router := handlers.Router()
	srv := &http.Server{Addr: serverUrl, Handler: router}
	go func() {
		log.WithFields(log.Fields{"url": serverUrl}).Info("listening")
		log.Fatal(srv.ListenAndServe())
	}()
	return srv
}

func getSignalChannel() chan os.Signal {
	killChannel := make(chan os.Signal)
	signal.Notify(killChannel, os.Kill, os.Interrupt, syscall.SIGTERM)
	return killChannel
}

func waitForKillSignal(c chan os.Signal) {
	killSignal := <- c
	switch killSignal {
	case os.Interrupt:
		log.Info("Got interrupt signal")
	case syscall.SIGTERM:
		log.Info("Got termination signal")
	}
}