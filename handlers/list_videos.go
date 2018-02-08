package handlers

import "net/http"

func listVideos(w http.ResponseWriter, _ *http.Request) {
	renderJson(w, getAvailableVideos())
}
