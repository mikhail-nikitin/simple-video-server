package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func getVideoInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	video, err := getVideoByKey(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if video != nil {
		renderJson(w, video)
	} else {
		http.NotFound(w, r)
	}
}
