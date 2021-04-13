package main

import (
	"encoding/json"
	"time"

	"github.com/JuanVF/gogame-server/sockets"
	"github.com/gorilla/websocket"
)

// Esta es la accion que se va a ejecutar cuando el usuario pide el algoritmo de fuerza bruta
func TestFuerzaBruta(ws *websocket.Conn, msg sockets.Message) {
	solucion := GetSolution(GetCategorias())

	// Enviamos la solucion al front
	solMsgJson, _ := json.Marshal(solucion[0])
	solMsg := sockets.Message{
		ID:   3,
		Json: string(solMsgJson),
	}

	sockets.GetInstance().SendTo(solMsg, ws)

	encontrada := make([]Categorias, 0)
	mensajes := make([]sockets.Message, 0)
	eliminadas := make(map[string]bool)

	// Medimos el tiempo
	initTime := GetCurrentTime()

	iteraciones, _ := FuerzaBruta(GetCategorias(), solucion[0], encontrada, &eliminadas, &mensajes)

	time := GetCurrentTime() - initTime

	// Enviamos al front
	sockets.GetInstance().SendTo(sockets.Message{
		ID:      2,
		Numbers: []int{int(time), iteraciones},
	}, ws)

	// Enviamos los resultados al usuario
	for _, mensaje := range mensajes {
		sockets.GetInstance().SendTo(mensaje, ws)
	}
}

// Esta es la accion que se va a ejecutar cuando el usuario pide el algoritmo de fuerza bruta
func TestFuerzaBrutaPura(ws *websocket.Conn, msg sockets.Message) {
	solucion := GetSolution(GetCategorias())

	// Enviamos la solucion al front
	solMsgJson, _ := json.Marshal(solucion[0])
	solMsg := sockets.Message{
		ID:   3,
		Json: string(solMsgJson),
	}

	sockets.GetInstance().SendTo(solMsg, ws)

	encontrada := make([]Categorias, 0)
	mensajes := make([]sockets.Message, 0)

	// Medimos el tiempo
	initTime := GetCurrentTime()

	iteraciones, _ := FuerzaBrutaCompleta(GetCategorias(), solucion[0], encontrada, &mensajes)

	time := GetCurrentTime() - initTime

	// Enviamos al front
	sockets.GetInstance().SendTo(sockets.Message{
		ID:      2,
		Numbers: []int{int(time), iteraciones},
	}, ws)

	// Enviamos los resultados al usuario
	for _, mensaje := range mensajes {
		sockets.GetInstance().SendTo(mensaje, ws)
	}
}

// Esta es la accion que se va a ejecutar cuando el usuario pide el algoritmo de fuerza bruta
func TestBacktracking(ws *websocket.Conn, msg sockets.Message) {
	solucion := GetSolution(GetCategorias())
	rest := GenerateRest(msg.Number, solucion[1])

	// Enviamos la solucion al front
	solMsgJson, _ := json.Marshal(solucion[0])
	solMsg := sockets.Message{
		ID:   3,
		Json: string(solMsgJson),
	}

	sockets.GetInstance().SendTo(solMsg, ws)

	// Inicializamos variables para bakctracking
	encontrada := make([]Categorias, 0)
	mensajes := make([]sockets.Message, 0)
	eliminadas := make(map[string]bool)

	// Medimos el tiempo
	initTime := GetCurrentTime()

	iteraciones, _ := Backtracking(GetCategorias(), solucion[0], encontrada, rest, &eliminadas, &mensajes)

	time := GetCurrentTime() - initTime

	// Enviamos al front
	sockets.GetInstance().SendTo(sockets.Message{
		ID:      2,
		Numbers: []int{int(time), iteraciones},
	}, ws)

	// Enviamos los resultados al usuario
	for _, mensaje := range mensajes {
		sockets.GetInstance().SendTo(mensaje, ws)
	}
}

// Desc: Retorna el tiempo del sistema en milisegundos
func GetCurrentTime() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
