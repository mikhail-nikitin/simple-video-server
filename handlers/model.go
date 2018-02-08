package handlers

type Video struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	Url       string `json:"url"`
}

type VideoList []Video

var allAvailableVideos = VideoList{
	Video{
		Id:        "d290f1ee-6c54-4b01-90e6-d701748f0851",
		Name:      "Black Retrospective Woman",
		Duration:  15,
		Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		Url:       "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
	},
	Video{
		Id:        "sldjfl34-dfgj-523k-jk34-5jk3j45klj34",
		Name:      "Go Rally TEASER-HD",
		Duration:  41,
		Thumbnail: "/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/screen.jpg",
		Url:       "/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/index.mp4",
	},
	Video{
		Id:        "hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345",
		Name:      "Танцор",
		Duration:  92,
		Thumbnail: "/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/screen.jpg",
		Url:       "/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/index.mp4",
	},
}

func getAvailableVideos() VideoList {
	return allAvailableVideos
}

func getVideoById(id string) *Video {
	for i := 0; i < len(allAvailableVideos); i++ {
		v := &allAvailableVideos[i]
		if v.Id == id {
			return v
		}
	}
	return nil
}