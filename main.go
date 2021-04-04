package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"time"

	"github.com/JuanVF/gogame-server/sockets"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandlerUsers(w http.ResponseWriter, req *http.Request) {
	ws, err := Upgrader.Upgrade(w, req, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	sockets.GetInstance().AddConn(ws)

	for {
		var msg sockets.Message

		err := ws.ReadJSON(&msg)

		if err != nil {
			log.Printf("Message error: %v", err)

			sockets.GetInstance().RemoveConn(ws)
			break
		}

		log.Printf("User sended: %v\n", msg)

		sockets.GetInstance().GetAction(msg.ID)(ws)
	}
}

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
	eliminadas := make(map[string]bool)
	fmt.Println("Solución:  ", solv)
	iteraciones1, _ := FuerzaBruta(posibilidades, solv, encontrada)
	iteraciones2, _, eliminadas := backtracking(posibilidades, solv, encontrada, rest, eliminadas)

	fmt.Printf("Iteraciones en fuerza bruta: %d\n", iteraciones1)
	fmt.Printf("Iteraciones en bacttracking: %d\n", iteraciones2)
}

func FuerzaBruta(posibilidades [][]string, solucion, encontrada []string) (int, bool) {

	// Caso base
	if len(posibilidades) == 0 {
		if Equals(solucion, encontrada) {
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

func select_elim(solv []string, try []string) string {
	rand.Seed(time.Now().UnixNano())
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

func backtracking(posibilidades [][]string, solucion, encontrada []string, rest [][]string, eliminadas map[string]bool) (int, bool, map[string]bool) {

	// Caso base
	if len(posibilidades) == 0 {
		if Equals(solucion, encontrada) {
			fmt.Println("try: ", encontrada, "		FINAL")
			return 1, true, eliminadas
		}
		i := select_elim(solucion, encontrada)
		eliminadas[i] = true
		fmt.Println("try: ", encontrada, "		eliminado:", i)
		return 1, false, eliminadas
	}

	amount := 0
	// Probamos cada array
	for _, posibilidad := range posibilidades[0] {
		generada := append(encontrada, posibilidad)
		if !gen_validation(encontrada, eliminadas) {
			continue
		}
		if sol_validation(generada, rest) {
			iteraciones, finded, eliminadas := backtracking(posibilidades[1:], solucion, generada, rest, eliminadas)
			amount += iteraciones
			if finded {
				return amount, true, eliminadas
			}
		}
	}
	return amount, false, eliminadas
}

func gen_validation(encontrada []string, eliminadas map[string]bool) bool {

	for _, x := range encontrada {
		if eliminadas[x] {
			return false
		}
	}
	return true
}

func sol_validation(solv []string, rest [][]string) bool {
	for i := 0; i < len(rest); i++ {
		if find(solv, rest[i][0]) || find(solv, rest[i][1]) {
			return false
		}
	}
	return true
}

func solution(posibilidades [][]string, rest [][]string) []string {
	aux := false
	solv := make([]string, len(posibilidades))
	for !aux {
		solv = try_select(posibilidades)
		aux = sol_validation(solv, rest)
	}
	return solv
}

func rest_generator(cant int, posibilidades [][]string) [][]string {

	rand.Seed(time.Now().UnixNano())

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
	rand.Seed(time.Now().UnixNano())
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
