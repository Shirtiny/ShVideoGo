package service

import (
	"shVideoGo/model"
	"shVideoGo/serializer"
)

// DeleteVideoService 创建视频的服务结构体（有点像类）
type DeleteVideoService struct {
}

// Delete 创建视频函数
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	//先查询该视频是否存在
	error := model.DB.First(&video, id).Error
	//若出错
	if error != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "未找到指定id的视频",
			Error: error.Error(),
		}
	}
	//存在则执行删除
	error = model.DB.Delete(&video).Error
	if error != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "删除失败",
			Error: error.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "删除成功",
	}
}
