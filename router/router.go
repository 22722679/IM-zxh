package router

import (
	"im/middlewares"
	"im/models"
	"im/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	//MySQL启动
	models.Init()
	//用户登录
	r.POST("/login", service.Login)

	//发送验证码
	r.POST("/send/code", service.SendCode)
	//用户注册
	r.POST("/register", service.Register)
	auth := r.Group("/u", middlewares.AuthCheck())
	//用户详情
	auth.GET("/user/detail", service.UserDetail)
	// 查询指定用户的个人信息
	auth.GET("/user/query", service.UserQuery)
	//发送和接收消息
	auth.GET("/websocket/message", service.WebsocketMessage)
	// 聊天记录列表
	auth.GET("/chat/list", service.ChatList)

	// 添加用户为好友
	auth.POST("/user/add", service.UserAdd)
	//删除好友
	auth.DELETE("/user/delete", service.UserDelete)

	// 视频流信息读取
	r.GET("/douyin/publish/list/", service.PublishLists)
	// 点赞列表查询
	auth.GET("/user/favorite/list/", service.FavoriteList)
	// 点赞功能
	auth.POST("/user/favorite/action/", service.FavoriteAction)
	return r
}
