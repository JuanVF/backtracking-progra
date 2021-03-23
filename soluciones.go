package main

import (
	"fmt"
	"reflect"
)

// TODO: El de restricciones solo eliminar uno
func Restricciones() []string {
	return nil
}

func FuerzaBruta(posibilidades [][]string, solucion, encontrada []string) (int, bool) {
	// Caso base
	if len(posibilidades) == 0 {
		if Equals(solucion, encontrada) {
			fmt.Println(encontrada)
			return 1, true
		}

		return 1, false
	}

	amount := 0

	// Probamos cada array
	for _, posibilidad := range posibilidades[0] {
		generada := append(encontrada, posibilidad)

		iteraciones, finded := FuerzaBruta(posibilidades[1:], solucion, generada)
		amount += iteraciones

		if finded {
			return amount, true
		}
	}

	return amount, false
}

func Equals(arr1, arr2 []string) bool {
	return reflect.DeepEqual(arr1, arr2)
}
