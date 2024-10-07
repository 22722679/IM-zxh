package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageBasic struct {
	UserIdentity string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	Data         string `bson:"data"`
	CreatedAt    int64  `bson:"create_at"`
	UpdateAt     int64  `bson:"update_at"`
	// "account":"账号",
	// "password":"密码",
	// "nickname":"昵称",
	// "sex":1,   // 0-未知， 1-男  2-女
	// "email":"邮箱",
	// "avatar":"头像",
	// "create_at":1,   //创建时间
	// "updata_at":1,   //更新时间

}

func (MessageBasic) CollectionName() string {
	return "message_basic"
}

func InsertOneMessageBasic(mb *MessageBasic) error {
	_, err := Mongo.Collection(MessageBasic{}.CollectionName()).InsertOne(context.Background(), mb)
	return err
}

func GetMessageListByRoomIdentity(roomIdentity string, limit, skip *int64) ([]*MessageBasic, error) {
	data := make([]*MessageBasic, 0)
	cursor, err := Mongo.Collection(MessageBasic{}.CollectionName()).
		Find(context.Background(), bson.M{"room_identity": roomIdentity},
			&options.FindOptions{
				Limit: limit,
				Skip:  skip,
				Sort: bson.D{{
					"create_at", -1,
				}},
			})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		mb := new(MessageBasic)
		err = cursor.Decode(mb)
		if err != nil {
			return nil, err
		}
		data = append(data, mb)
	}
	return data, nil
}
