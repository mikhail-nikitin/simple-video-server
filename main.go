package main

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/mikhail-nikitin/simple-video-server/handlers"
	"os"
	"context"
	"os/signal"
	"syscall"
)

const logFileName = "my.log"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	logFile := startLoggingToFile()
	defer logFile.Close()

	const serverUrl = ":8000"
	srv := startServer(serverUrl)

	signalChannel := getSignalChannel()
	handleSignals(signalChannel, srv, logFile)
	srv.Shutdown(context.Background())
}
func startLoggingToFile() *os.File {
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	}
	return file
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
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Kill, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR1)
	return signalChannel
}

func handleSignals(c chan os.Signal, server *http.Server, logFile *os.File) {
	for {
		killSignal := <- c
		switch killSignal {
		case os.Interrupt:
			log.Info("Got interrupt signal")
			server.Shutdown(context.Background())
			log.Info("Gracefully exited")
			return

		case syscall.SIGTERM:
			log.Info("Got termination signal")
			server.Shutdown(context.Background())
			log.Info("Gracefully exited")
			return

		case syscall.SIGUSR1:
			log.Info("Reopening log files")
			logFile.Close()
			logFile = startLoggingToFile()
			log.Info("Reopened log")
		}
	}
}