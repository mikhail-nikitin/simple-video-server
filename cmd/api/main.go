package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mikhail-nikitin/simple-video-server/daemon"
	"github.com/mikhail-nikitin/simple-video-server/handlers"
	log "github.com/sirupsen/logrus"
)

const logFileName = "api.log"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	logFile := daemon.StartLoggingToFile(logFileName)
	defer logFile.Close()

	db, err := sql.Open("mysql", "video_server:Q1234@/simple_video_server")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	const serverUrl = ":8000"
	router := handlers.Router(db)
	srv := daemon.StartServer(serverUrl, router)

	signalChannel := daemon.CreateSignalChannel()
	daemon.HandleSignals(signalChannel, srv, logFile)
	srv.Shutdown(context.Background())
}
