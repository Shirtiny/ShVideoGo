package api

import (
	"shVideoGo/serializer"

	"github.com/gin-gonic/gin"
)

//TestVideo 测试接口
func TestVideo(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 200,
		Msg:  "测试成功",
	})
}
