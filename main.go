package main

import (
	"log"
	"net/http"

	"github.com/JuanVF/gogame-server/sockets"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandlerUsers(w http.ResponseWriter, req *http.Request) {
	ws, err := Upgrader.Upgrade(w, req, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	sockets.GetInstance().AddConn(ws)

	for {
		var msg sockets.Message

		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Printf("Message error: %v", err)

			sockets.GetInstance().RemoveConn(ws)
			break
		}

		log.Printf("User sended: %v\n", msg)

		sockets.GetInstance().GetAction(msg.ID)(ws)
	}
}

func main() {
	sockets.GetInstance().AddAction(0, TestFuerzaBruta)

	server := NewServer(5000)

	server.Handle("/api", "GET", HandlerUsers)

	server.Listen()
}
