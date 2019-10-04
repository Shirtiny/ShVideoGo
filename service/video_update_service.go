package service

import (
	"shVideoGo/model"
	"shVideoGo/serializer"
)

// UpdateVideoService 创建视频的服务结构体（有点像类）
type UpdateVideoService struct {
	VideoID    int    `form:"video_id" json:"video_id"`
	VideoTitle string `form:"video_title" json:"video_title" binding:"required"`
	VideoInfo  string `form:"video_info" json:"video_info"`
}

// Update 修改更新视频函数
func (service *UpdateVideoService) Update(id string) serializer.Response {
	//绑定参数，模型:service.xx
	video := model.Video{}
	//先查询数据库，找到对应user，把错误返回
	error := model.DB.First(&video, id).Error
	if error != nil {
		return serializer.Response{
			Code: 500,
			Msg:  "数据库中未能找到对应视频",
			//返回错误信息
			Error: error.Error(),
		}
	}

	//修改，把service通过c绑定，接收到的json数据，赋给model
	video.VideoID = service.VideoID
	video.VideoTitle = service.VideoTitle
	video.VideoInfo = service.VideoInfo
	//把video保存到数据库
	error = model.DB.Save(&video).Error
	if error != nil {
		return serializer.Response{
			Code: 500,
			Msg:  "视频保存失败",
			//返回错误信息
			Error: error.Error(),
		}
	}

	return serializer.Response{
		//返回序列化的video，需要调用video的序列化器
		Data: serializer.BuildVideo(video),
		Code: 200,
		Msg:  "视频保存成功",
	}
}
