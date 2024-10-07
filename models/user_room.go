package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRoom struct {
	UserIdentity string `bson:"user_identity"`
	RoomIdentity string `bson:"room_identity"`
	RoomType     int    `bson:"room_type"` //房间类型  [1-独聊房间  2-群聊房间]
	CreatedAt    int64  `bson:"create_at"`
	UpdateAt     int64  `bson:"update_at"`
}

func (UserRoom) CollectionName() string {
	return "user_room"
}

func GetUserRoomByUserIdentityRoomIdentity(userIdentity, roomIdentity string) (*UserRoom, error) {
	ur := new(UserRoom)
	err := Mongo.Collection(UserRoom{}.CollectionName()).
		FindOne(context.Background(), bson.D{{"user_identity", userIdentity}, {"room_identity", roomIdentity}}).
		Decode(ur)
	return ur, err
}

func GetUserRoomByRoomIdentity(roomIdentity string) ([]*UserRoom, error) {
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).Find(context.Background(), bson.D{{"room_identity", roomIdentity}})
	if err != nil {
		return nil, err
	}
	urs := make([]*UserRoom, 0)
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			//t.Fatal(err)
			//fmt.Println("error", err)
			return nil, err
		}
		urs = append(urs, ur)
	}
	return urs, nil
}

func JudgeUserIdFriend(userIdentity1, userIdentity2 string) bool {
	//查询userIdentity1 单聊列表
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).
		Find(context.Background(), bson.D{{"user_identity", userIdentity1}, {"room_type", 1}})
	roomIdentities := make([]string, 0)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return false
	}
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			log.Printf("Decode Error:%v\n", err)
			return false
		}
		roomIdentities = append(roomIdentities, ur.RoomIdentity)
	}
	//	获取关联  userIdentity2 单聊房间数
	cnt, err := Mongo.Collection(UserRoom{}.CollectionName()).
		//拿到userIdentity2文件的数量
		CountDocuments(context.Background(), bson.M{"user_identity": userIdentity2, "room_identity": bson.M{"$in": roomIdentities}})
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return false
	}
	if cnt > 0 {
		return true
	}
	return false
}

func InsertOneUserRoom(ur *UserRoom) error {
	_, err := Mongo.Collection(UserRoom{}.CollectionName()).InsertOne(context.Background(), ur)
	return err
}

func GetUserRoomIdentity(userIdentity1, userIdentity2 string) string {
	//查询userIdentity1 单聊列表
	cursor, err := Mongo.Collection(UserRoom{}.CollectionName()).
		Find(context.Background(), bson.D{{"user_identity", userIdentity1}, {"room_type", 1}})
	roomIdentities := make([]string, 0)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return ""
	}
	for cursor.Next(context.Background()) {
		ur := new(UserRoom)
		err := cursor.Decode(ur)
		if err != nil {
			log.Printf("Decode Error:%v\n", err)
			return ""
		}
		roomIdentities = append(roomIdentities, ur.RoomIdentity)
	}
	//	获取关联  userIdentity2 单聊房间数
	ur := new(UserRoom)
	err = Mongo.Collection(UserRoom{}.CollectionName()).
		//拿到userIdentity2文件的数量
		FindOne(context.Background(), bson.M{"user_identity": userIdentity2, "room_identity": bson.M{"$in": roomIdentities}}).Decode(ur)
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return ""
	}
	return ur.RoomIdentity
}


func DeleteUserRoom(roomIdentity string) error {
	_, err := Mongo.Collection(UserRoom{}.CollectionName()).
		DeleteOne(context.Background(), bson.M{"room_identity": roomIdentity})
	if err != nil {
		log.Printf("[DB ERROR]:%v\n", err)
		return err
	}
	return nil
}