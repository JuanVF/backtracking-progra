import {useState, useEffect} from 'react'

import CardsContainer from '../components/cardsContainer'
import cardsData from '../cards.json'

const wsurl = "ws://localhost:5000/api"

// Definimos el websocket
const ws = new WebSocket(wsurl)
const queue = []
let solution = []

function CardsSection(){
    // Hooks de la app
    const [cardsList, setCardsList] = useState(JSON.parse(JSON.stringify(cardsData)))
    const [loading, setLoading] = useState(false)

    // Acciones permitidas para la comunicacion entre websockets
    const wsActions = new Map([
        [0, function(msg){ // Funcion para una solucion incorrecta
            let solucion = JSON.parse(msg.json)
            let type = "selected"
    
            setCardsList(SetPosibilidades(solucion, type))
        }],
        [1, function(msg){ // Funcion para una solucion correcta
            let solucion = JSON.parse(msg.json)
            let type = "answer"
    
            setCardsList(SetPosibilidades(solucion, type))
            console.log(SetPosibilidades(solucion, type))
        }],
        [3, function(msg){ // Funcion para la solucion
            let solucion = JSON.parse(msg.json)

            solution = ParseSolution(solucion)
            console.log(solution)
        }]
    ])

    // Se encarga de ejecutar la cola de mensajes que proviene del server
    const Listener = () => {
        if (queue[0] !== undefined){

            let msg = queue.shift()
            let action = wsActions.get(msg.ID)

            if (action !== undefined) {
                action(msg)
            }else{
                console.log(msg)
            }
        }
    }

    setInterval(Listener, 100)

    // Effect para activar la conexion con el websocket server
    useEffect(()=>{
        // Evento de conexion
        ws.onopen = event => {
            console.log("Conexion con el websocket del servidor iniciada...")

            let msg = {
                ID : 1,
                number : 5
            }

            ws.send(JSON.stringify(msg))
        }

        // Evento de error
        ws.onerror = event => {
            console.log(event)
        }
    })

    // Effect para recibir mensajes
    useEffect(()=>{
        // Evento de mensaje
        ws.onmessage = (event) => {
            let msg = JSON.parse(event.data)

            queue.push(msg)
        }

        setInterval(()=>{
            setLoading(true)
            setLoading(false)
        }, 300)
    }, [])

    return !loading ? (
        <div>
            <p className="cards-container-title">Cartas disponibles</p>
            <CardsContainer cardsList={cardsList}/>
            <p className="cards-container-title">Solución generada</p>
            <CardsContainer cardsList={solution} type="default"/>
            <p className="cards-container-title">Estadisticas</p>
        </div>) : (<div></div>)
}
    
// Dada una posible solucion actualizamos el estado de las cartas
// Nota: se asume que las cartas van ordenadas por categoria
// al igual que la solucion que se pasa por parametro
// esto por eficiencia
const SetPosibilidades = (solucion, type)=>{
    let cards = JSON.parse(JSON.stringify(cardsData))

    for (let i = 0; i < cards.length; i++){
        SetPosibilidad(cards[i].posibilidades, solucion[i].Posibilidades[0], type)
    }

    return cards
}


// Permite alterar el tipo de texto de alguna carta
const SetPosibilidad = (posibilidades, posibilidad, type)=>{
    let index = posibilidades.findIndex(element => {
        return element.text.toLowerCase() === posibilidad.toLowerCase()
    })

    // La solucion existe en el array
    if (index >= 0){
        posibilidades[index].type = type
    }
}

// Se encarga de parsear la solucion
const ParseSolution = (solucion) => {
    let type = "default"

    for (let i = 0; i < solucion.length; i++){
        solucion[i].posibilidades = [{
            text : solucion[i].Posibilidades[0], 
            type : type
        }]
    }

    return solucion
}

export default CardsSection