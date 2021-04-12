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
	records := make(map[string]plotter.XYs)

	for i := 1; i <= 20; i++ { // Restricciones
		// Hacemos una media para que sea mas correcto
		var promIteracionesfb float64 = 0
		var promIteraciones float64 = 0

		for j := 0; j < 100; j++ { // Hacemos una media
			rest := GenerateRest(i, GetCategorias())
			solucion := GetSolution(GetCategorias(), rest)
			encontrada := make([]Categorias, 0)
			mensajes := make([]sockets.Message, 0)
			eliminadas := make(map[string]bool)

			// Medimos el tiempo
			//initTime := GetCurrentTime()

			iteracionesfb, _ := FuerzaBruta(GetCategorias(), solucion, encontrada, &mensajes, rest)

			iteraciones, _ := Backtracking(GetCategorias(), solucion, encontrada, rest, &eliminadas, &mensajes)

			//time := GetCurrentTime() - initTime

			promIteraciones += float64(iteraciones)
			promIteracionesfb += float64(iteracionesfb)
			//promTiempo += float64(time)
		}

		promIteraciones /= 100
		promIteracionesfb /= 100
		//promTiempo /= 100

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

	if err := p.Save(7*vg.Inch, 7*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
