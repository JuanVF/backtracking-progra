package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	sospechoso := []string{"El/La mejor amigo(a)", "El/la novio(a)", "El/la vecino(a)", "El mensajero", "El extraño", "El/la hermanastro(a)", "El/la colega de trabajo"}
	arma := []string{"Pistola", "Cuchillo", "Machete", "Pala", "Bate", "Botella", "Tubo", "Cuerda"}
	motivo := []string{"Venganza", "Celos", "Dinero", "Accidente", "Drogas", "Robo"}
	cuerpo := []string{"Cabeza", "Pecho", "Abdomen", "Espalda", "Piernas", "Brazos"}
	lugar := []string{"Sala", "Comedor", "Baño", "Terraza", "Cuarto", "Garage", "Patio", "Balcón", "Cocina"}

	posibilidades := [][]string{sospechoso, arma, motivo, cuerpo, lugar}
	rest := rest_generator(15, posibilidades)
	solv := solution(posibilidades, rest)
	encontrada := []string{}
	eliminadas := []string{}

	iteraciones1, _ := FuerzaBruta(posibilidades, solv, encontrada)
	iteraciones2, _, _ := backtracking(posibilidades, solv, encontrada, rest, eliminadas)
	fmt.Println(rest)
	fmt.Println(solv)

	fmt.Printf("Iteraciones en fuerza bruta: %d\n", iteraciones1)
	fmt.Printf("Iteraciones en bacttracking: %d\n", iteraciones2)
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

func select_elim(try []string, solv []string) string {
	delete := ""
	for delete == "" {
		i := rand.Intn(len(try))
		delete = try[i]
		if find(solv, delete) {
			delete = ""
		}
	}
	return delete
}

func backtracking(posibilidades [][]string, solucion, encontrada []string, rest [][]string, eliminadas []string) (int, bool, string) {

	// Caso base
	if len(posibilidades) == 0 {
		fmt.Println(encontrada)
		if Equals(solucion, encontrada) {
			return 1, true, "*"
		}
		elm := select_elim(solucion, encontrada)
		println(elm)
		return 1, false, elm
	}

	amount := 0
	elim := ""
	// Probamos cada array
	for _, posibilidad := range posibilidades[0] {
		generada := append(encontrada, posibilidad)
		if sol_validation(generada, rest, eliminadas) {
			iteraciones, finded, elim := backtracking(posibilidades[1:], solucion, generada, rest, eliminadas)
			if elim != "" {
				eliminadas = append(eliminadas, elim)
			}
			amount += iteraciones
			if finded {
				return amount, true, elim
			}
		}
	}

	return amount, false, elim
}

func sol_validation(solv []string, rest [][]string, eliminadas []string) bool {
	for i := 0; i < len(rest); i++ {
		if find(solv, rest[i][0]) || find(solv, rest[i][1]) {
			return false
		}
	}
	for i := 0; i < len(eliminadas); i++ {
		if find(solv, eliminadas[i]) {
			return false
		}
	}
	return true
}

func solution(posibilidades [][]string, rest [][]string) []string {
	aux := false
	eliminadas := []string{}
	solv := make([]string, len(posibilidades))
	for !aux {
		solv = try_select(posibilidades)
		aux = sol_validation(solv, rest, eliminadas)
	}
	return solv
}

func rest_generator(cant int, posibilidades [][]string) [][]string {

	rest := make([][]string, cant)

	for i := 0; i < cant; i++ {

		rest[i] = make([]string, 2)
		l1 := posibilidades[rand.Intn(len(posibilidades))]
		l2 := l1

		for Equals(l1, l2) {
			l2 = posibilidades[rand.Intn(len(posibilidades))]
		}

		rest[i][0] = l1[rand.Intn(len(l1))]
		rest[i][1] = l2[rand.Intn(len(l2))]
	}

	return rest
}

func try_select(posibilidades [][]string) []string {
	solv := make([]string, 5)
	for i := 0; i < len(posibilidades); i++ {
		solv[i] = posibilidades[i][rand.Intn(len(posibilidades[i]))]
	}
	return solv
}

func Equals(arr1, arr2 []string) bool {
	return reflect.DeepEqual(arr1, arr2)
}

func find(a []string, x string) bool {
	for i := 0; i < len(a); i++ {
		if x == a[i] {
			return true
		}
	}
	return false
}
