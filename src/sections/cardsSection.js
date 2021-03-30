import {useState, useEffect} from 'react'

import CardsContainer from '../components/cardsContainer'
import cardsData from '../cards.json'

const wsurl = "ws://localhost:5000/api"

// Definimos el websocket
const ws = new WebSocket(wsurl)

function CardsSection(){
    const [cardsList, setCardsList] = useState(JSON.parse(JSON.stringify(cardsData)))
    const [loading, setLoading] = useState(false)
    
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

    useEffect(()=>{
        // Evento de conexion
        ws.onopen = event => {
            console.log("Conexion con el websocket del servidor iniciada...")

            let msg = {
                ID : 0
            }

            ws.send(JSON.stringify(msg))
        }

        // Evento de error
        ws.onerror = event => {
            console.log(event)
        }
    })

    useEffect(()=>{
        let wsActions = new Map([
            [0, function(msg){
                let solucion = JSON.parse(msg.json)
                let type = "selected"

                setCardsList(SetPosibilidades(solucion, type))
            }]
        ])

        // Evento de mensaje
        ws.onmessage = (event) => {
                setLoading(true)
                console.log(event.data)
                let msg = JSON.parse(event.data)

                let action = wsActions.get(msg.ID)

                if (action !== undefined) {
                    action(msg)
                }else{
                    console.log(msg)
                }

                setLoading(false)
            }
    }, [])

    return !loading ? (
        <div>
            <p className="cards-container-title">Cartas disponibles</p>
            <CardsContainer cardsList={cardsList}/>
            <p className="cards-container-title">Solución generada</p>
            <CardsContainer cardsList={[]} type="default"/>
        </div>) : (<div></div>)
}

export default CardsSection