package handlers

import (
	"github.com/segmentio/ksuid"
)

type Video struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	Url       string `json:"url"`
}

type VideoList []Video

func getAvailableVideos() (VideoList, error) {
	rows, err := dbConn.Query("SELECT video_key, title, duration, url, thumbnail_url FROM video")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make(VideoList, 0)
	for rows.Next() {
		var video Video
		err = rows.Scan(&video.Id, &video.Name, &video.Duration, &video.Url, &video.Thumbnail)
		if err != nil {
			return nil, err
		}
		list = append(list, video)
	}
	return list, nil
}

func getVideoByKey(key string) (*Video, error) {
	rows, err := dbConn.Query("SELECT video_key, title, duration, url, thumbnail_url FROM video WHERE video_key = ?", key)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var video Video
	if rows.Next() {
		err = rows.Scan(&video.Id, &video.Name, &video.Duration, &video.Url, &video.Thumbnail)
		if err != nil {
			return nil, err
		}
		return &video, nil
	} else {
		return nil, nil
	}
}

func generateVideoKey() string {
	return ksuid.New().String()
}

func addVideo(key string, name string, url string, thumbnailUrl string) error {
	rows, err := dbConn.Query(`INSERT INTO video SET video_key = ?, title = ?, url = ?, thumbnail_url = ?`, key, name, url, thumbnailUrl)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}
