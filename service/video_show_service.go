package service

import (
	"shVideoGo/model"
	"shVideoGo/serializer"
)

//ShowVideoService 结构体 注意这个结构体为空
type ShowVideoService struct {
}

//Show show函数有个string类型的id
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	error := model.DB.Find(&video, id).Error
	//若出错
	if error != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "未找到指定id的视频",
			Error: error.Error(),
		}
	}
	//成功则返回数据
	return serializer.Response{
		Code: 200,
		Msg:  "找到了该视频",
		Data: serializer.BuildVideo(video),
	}
}
