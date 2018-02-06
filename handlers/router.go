package handlers

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", listVideos).Methods(http.MethodGet)
	s.HandleFunc("/video/d290f1ee-6c54-4b01-90e6-d701748f0851", getVideoInfo).Methods(http.MethodGet)
	return logMiddleware(s)
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("start serving request")
		h.ServeHTTP(w, r)
		log.Info("request have been served")
	})
}
