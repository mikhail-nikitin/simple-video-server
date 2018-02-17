package daemon

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func CreateSignalChannel() chan os.Signal {
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Kill, os.Interrupt, syscall.SIGTERM, syscall.Signal(30))
	return signalChannel
}

func HandleSignals(c chan os.Signal, server *http.Server, logFile *os.File) {
	for {
		killSignal := <-c
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

		case syscall.Signal(30):
			log.Info("Reopening log files")
			logFile.Close()
			logFile = StartLoggingToFile(logFile.Name())
			log.Info("Reopened log")
		}
	}
}
