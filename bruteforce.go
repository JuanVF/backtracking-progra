package main

import (
	"encoding/json"
	"reflect"

	"github.com/JuanVF/gogame-server/sockets"
)

// Algoritmo de fuerza bruta para encontrar la solucion
// Este verifica cada una de las posibilidades
func FuerzaBruta(categorias, solucion, encontrada []Categorias, eliminadas *map[string]bool, mensajes *[]sockets.Message) (int, bool) {
	// Caso base
	if len(categorias) == 0 {
		json, _ := json.Marshal(encontrada)
		message := sockets.Message{
			ID:   0,
			Json: string(json),
		}

		// Determinamos si es solucion
		isSolution := reflect.DeepEqual(solucion, encontrada)

		// Si no es solucion se solicita una "pista"
		if !isSolution {
			message.ID = 1
			eliminada := SelectElim(solucion, encontrada)

			(*eliminadas)[eliminada] = true
		}

		(*mensajes) = append(*mensajes, message)

		return 1, isSolution
	}

	amount := 0

	// Probamos cada posibilidad
	for _, posibilidad := range categorias[0].Posibilidades {

		if (*eliminadas)[posibilidad] {
			continue
		}

		generada := append(encontrada, Categorias{
			Categoria:     categorias[0].Categoria,
			Posibilidades: []string{posibilidad},
		})

		iteraciones, finded := FuerzaBruta(categorias[1:], solucion, generada, eliminadas, mensajes)
		amount += iteraciones

		if finded {
			return amount + 1, true
		}
	}

	return amount + 1, false
}
