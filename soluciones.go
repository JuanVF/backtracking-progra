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
	for _, restriccion := range rest { // Maximo k iteraciones
		if Find(solv, restriccion) { // Maximo n iteraciones
			return false
		}
	}

	return true
}

// Se genera una solucion sin validar que sea valida
// Complejidad del algoritmo: O(n)
func GetSolution(posibilidades []Categorias, rest [][]string) []Categorias {
	rand.Seed(time.Now().UnixNano())
	solv := make([]Categorias, len(posibilidades))

	posibilidades = GetValidPosibilities(posibilidades, rest)

	for i := 0; i < len(posibilidades); i++ { // Maximo n iteraciones
		length := len(posibilidades[i].Posibilidades)
		randPos := posibilidades[i].Posibilidades[rand.Intn(length)]

		solv[i] = Categorias{Categoria: posibilidades[i].Categoria}
		solv[i].Posibilidades = append(solv[i].Posibilidades, randPos)
	}

	return solv
}

// Dadas unas restricciones retorna posibilidades validas para generar una solucion
// Complejidad del algoritmo: O(kn)
func GetValidPosibilities(categorias []Categorias, rest [][]string) []Categorias {
	tmp := make([]Categorias, 0)
	restMap := make(map[string]bool)
	added := make(map[string]bool)

	for i := 0; i < len(categorias); i++ { // Maximo 5 iteraciones
		tmp = append(tmp, Categorias{Categoria: categorias[i].Categoria})

		// Este codigo necesita refactor
		for _, posibilidad := range categorias[i].Posibilidades { // Maximo n iteraciones
			for _, restriccion := range rest { // Maximo k iteraciones
				if restMap[posibilidad] || added[posibilidad] {
					continue
				}

				switch posibilidad {
				case restriccion[0]:
					restMap[restriccion[1]] = true
					break
				case restriccion[1]:
					restMap[restriccion[0]] = true
					break
				}

				added[posibilidad] = true
				tmp[i].Posibilidades = append(tmp[i].Posibilidades, posibilidad)
			}
		}
	}

	return tmp
}

// Genera las restricciones aleatorias
// Complejidad del algoritmo: O(n)
func GenerateRest(cant int, posibilidades []Categorias) [][]string {
	rand.Seed(time.Now().UnixNano())

	rest := make([][]string, cant)

	for i := 0; i < cant; i++ { // Maximo de iteraciones n
		rest[i] = make([]string, 2)
		index := rand.Intn(len(posibilidades))

		l1 := posibilidades[index]

		// Eliminamos el l1 del indice para no repetirlo
		tmp := posibilidades[:index]
		tmp = append(tmp, posibilidades[index+1:]...)

		l2 := tmp[rand.Intn(len(tmp))]

		rest[i][0] = l1.Posibilidades[rand.Intn(len(l1.Posibilidades))]
		rest[i][1] = l2.Posibilidades[rand.Intn(len(l2.Posibilidades))]
	}

	return rest
}

// Dada una restriccion la busca en la posible solucion
// Coste algoritmico: O(n)
func Find(solv []Categorias, rest []string) bool {
	count := 0

	for _, categoria := range solv { // Maximo 5 iteraciones
		for _, posibilidad := range categoria.Posibilidades { // Maximo n iteraciones
			if rest[0] == posibilidad || rest[1] == posibilidad {
				count++
			}

			if count >= 2 {
				return true
			}
		}
	}

	return false
}
