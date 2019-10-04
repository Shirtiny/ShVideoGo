package service

import (
	"shVideoGo/model"
	"shVideoGo/serializer"
)

//ListVideosService 结构体 注意这个结构体为空
type ListVideosService struct {
}

//List 显示视频列表
func (service *ListVideosService) List() serializer.Response {
	var videos []model.Video
	error := model.DB.Find(&videos).Error
	if error != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "查找视频时出错，可能是数据库错误",
			Error: error.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "查询列表完成",
		Data: serializer.BuildVideos(videos),
	}
}
