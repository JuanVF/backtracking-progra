package main

import (
	"math/rand"
	"reflect"
	"time"
)

// Algoritmo de Backtracking
func Backtracking(categorias, solucion, encontrada []Categorias, rest [][]string, eliminadas *map[string]bool) (int, bool) {
	// Caso base
	if len(categorias) == 0 {
		isSolution := reflect.DeepEqual(solucion, encontrada)

		if !isSolution {
			eliminada := SelectElim(solucion, encontrada)

			(*eliminadas)[eliminada] = true
		}

		return 1, isSolution
	}

	amount := 0

	// Probamos cada array
	for _, posibilidad := range categorias[0].Posibilidades {
		if (*eliminadas)[posibilidad] {
			continue
		}

		generada := append(encontrada, Categorias{
			Categoria:     categorias[0].Categoria,
			Posibilidades: []string{posibilidad},
		})

		if isRightSolution(generada, rest) {
			iteraciones, finded := Backtracking(categorias[1:], solucion, generada, rest, eliminadas)
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
