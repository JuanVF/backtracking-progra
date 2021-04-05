package main

import (
	"encoding/json"
	"reflect"

	"github.com/JuanVF/gogame-server/sockets"
)

// Algoritmo de fuerza bruta para encontrar la solucion
// Este verifica cada una de las posibilidades
func FuerzaBruta(categorias, solucion, encontrada []Categorias, mensajes *[]sockets.Message) (int, bool) {
	// Caso base
	if len(categorias) == 0 {
		json, _ := json.Marshal(encontrada)
		message := sockets.Message{
			ID:   0,
			Json: string(json),
		}

		isSolution := reflect.DeepEqual(solucion, encontrada)

		if isSolution {
			message.ID = 1
		}

		(*mensajes) = append(*mensajes, message)

		return 1, isSolution
	}

	amount := 0

	// Probamos cada posibilidad
	for _, posibilidad := range categorias[0].Posibilidades {
		generada := append(encontrada, Categorias{
			Categoria:     categorias[0].Categoria,
			Posibilidades: []string{posibilidad},
		})

		iteraciones, finded := FuerzaBruta(categorias[1:], solucion, generada, mensajes)
		amount += iteraciones

		if finded {
			return amount + 1, true
		}
	}

	return amount + 1, false
}
