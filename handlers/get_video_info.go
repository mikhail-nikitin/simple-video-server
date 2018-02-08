package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
)

func getVideoInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	videoId := vars["id"]
	if video := getVideoById(videoId); video != nil {
		renderJson(w, video)
	} else {
		http.NotFound(w, r)
	}
}
