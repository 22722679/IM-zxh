package models

import (
	"context"
	"fmt"
	"log"
	"time"
	"im/config"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
		_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)


var Mongo = InitMongo()
var RDB = InitRedis()
var db *sqlx.DB


func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("Connect MongoDB Error: ", err)
		return nil
	}
	fmt.Println("connect Success")
	return client.Database("im")
}


// 未使用
func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:63791",
	})
}
//mysql  DB
func Init() (err error) {

	db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/douyin?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.PassWord, config.IpMessage))

	if err != nil {

		zap.L().Error("连接数据库失败,错误码为：%v\n", zap.Error(err))

		return

	}


	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns")) //设置数据库的最大打开连接数

	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns")) //设置连接空闲的最大时间,过期的连接可能会在重用之前惰性关闭。

	return

}