package models

import "mime/multipart"

//发布视频请求

type PublishActionRequest struct {
	Data  *multipart.FileHeader `json:"data"`  //上传的视频
	Token string                `json:"token"` //用户权限鉴别
	Tiele string                `json:"title"` //视频的标题
}

type PublishListResponse struct {
	StatusCode int64   `json:"status_code"` //状态码，成功200，失败其他
	StatusMsg  string  `json:"status_msg"`  //返回状态描述
	UserList   []Video `json:"user_list"`   //用户信息列表

}
