package handlers

import (
	"net/http"
	json "encoding/json"
	"io"
	"github.com/sirupsen/logrus"
)

func getVideoInfo(w http.ResponseWriter, _ *http.Request) {
	v := Video{
		Id:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospective Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		Url: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
	}

	s, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; character-set=utf8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(s)); err != nil {
		logrus.WithField("err", err.Error()).Error("write response error")
	}
}