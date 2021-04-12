package main

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"time"

	"github.com/JuanVF/gogame-server/sockets"
)

// Algoritmo de Backtracking
// rest = z
// solucion = m
// encontrada = t
func Backtracking(categorias, solucion, encontrada []Categorias, rest [][]string, eliminadas *map[string]bool, mensajes *[]sockets.Message) (int, bool) {
	// Caso base
	if len(categorias) == 0 {
		json, _ := json.Marshal(encontrada)
		message := sockets.Message{
			ID:   1,
			Json: string(json),
		}

		// Determinamos si la solucion encontrada es la correcta
		isSolution := reflect.DeepEqual(solucion, encontrada) // m*t

		// Si no es solucion se solicita una "pista"
		if !isSolution {
			message.ID = 0
			eliminada := SelectElim(solucion, encontrada, *eliminadas)

			(*eliminadas)[eliminada] = true
		}

		(*mensajes) = append(*mensajes, message)

		return 1, isSolution
	}

	amount := 0

	// Probamos cada array
	// O(n)
	for _, posibilidad := range categorias[0].Posibilidades {
		// Vamos generando la solucion
		// Generada = t
		generada := append(encontrada, Categorias{
			Categoria:     categorias[0].Categoria,
			Posibilidades: []string{posibilidad},
		})

		if len(generada)%2 == 0 {
			solution, eliminada := isRightSolution(generada, rest)

			(*eliminadas)[eliminada] = true
			if solution {
				continue
			}
		}

		/*if IsDeletedInSolution(generada, *eliminadas) {
			continue
		}*/

		iteraciones, finded := Backtracking(categorias[1:], solucion, generada, rest, eliminadas, mensajes)
		amount += iteraciones

		if finded {
			return amount + 1, true
		}
	}
	return amount + 1, false
}

// Busca una opcion aleatoria y la retorna
// Coste del algoritmo O(n) : n => tamano de sol
func SelectElim(sol, encontrada []Categorias, eliminadas map[string]bool) string {
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
			if !toDelete[posibilidad] && !eliminadas[posibilidad] {
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
