package service

import (
	"im/config"
	"im/helper"
	"im/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)
 
var upgrader = websocket.Upgrader{}
var wc = make(map[string]*websocket.Conn)

func WebsocketMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统异常" + err.Error(),
		})
		return
	}
	defer conn.Close()
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	wc[uc.Identity] = conn
	for {
		ms := new(config.MessageStruct)
		err := conn.ReadJSON(ms)
		if err != nil {
			log.Printf("Read Error:%v\n", err)
			return
		}
		//  TODO: 判断用户是否属于消息体的房间
		_, err = models.GetUserRoomByUserIdentityRoomIdentity(uc.Identity, ms.RoomIdentity)
		if err != nil {
			log.Printf("UserIdentity:%v RoomIdentity:%v Not exists\n", uc.Identity, ms.RoomIdentity)
			return
		}
		// TODO: 保存消息
		mb := &models.MessageBasic{
			UserIdentity: uc.Identity,
			RoomIdentity: ms.RoomIdentity,
			Data:         ms.Message,
			CreatedAt:    time.Now().Unix(),
			UpdateAt:     time.Now().Unix(),
		}
		err = models.InsertOneMessageBasic(mb)
		if err != nil {
			log.Printf("[DB ERROR]:%v\n", err)
			return
		}
		// TODO: 获取在特定房间的在线用户
		userRooms, err := models.GetUserRoomByRoomIdentity(ms.RoomIdentity)
		if err != nil {
			log.Printf("[DB ERRIR]:%v\n", err)
			return
		}
		for _, room := range userRooms {
			if cc, ok := wc[room.UserIdentity]; ok {
				err := cc.WriteMessage(websocket.TextMessage, []byte(ms.Message))
				if err != nil {
					log.Printf("Write Message Error:%v\n", err)
					return
				}
			}
		}
	}
}
