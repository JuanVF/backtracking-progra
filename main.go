package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JuanVF/gogame-server/sockets"
	"github.com/gorilla/websocket"
)

// Variable para generar la coneccion de websocket
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handler para mantener la coneccion por websocket
func HandlerUsers(w http.ResponseWriter, req *http.Request) {
	ws, err := Upgrader.Upgrade(w, req, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	// Agregamos el usuario a la lista de usuarios conectados
	sockets.GetInstance().AddConn(ws)

	// Esto funciona basicamente como un listener del usuario
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
	rest := GenerateRest(20, GetCategorias())
	solv := GetSolution(GetCategorias(), rest)

	encontrada := []Categorias{}
	eliminadas := make(map[string]bool)
	mensajes := make([]sockets.Message, 0)

	iteraciones2, _ := Backtracking(GetCategorias(), solv, encontrada, rest, &eliminadas)
	iteraciones, _ := FuerzaBruta(GetCategorias(), solv, encontrada, &mensajes)
	fmt.Printf("Iteraciones en Fuerza bruta: %d\n", iteraciones)
	fmt.Printf("Iteraciones en Backtracking: %d\n", iteraciones2)
}
