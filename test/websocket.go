package test

import (
	"flag"
	"log"
	"net/http"

	//"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)
	
var addr = flag.String("addr", "localhost:8080", "http service address") //服务端地址

var upgrader = websocket.Upgrader{} //  将http分解为websocket
var ws = make(map[*websocket.Conn]struct{})

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	ws[c] = struct{}{}
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		for conn := range ws {
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

//HTTP->websocket
// func main() {
// 	//func TestServerServer(t *testing.T) {
// 	//
// 	http.HandleFunc("/echo", echo)
// 	log.Fatal(http.ListenAndServe(*addr, nil))
// }

// gin->websocket
func main() {
	r := gin.Default()
	//路由
	r.GET("/echo", func(ctx *gin.Context) {
		echo(ctx.Writer, ctx.Request)
	})
	r.Run(":8080")

}
