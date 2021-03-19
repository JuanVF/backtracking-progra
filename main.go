package main

import (
	"log"
	"net/http"

	"github.com/JuanVF/gogame-server/sockets"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{}
var Handler *sockets.SocketHandler = sockets.GetInstance()

func HandlerExample(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrader.Upgrade(w, r, nil) // Con esto generamos el websocket

	// Esto es un estandar en GoLang, manejar los errors asi
	if err != nil {
		log.Fatal(err)
	}

	// Aqui agregamos una conexion
	Handler.AddConn(ws)

	// Aqui hacemos un Listener para escuchar al usuario
	for {
		var msg sockets.Message

		// Esperamos que el usuario mande algo
		err := ws.ReadJSON(&msg)

		// Si hubo un error desconectamos al usaurio
		if err != nil {
			Handler.RemoveConn(ws)

			break
		}

		// Sacamos la accion que viene del ID del mensaje y la ejecutamos
		action := Handler.GetAction(msg.ID)

		action()
	}
}

func FuncionEjemplo() {
	// Aqui hacemos cualquier cosa x
}

func main() {
	server := NewServer(5000)

	server.Handle("/api", "GET", HandlerExample)

	// Aqui decimos que con el ID 1, agregamos la funcion de ejemplo
	Handler.AddAction(1, FuncionEjemplo)

	server.Listen()
}
