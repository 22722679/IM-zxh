package config

//import "os"

var MailPassword = "duibuqi.123"//os.Getenv("MailPassword")

type MessageStruct struct {
	Message 	 string `json:"message"`
	RoomIdentity string `json:"room_identity"`
}


var RegisterPrefix = "TOKEN_"
var ExpireTime = 300


//ftp服务器地址

const FtpIp = "localhost:21"

const FtpName = "a"

const FtpPass = "duibuqi.123"

const HeartTime = 60 * 60



//mysql登录信息

const User = "root"

const PassWord = "duibuqi.123"

const IpMessage = "106.54.210.139:3306"

// oss 视频存储地方
var TencentSecretId = "AKIDaGxmkZGk38AB884JMEHzZ5EQaaNyLYYY"
var TencentSecretKey = "799CF71wkRzNgr38FNMFyaDoLk1CE5mH"
var TencentUploadDir = "/dousheng"
var CosBucket = "https://lj-1307665894.cos.ap-nanjing.myqcloud.com/"

//转入接口

const Port = 8080


