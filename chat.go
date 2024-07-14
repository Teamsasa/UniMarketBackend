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
		// エラーが発生した場合、exitはせずに関数から抜ける
		log.Fatal(err)
	}
	defer ws.Close()

	// データベースから過去のメッセージを取得
	rows, err := db.Query("SELECT id, sender, content FROM chat_history")
	if err != nil {
		log.Println("error in query: ", err)
		delete(clients, ws)
		return
	}

	// チャット履歴を構造体の配列に格納
	chat_historys := []Message{}
	for rows.Next() {
		var chat_history Message
		err := rows.Scan(&chat_history.ID, &chat_history.Sender, &chat_history.Content)
		if err != nil {
			log.Println("error in scan: ", err)
			return
		}
		chat_historys = append(chat_historys, chat_history)
	}

	// チャット履歴をクライアントに送信
	SendContent := sendHistory{
		Type: "history",
		Data: chat_historys,
	}
	err = ws.WriteJSON(SendContent)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

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

		// メッセージをデータベースに保存
		_, err := db.Exec("INSERT INTO chat_history (sender, content) VALUES ($1, $2)", message.Sender, message.Content)
		if err != nil {
			log.Println("error in insert: ", err)
			continue
		}

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
