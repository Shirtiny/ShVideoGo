package service

import (
	"shVideoGo/model"
	"shVideoGo/serializer"

	"github.com/gin-gonic/gin"
)

// CreateVideoService 创建视频的服务结构体（有点像类）
type CreateVideoService struct {
	VideoID    int    `form:"video_id" json:"video_id"`
	VideoTitle string `form:"video_title" json:"video_title" binding:"required"`
	VideoInfo  string `form:"video_info" json:"video_info"`
}

// Create 创建视频函数
func (service *CreateVideoService) Create(c *gin.Context) serializer.Response {
	//绑定参数，模型:service.xx
	video := model.Video{
		VideoID:    service.VideoID,
		VideoTitle: service.VideoTitle,
		VideoInfo:  service.VideoInfo,
	}
	//操作数据库，把错误返回
	error := model.DB.Create(&video).Error
	if error != nil {
		return serializer.Response{
			Code: 500,
			Msg:  "数据库操作失败",
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
