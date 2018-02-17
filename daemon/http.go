package daemon

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func StartServer(serverUrl string, router http.Handler) *http.Server {
	srv := &http.Server{Addr: serverUrl, Handler: router}
	go func() {
		log.WithFields(log.Fields{"url": serverUrl}).Info("listening")
		log.Fatal(srv.ListenAndServe())
	}()
	return srv
}
