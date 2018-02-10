package handlers

import "net/http"

func listVideos(w http.ResponseWriter, _ *http.Request) {
	list, err := getAvailableVideos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderJson(w, list)
}
