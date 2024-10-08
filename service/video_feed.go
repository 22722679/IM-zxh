package service

import (
	"im/config"
	"im/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PublishLists(c *gin.Context) {
	//用user_id查询视频信息
	// 用户基本信息
	userId := c.Query("user_id")
	author, err := models.GetUserBasicByAccount(userId)
	if err != nil {
		c.JSON(http.StatusOK, models.PublishListResponse{
			StatusCode: 200,
			StatusMsg:  "查询用基本信息失败",
			UserList:   nil,
		})
		return
	}
	//视频数据进行封装
	videos := make([]models.Video, 0)
	videos = append(videos, models.Video{
		Author:        *author,
		CommentCount:  config.CommentCount,
		CoverUrl:      config.CoverUrl,
		FavoriteCount: config.FavoriteCount,
		Identity:      strconv.Itoa(config.Id),
		PlayUrl:       config.PlayUrl,
		Title:         config.Tiele,
	})

	c.JSON(http.StatusOK, models.PublishListResponse{
		StatusCode: 200,
		StatusMsg:  "查询成功",
		UserList:   videos,
	})
}

// 视频信息处理
func VideoInfoHandler(c *gin.Context) {
	data, _ := models.SelectVideoInfo()
	ResponseSuccess(c, data)

}
func ResponseSuccess(c *gin.Context, videoList []*models.VideoInfo) {
	c.JSON(http.StatusOK, &models.ResponseData{
		Code:      models.CodeSuccess,
		Msg:       models.CodeSuccess.Msg(),
		VideoInfo: &videoList, //返回的应该为数据库中的信息，该信息videoList用mysql/feed中的SelectVideoInfo()获得
	})
}
