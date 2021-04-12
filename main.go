package main

import (
	"log"
	"net/http"

	"github.com/JuanVF/gogame-server/sockets"
	"github.com/gorilla/websocket"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Variable para generar la coneccion de websocket
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Handler para mantener la coneccion por websocket
func HandlerUsers(w http.ResponseWriter, req *http.Request) {
	ws, err := Upgrader.Upgrade(w, req, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	// Agregamos el usuario a la lista de usuarios conectados
	sockets.GetInstance().AddConn(ws)

	// Esto funciona basicamente como un listener del usuario
	for {
		var msg sockets.Message

		// Esto se queda esperando a un mensaje del usuario
		err := ws.ReadJSON(&msg)

		// Si hay error lo desconectamos
		if err != nil {
			log.Printf("Message error: %v", err)

			sockets.GetInstance().RemoveConn(ws)
			return
		}

		// Hacemos print de lo que el usuario envio
		log.Printf("User sended: %v\n", msg)

		// Ejecutamos una accion en base a lo que el usuario envio
		sockets.GetInstance().GetAction(msg.ID)(ws, msg)
	}
}

func main() {
	/*solucion := GetSolution(GetCategorias())
	rest := GenerateRest(5, solucion[1])

	encontrada := make([]Categorias, 0)
	mensajes := make([]sockets.Message, 0)
	eliminadas := make(map[string]bool)

	//FuerzaBruta(GetCategorias(), solucion, encontrada, &eliminadas, &mensajes)

	Backtracking(GetCategorias(), solucion[0], encontrada, rest, &eliminadas, &mensajes)

	fmt.Println(solucion)
	fmt.Println(eliminadas)
	fmt.Println(rest)*/
	GenerarGraficas()
}

func GenerarGraficas() {
	records := make(map[string]plotter.XYs)
	recordsTime := make(map[string]plotter.XYs)

	for i := 1; i <= 10; i++ { // Restricciones
		// Hacemos una media para que sea mas correcto
		var promIteracionesfb float64 = 0
		var promTiempofb float64 = 0
		var promIteraciones float64 = 0
		var promTiempo float64 = 0

		solucion := GetSolution(GetCategorias())

		solv := []Categorias{
			{
				Categoria:     "sospechoso",
				Posibilidades: []string{"El/la colega de trabajo"},
			},
			{
				Categoria:     "arma",
				Posibilidades: []string{"Cuerda"},
			},
			{
				Categoria:     "motivo",
				Posibilidades: []string{"Robo"},
			},
			{
				Categoria:     "cuerpo",
				Posibilidades: []string{"Brazos"},
			},
			{
				Categoria:     "lugar",
				Posibilidades: []string{"Cocina"},
			},
		}

		rest := GenerateRest(i, solucion[1])

		rep := 50

		for j := 0; j < rep; j++ { // Hacemos una media
			encontrada := make([]Categorias, 0)
			mensajes := make([]sockets.Message, 0)
			eliminadas := make(map[string]bool)

			// Medimos el tiempo
			initTime := GetCurrentTime()

			iteraciones, _ := Backtracking(GetCategorias(), solv, encontrada, rest, &eliminadas, &mensajes)

			time := GetCurrentTime() - initTime

			promTiempo += float64(time)

			initTime = GetCurrentTime()

			iteracionesfb, _ := FuerzaBruta(GetCategorias(), solv, encontrada, &eliminadas, &mensajes)

			time = GetCurrentTime() - initTime

			promTiempofb += float64(time)
			promIteraciones += float64(iteraciones)
			promIteracionesfb += float64(iteracionesfb)
		}

		promIteraciones /= float64(rep)
		promIteracionesfb /= float64(rep)
		promTiempo /= float64(rep)
		promTiempofb /= float64(rep)

		recordsTime["Backtracking"] = append(recordsTime["Backtracking"], plotter.XY{
			X: float64(i),
			Y: promTiempo,
		})

		recordsTime["FuerzaBruta"] = append(recordsTime["FuerzaBruta"], plotter.XY{
			X: float64(i),
			Y: promTiempofb,
		})

		records["Backtracking"] = append(records["Backtracking"], plotter.XY{
			X: float64(i),
			Y: promIteraciones,
		})

		records["FuerzaBruta"] = append(records["FuerzaBruta"], plotter.XY{
			X: float64(i),
			Y: promIteracionesfb,
		})
	}

	p, err := plot.New()
	ptime, err := plot.New()

	if err != nil {
		return
	}

	p.Title.Text = "Rendimiento segun restricciones"
	p.X.Label.Text = "Restricciones"
	p.Y.Label.Text = "Iteraciones"

	err = plotutil.AddLinePoints(p,
		"Backtracking", records["Backtracking"],
		"FuerzaBruta", records["FuerzaBruta"])

	if err != nil {
		return
	}

	if err := p.Save(7*vg.Inch, 7*vg.Inch, "iteraciones.png"); err != nil {
		panic(err)
	}

	ptime.Title.Text = "Rendimiento segun Tiempo"
	ptime.X.Label.Text = "Restricciones"
	ptime.Y.Label.Text = "Tiempo"

	err = plotutil.AddLinePoints(ptime,
		"Backtracking", recordsTime["Backtracking"],
		"FuerzaBruta", recordsTime["FuerzaBruta"])

	if err != nil {
		return
	}

	if err := ptime.Save(7*vg.Inch, 7*vg.Inch, "tiempo.png"); err != nil {
		panic(err)
	}
}
