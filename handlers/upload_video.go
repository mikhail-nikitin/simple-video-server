package handlers

import (
	"net/http"
	"io"
	log "github.com/sirupsen/logrus"
	"os"
)

func uploadVideo(w http.ResponseWriter, r *http.Request) {
	videoFile, header, error := r.FormFile("file[]")
	if error != nil {
		http.Error(w, "Unable to read \"video\" file", http.StatusBadRequest)
		log.Error(error.Error())
		return
	}

	if header.Header.Get("Content-Type") != "video/mp4" {
		http.Error(w, "Incompatible content type", http.StatusBadRequest)
		return
	}

	var outputFile *os.File
	outputFile, error = os.OpenFile("content/test/index.mp4", os.O_CREATE|os.O_WRONLY, 0666)
	defer outputFile.Close()
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		log.Error(error.Error())
		return
	}

	_, error = io.Copy(outputFile, videoFile)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		log.Error(error.Error())
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
