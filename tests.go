package main

import (
	"fmt"
	"time"

	"github.com/JuanVF/gogame-server/sockets"
	"github.com/gorilla/websocket"
)

// Esta es la accion que se va a ejecutar cuando el usuario pide el algoritmo de fuerza bruta
func TestFuerzaBruta(ws *websocket.Conn) {
	solucion := []Categorias{
		{
			Categoria:     "sospechoso",
			Posibilidades: []string{"El/la colega de trabajo"},
		},
		{
			Categoria:     "arma",
			Posibilidades: []string{"Cuerda"},
		},
		{
			Categoria:     "motivo",
			Posibilidades: []string{"Robo"},
		},
		{
			Categoria:     "cuerpo",
			Posibilidades: []string{"Brazos"},
		},
		{
			Categoria:     "lugar",
			Posibilidades: []string{"Cocina"},
		},
	}

	encontrada := make([]Categorias, 0)
	mensajes := make([]sockets.Message, 0)

	initTime := GetCurrentTime()

	FuerzaBruta(GetCategorias(), solucion, encontrada, &mensajes)

	time := GetCurrentTime() - initTime

	fmt.Printf("Termino en %dms", time)

	sockets.GetInstance().SendTo(sockets.Message{
		ID:     2,
		Number: int(time),
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
