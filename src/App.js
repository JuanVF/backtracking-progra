import './css/app.css';
import CardsContainer from './components/cardsContainer'

function App() {
    let cardsList = [
      {
        type : "right",
        text : "El/la mejor amigo(a)"
      },
      {
        type : "discarded",
        text : "El/la novio(a)"
      },
      {
        type : "default",
        text : "El/la vecino(a)"
      },
      {
        type : "selected",
        text : "El mensajero"
      },
      {
        type : "default",
        text : "El extraño"
      }
    ]

    let containers = [
      {
        title : "Sospechosos",
        cardsList : cardsList
      },
      {
        title : "Armas",
        cardsList : cardsList
      },
      {
        title : "Lugar",
        cardsList : cardsList
      },
      {
        title : "Motivo",
        cardsList : cardsList
      },
    ]

    return (
      <div>
        <CardsContainer cardsList={containers}/>
      </div>
    );
}

export default App
