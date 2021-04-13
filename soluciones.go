package main

import (
	"math/rand"
	"time"
)

// Esta estructura es para organizar los datos por categorias
type Categorias struct {
	Categoria     string   `json:"Categoria"`
	Posibilidades []string `json:"Posibilidades"`
}

// Esto se puede sustituir por un .json
// Pero retorna una lista con todas las categorias
func GetCategorias() []Categorias {
	categorias := make([]Categorias, 5)

	categorias[0] = Categorias{
		Categoria:     "sospechoso",
		Posibilidades: []string{"El/la mejor amigo(a)", "El/la novio(a)", "El/la vecino(a)", "El mensajero", "El extraño", "El/la hermanastro(a)", "El/la colega de trabajo"},
	}

	categorias[1] = Categorias{
		Categoria:     "arma",
		Posibilidades: []string{"Pistola", "Cuchillo", "Machete", "Pala", "Bate", "Botella", "Tubo", "Cuerda"},
	}

	categorias[2] = Categorias{
		Categoria:     "motivo",
		Posibilidades: []string{"Venganza", "Celos", "Dinero", "Accidente", "Drogas", "Robo"},
	}

	categorias[3] = Categorias{
		Categoria:     "cuerpo",
		Posibilidades: []string{"Cabeza", "Pecho", "Abdomen", "Espalda", "Piernas", "Brazos"},
	}

	categorias[4] = Categorias{
		Categoria:     "lugar",
		Posibilidades: []string{"Sala", "Comedor", "Baño", "Terraza", "Cuarto", "Garage", "Patio", "Balcón", "Cocina"},
	}

	return categorias
}

// Determina si una solución es válida
// Sea k : tamano de restricciones, n : tamano de solv
// Coste algoritmico: O(kn)
func isRightSolution(solv []Categorias, rest [][]string) bool {
	mapaRest := make(map[string]string)
	mapaSoluciones := make(map[string]bool)

	// Maximo k iteraciones
	for _, restricciones := range rest {
		mapaRest[restricciones[0]] = restricciones[1]
	}

	// Maximo n iteraciones
	for _, soluciones := range solv {
		mapaSoluciones[soluciones.Posibilidades[0]] = true
	}

	// Maximo n iteraciones
	for solucion := range mapaSoluciones {
		current := mapaRest[solucion]

		if mapaSoluciones[current] {
			return false
		}
	}

	return true

}

// Se genera una solucion sin validar que sea valida
// Complejidad del algoritmo: O(n)
func GetSolution(posibilidades []Categorias) [][]Categorias {
	categorias := make([]Categorias, len(posibilidades))
	rst := make([]Categorias, len(posibilidades))

	for i := 0; i < len(categorias); i++ {
		index := rand.Intn(len(posibilidades[i].Posibilidades))

		current := Categorias{
			Categoria: posibilidades[i].Categoria,
		}

		categorias[i] = current
		categorias[i].Posibilidades = []string{posibilidades[i].Posibilidades[index]}

		rst[i] = current
		rst[i].Posibilidades = posibilidades[i].Posibilidades[:index+1]
		rst[i].Posibilidades = append(rst[i].Posibilidades, posibilidades[i].Posibilidades[index:]...)
	}

	return [][]Categorias{categorias, rst}
}

// Genera las restricciones aleatorias
// Complejidad del algoritmo: O(n)
func GenerateRest(cant int, posibilidades []Categorias) [][]string {
	rand.Seed(time.Now().UnixNano())

	rest := make([][]string, cant)

	for i := 0; i < cant; i++ { // Maximo de iteraciones n
		rest[i] = make([]string, 2)
		index := rand.Intn(len(posibilidades))

		lista1 := posibilidades[index]

		// Eliminamos el lista1 del indice para no repetirlo
		tmp := posibilidades[:index]
		tmp = append(tmp, posibilidades[index+1:]...)

		lista2 := tmp[rand.Intn(len(tmp))]

		rest[i][0] = lista1.Posibilidades[rand.Intn(len(lista1.Posibilidades))]
		rest[i][1] = lista2.Posibilidades[rand.Intn(len(lista2.Posibilidades))]
	}

	return rest
}

// Dada una restriccion la busca en la posible solucion
// Coste algoritmico: O(5)
func Find(solv []Categorias, rest []string) bool {
	count := 0

	for _, categoria := range solv { // Maximo 5 iteraciones
		if rest[0] == categoria.Posibilidades[0] || rest[1] == categoria.Posibilidades[0] {
			count++
		}

		if count >= 2 {
			return true
		}
	}

	return false
}
