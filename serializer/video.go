package serializer

import "shVideoGo/model"

//Video 视频序列化器结构体
type Video struct {
	VideoID    int    `json:"video_id"`
	VideoTitle string `json:"video_title"`
	VideoInfo  string `json:"video_info"`
}

//BuildVideo 序列化视频
func BuildVideo(video model.Video) Video {
	return Video{
		VideoID:    video.VideoID,
		VideoTitle: video.VideoTitle,
		VideoInfo:  video.VideoInfo,
	}
}

//BuildVideos 序列化视频列表
func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
