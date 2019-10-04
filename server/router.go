package server

import (
	"os"
	"shVideoGo/api"
	"shVideoGo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	videoAPI := r.Group("/api/video")

	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		//测试
		v1.GET("test", api.TestVideo)
		//增加一个视频
		v1.POST("addVideo", api.AddVideo)

		//增加一个视频，换个group试试
		videoAPI.POST("createVideo", api.AddVideo)
		//展示一个视频的详情/:id
		videoAPI.GET("showVideo/:id", api.ShowVideo)
		//展示所有视频的列表
		videoAPI.GET("listVideos", api.ListVideos)
		//删除一个视频
		videoAPI.DELETE("deleteVideo/:id", api.DeleteVideo)
		//更改一个视频
		videoAPI.PUT("updateVideo/:id", api.UpdateVideo)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
