package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const contentTypeJson = "application/json; character-set=utf8"
const headerContentType = "Content-Type"

func renderJson(w http.ResponseWriter, response interface{}) {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(headerContentType, contentTypeJson)
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(responseBytes)); err != nil {
		logrus.WithField("err", err.Error()).Error("was unable to write response")
	}
}
