package models

import (
	"context"

	//"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBasic struct {
	Identity      string `bson:"identity"`
	Account       string `bson:"account"`                            // 账号
	Password      string `bson:"password"`                           // 密码
	Nickname      string `bson:"nickname"`                           // 昵称
	Sex           int    `bson:"sex"`                                // 0-未知， 1-男  2-女
	Email         string `bson:"email"`                              // 邮箱
	Avatar        string `bson:"avatar"`                             // 头像
	FollowCount   int64  `json:"follow_count" db:"follow_count"`     //关注总数
	FollowerCount int64  `json:"follower_count" db:"follower_count"` //粉丝总数
	IsFollow      bool   `json:"is_follow" db:"is_follow"`           //是否被关注
	CreatedAt     int64  `bson:"create_at"`                          //创建时间
	UpdateAt      int64  `bson:"update_at"`                          //更新时间
	//BackgroundImages string `json:"background_images"`                  //用户个人页面顶部大图
	//Signature        string `json:"signature"`                          //个人简介
	//TotalFavorited   string `json:"total_favorited"`                    //获赞数量
	//WorkCount        int64  `json:"work_count"`                         //作品数
	//FavoriteCount    int64  `json:"favorite_count"`                      //喜欢数
}

func (UserBasic) CollectionName() string {
	return "user_basic"
}

func GetUserBasicByAccountPassword(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}, {"password", password}}).
		Decode(ub)
	return ub, err
}

func GetUserBasicByIdentity(identity string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"identity", identity}}).
		Decode(ub)
	return ub, err
}

func GetUserBasicByAccount(account string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection(UserBasic{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"account", account}}).
		Decode(ub)
	return ub, err
}

func GetUserBasicCountByEmail(email string) (int64, error) {
	return Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"email", email}})
}

func GetUserBasicCountByAccount(account string) (int64, error) {
	return Mongo.Collection(UserBasic{}.CollectionName()).
		CountDocuments(context.Background(), bson.D{{"account", account}})
}

func InsertOneUserBasic(ub *UserBasic) error {
	_, err := Mongo.Collection(UserBasic{}.CollectionName()).InsertOne(context.Background(), ub)
	return err
}
