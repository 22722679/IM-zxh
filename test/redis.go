package main

import (
	"context"
	//"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

//testSet
//建立连接
func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:63791",
		Password: "", //no password set
		DB:       0,  //use default DB
	})
	err := rdb.Set(ctx, "key", "valueaaa", time.Second*30).Err()
	if err != nil {
		panic(err)
	}
}

// testGet
// func main() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:63791",
// 		Password: "", //no password set
// 		DB:       0,  //use default DB
// 	})
// 	r, err := rdb.Get(ctx, "key").Result()
//     if err != nil {
//         //t.Fatal(err)
//         fmt.Printf("error:--%v",err)
//         return
//     }
//     fmt.Println(r)
// }
