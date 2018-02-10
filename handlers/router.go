package handlers

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"database/sql"
)

var dbConn *sql.DB

func Router(db *sql.DB) http.Handler {
	dbConn = db
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", listVideos).Methods(http.MethodGet)
	s.HandleFunc("/video/{id}", getVideoInfo).Methods(http.MethodGet)
	s.HandleFunc("/video", uploadVideo).Methods(http.MethodPost)
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
