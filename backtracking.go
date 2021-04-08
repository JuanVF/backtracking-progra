package main

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"time"

	"github.com/JuanVF/gogame-server/sockets"
)

// Algoritmo de Backtracking
func Backtracking(categorias, solucion, encontrada []Categorias, rest [][]string, eliminadas *map[string]bool, mensajes *[]sockets.Message) (int, bool) {
	// Caso base
	if len(categorias) == 0 {
		json, _ := json.Marshal(encontrada)
		message := sockets.Message{
			ID:   1,
			Json: string(json),
		}

		// Determinamos si la solucion encontrada es la correcta
		isSolution := reflect.DeepEqual(solucion, encontrada)

		// Si no es solucion se solicita una "pista"
		if !isSolution {
			message.ID = 0
			eliminada := SelectElim(solucion, encontrada)

			(*eliminadas)[eliminada] = true
		}

		(*mensajes) = append(*mensajes, message)

		return 1, isSolution
	}

	amount := 0

	// Probamos cada array
	for _, posibilidad := range categorias[0].Posibilidades {
		if (*eliminadas)[posibilidad] {
			continue
		}

		// Vamos generando la solucion
		generada := append(encontrada, Categorias{
			Categoria:     categorias[0].Categoria,
			Posibilidades: []string{posibilidad},
		})

		// Si es una solucion correcta
		if isRightSolution(generada, rest) {
			iteraciones, finded := Backtracking(categorias[1:], solucion, generada, rest, eliminadas, mensajes)
			amount += iteraciones

			if finded {
				return amount + 1, true
			}
		}
	}
	return amount + 1, false
}

// Busca una opcion aleatoria y la retorna
// Coste del algoritmo O(n) : n => tamano de sol
func SelectElim(sol, encontrada []Categorias) string {
	rand.Seed(time.Now().UnixNano())

	tmp := []string{}
	toDelete := make(map[string]bool)

	// Esto parece O(n*n) pero es O(5n)
	for _, categoria := range sol { // Maximo 5 loops
		for _, posibilidad := range categoria.Posibilidades { // Maximo n loops
			toDelete[posibilidad] = true
		}
	}

	// Lo mismo aplica aqui
	for _, categoria := range encontrada {
		for _, posibilidad := range categoria.Posibilidades {
			if !toDelete[posibilidad] {
				tmp = append(tmp, posibilidad)
			}
		}
	}

	deleted := ""

	if len(tmp) > 0 {
		deleted = tmp[rand.Intn(len(tmp))]
	}

	return deleted
}
