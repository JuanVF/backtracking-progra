package main

import (
	"encoding/json"
	"reflect"

	"github.com/JuanVF/gogame-server/sockets"
)

type Categorias struct {
	Categoria     string   `json:"Categoria"`
	Posibilidades []string `json:"Posibilidades"`
}

func GetCategorias() []Categorias {
	categorias := make([]Categorias, 5)

	categorias[0] = Categorias{
		Categoria:     "sospechoso",
		Posibilidades: []string{"El/La mejor amigo(a)", "El/la novio(a)", "El/la vecino(a)", "El mensajero", "El extraño", "El/la hermanastro(a)", "El/la colega de trabajo"},
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

func FuerzaBruta2(categorias, solucion, encontrada []Categorias, mensajes *[]sockets.Message) (int, bool) {
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

	// Probamos cada array
	for _, posibilidad := range categorias[0].Posibilidades {
		generada := append(encontrada, Categorias{
			Categoria:     categorias[0].Categoria,
			Posibilidades: []string{posibilidad},
		})

		iteraciones, finded := FuerzaBruta2(categorias[1:], solucion, generada, mensajes)
		amount += iteraciones

		if finded {
			return amount, true
		}
	}

	return amount, false
}
