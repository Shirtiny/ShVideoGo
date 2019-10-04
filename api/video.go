package api

import (
	"shVideoGo/service"

	"github.com/gin-gonic/gin"
)

//AddVideo 增加一个视频
func AddVideo(c *gin.Context) {

	//拿到service对象
	var service service.CreateVideoService
	//绑定service
	error := c.ShouldBind(&service)
	if error != nil {
		//service绑定失败
		c.JSON(500, ErrorResponse(error))
	} else {
		//调用创建视频的service的create方法result :=
		result := service.Create(c)
		c.JSON(200, result)
	}
}

//ShowVideo 展示视频的详情
func ShowVideo(c *gin.Context) {
	//service
	service := service.ShowVideoService{}
	//绑定(可以不绑定，因为结构体为空，绑定的话必须传个json过来，不过这个是get请求没有body)
	error := c.ShouldBind(&service)
	if error != nil {
		c.JSON(500, ErrorResponse(error))
	} else {
		result := service.Show(c.Param("id"))
		c.JSON(200, result)
	}
}

//ListVideos 展示所有视频的列表
func ListVideos(c *gin.Context) {
	//service
	service := service.ListVideosService{}
	error := c.ShouldBind(&service)
	if error != nil {
		c.JSON(500, "服务器出错")
	} else {
		result := service.List()
		c.JSON(200, result)
	}

}

//DeleteVideo 删除一个视频
func DeleteVideo(c *gin.Context) {
	//sercice
	service := service.DeleteVideoService{}
	//结构体为空，不需要接json参数，不用shoudbind
	result := service.Delete(c.Param("id"))
	c.JSON(200, result)

}

//UpdateVideo 更新一个视频
func UpdateVideo(c *gin.Context) {
	//service
	service := service.UpdateVideoService{}
	//绑定
	error := c.ShouldBind(&service)
	if error != nil {
		c.JSON(500, ErrorResponse(error))
	} else {
		result := service.Update(c.Param("id"))
		c.JSON(200, result)
	}

}
