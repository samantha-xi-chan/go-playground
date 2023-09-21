package play132_websocktserver

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许任何来源的WebSocket连接
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	log.Println(" - - - WebSocket 连接已建立  - - - ")

	for {
		// 读取客户端发送的消息
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// 将消息原样发送回客户端
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

		//conn.Close()
	}
}

func Play() {
	addr := ":8080"
	http.HandleFunc("/socket", handleConnection)
	log.Println("WebSocket服务器已启动，监听地址： ", addr)
	http.ListenAndServe(addr, nil)
}
