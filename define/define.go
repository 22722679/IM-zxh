package define

//import "os"

var MailPassword = "duibuqi.123"//os.Getenv("MailPassword")

type MessageStruct struct {
	Message 	 string `json:"message"`
	RoomIdentity string `json:"room_identity"`
}


var RegisterPrefix = "TOKEN_"
var ExpireTime = 300