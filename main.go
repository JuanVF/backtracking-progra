package main

import (
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

		// Esto se queda esperando a un mensaje del usuario
		err := ws.ReadJSON(&msg)

		// Si hay error lo desconectamos
		if err != nil {
			log.Printf("Message error: %v", err)

			sockets.GetInstance().RemoveConn(ws)
			return
		}

		// Hacemos print de lo que el usuario envio
		log.Printf("User sended: %v\n", msg)

		// Ejecutamos una accion en base a lo que el usuario envio
		sockets.GetInstance().GetAction(msg.ID)(ws, msg)
	}
}

func main() {
	// Agregamos las funciones para que el front envie peticiones
	// para fuerza bruta y backtracking
	sockets.GetInstance().AddAction(0, TestFuerzaBruta)
	sockets.GetInstance().AddAction(1, TestBacktracking)

	// Inicializamos el server
	server := NewServer(5000)

	// En la ruta /api por metodo get se va a manejar el websocket
	server.Handle("/api", "GET", HandlerUsers)

	// El servidor inicia
	server.Listen()
}
