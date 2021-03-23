package main

import (
	"fmt"
)

func main() {
	sospechoso := []string{"El/La mejor amigo(a)", "El/la novio(a)", "El/la vecino(a)", "El mensajero", "El extraño", "El/la hermanastro(a)", "El/la colega de trabajo"}
	arma := []string{"Pistola", "Cuchillo", "Machete", "Pala", "Bate", "Botella", "Tubo", "Cuerda"}
	motivo := []string{"Venganza", "Celos", "Dinero", "Accidente", "Drogas", "Robo"}
	cuerpo := []string{"Cabeza", "Pecho", "Abdomen", "Espalda", "Piernas", "Brazos"}
	lugar := []string{"Sala", "Comedor", "Baño", "Terraza", "Cuarto", "Garage", "Patio", "Balcón", "Cocina"}

	posibilidades := [][]string{sospechoso, arma, motivo, cuerpo, lugar}
	solucion := []string{"El/la colega de trabajo", "Cuerda", "Robo", "Brazos", "Cocina"}
	encontrada := []string{}

	iteraciones, _ := FuerzaBruta(posibilidades, solucion, encontrada)

	fmt.Printf("Iteraciones en fuerza bruta: %d\n", iteraciones)
}
