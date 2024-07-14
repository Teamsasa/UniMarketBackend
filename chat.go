package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// 接続されるクライアント
var clients = make(map[*websocket.Conn]bool)

// メッセージブロードキャストチャネル
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{
	// 気が向いたら、もっと厳しくする
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

type Message struct {
	ID      int    `json:"id"`
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

type SendMessage struct {
	Type string  `json:"type"` // "message"
	Data Message `json:"data"`
}

type sendHistory struct {
	Type string    `json:"type"` // "history"
	Data []Message `json:"data"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ws handler called...")

	// 送られてきたGETリクエストをWebSocketにアップグレード
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// クライアントを登録
	clients[ws] = true

	for {
		var message Message
		// 新しいメッセージをJSONとして読み込み、Message構造体にマッピング
		err := ws.ReadJSON(&message)

		fmt.Println("message received: ", message)

		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// 受け取ったメッセージをbroadcastチャネルに送る
		broadcast <- message
	}
}

func handleMessages() {
	for {
		// broadcastチャネルから次のメッセージを受け取る
		message := <-broadcast

		SendContent := SendMessage{
			Type: "message",
			Data: message,
		}

		// クライアントにメッセージを送信
		for client := range clients {

			err := client.WriteJSON(SendContent)

			fmt.Println("message sent: ", message)

			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
