import './css/app.css';
import Cards from './components/cards'

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

    return (
      <div>
        <Cards title="Sospechosos" cardsList={cardsList}/>
      </div>
    );
}

export default App
