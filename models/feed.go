package models

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

type ResCode int64

const CodeSuccess ResCode = 0

//视频流信息

type VideoInfo struct {
	Id            uint      `json:"id" db:"id`
	Identity      string    `json:"identity" db:"author_id"`
	PlayUrl       string    `json:"play_url" db:"play_url"`
	CoverUrl      string    `json:"cover_url" db:"cover_url"`
	FavoriteCount int       `json:"favorite_count" db:"favorite_count"`
	CommentCount  int       `json:"comment_count" db:"comment_count"`
	Title         string    `json:"title" db:"title"`
	Author        UserBasic `json:"author"` //视频作者信息
}

type Video struct {
	Identity string `json:"identity" db:"author_id"`

	PlayUrl string `json:"play_url" db:"play_url"`

	CoverUrl string `json:"cover_url" db:"cover_url"`

	FavoriteCount int `json:"favorite_count" db:"favorite_count"`

	CommentCount int `json:"comment_count" db:"comment_count"`

	Title string `json:"title" db:"title"`

	Author UserBasic `json:"author"` //视频作者信息

	IsFavorite bool `json:"is_favorite"` //是否点赞

}

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "success",
}

type ResponseData struct {
	Code      ResCode     `json:"status_code"`
	Msg       interface{} `json:"status_msg"`
	VideoInfo interface{} `json:"video_list"`
}

func (rc ResCode) Msg() string {
	msg := codeMsgMap[rc] //若错误返回空
	return msg
}

func SelectVideoInfo() (videoList []*VideoInfo, err error) {

	sqlStr := "select author_id, play_url, cover_url, favorite_count, comment_count, title from video_infos"
	if err := db.Select(&videoList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("数据库中没有视频")
		}
	}
	return
}

func SelectVideoInfoListByUserId(id int64) ([]*Video, error) {
	var UserId []*Video
	VideoList := fmt.Sprintf("select author_id from video_infos where id = %d", id)
	if err := db.Select(&UserId, VideoList); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("没有该用户")
		}
	}
	return UserId, nil
}
