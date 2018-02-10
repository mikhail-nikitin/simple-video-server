package handlers

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
)

func uploadVideo(w http.ResponseWriter, r *http.Request) {
	videoFile, header, err := r.FormFile("file[]")
	if err != nil {
		http.Error(w, "Unable to read \"video\" file", http.StatusBadRequest)
		log.Error(err.Error())
		return
	}

	if header.Header.Get("Content-Type") != "video/mp4" {
		http.Error(w, "Incompatible content type", http.StatusBadRequest)
		return
	}

	videoKey := generateVideoKey()
	videoFilePath := getOutputFileName(videoKey)
	os.MkdirAll(path.Dir(videoFilePath), 0666)

	var outputFile *os.File
	outputFile, err = os.OpenFile(videoFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	defer outputFile.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err.Error())
		return
	}

	err = addVideo(
		videoKey,
		"Title",
		getVideoUrl(videoKey),
		"http://via.placeholder.com/240x180")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err.Error())
		return
	}

	_, err = io.Copy(outputFile, videoFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

func getOutputFileName(videoKey string) string {
	return "content/" + videoKey + "/index.mp4"
}

func getVideoUrl(videoKey string) string {
	return getOutputFileName(videoKey)
}
