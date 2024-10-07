package helper

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	//"go/token"

	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"github.com/zhangzhaojian/go-mongodb/x/mongo/driver/uuid"
)

type UserClaims struct {
	//Identity  string  `json:"identity"`
	Identity string `json:"identity"`
	Email    string `json:email`
	jwt.StandardClaims
}

// 生成md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("im")

// 生成token
func GenerateToken(identity, email string) (string, error) {
	UserClaim := &UserClaims{
		Identity:       identity,
		Email:          email,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// SendCode
// 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "zhangxiaohu<zdefined@163.com>"
	//e.From = "Get <getcharzhaopan@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送， 请查收"
	e.HTML = []byte("<h1>您的验证码：" + code + "</h1>")

	//return e.Send("smtp.qq.com:25", smtp.PlainAuth("", "781824252@qq.com", "jfpgzecoiupnbcce", "smtp.qq.com"))
	return e.Send("smtp.163.com:25", smtp.PlainAuth("", "zdefined@163.com", "QDad2Xie553x6n6d", "smtp.163.com"))
	
	
	//e.SendWithTLS("smtp.qq.com:25",
	//smtp.PlainAuth("", "781824252@qq.com", "jfpgzecoiupnbcce", "stmp.qq.com"),
	//&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})

}

// 验证码生成
func GetCode() string {
	rand.Seed(time.Now().UnixNano())

	res := ""
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(10))
	}
	return res
}

// 生成唯一码 identity
func GetUUID() string {
	u, err := uuid.New()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", u)
}

