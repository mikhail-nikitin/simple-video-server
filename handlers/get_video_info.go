package handlers

import "net/http"

func getVideoInfo(w http.ResponseWriter, _ *http.Request) {
	videoDetails := Video{
		Id:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospective Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		Url:       "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
	}
	renderJson(w, videoDetails)
}
