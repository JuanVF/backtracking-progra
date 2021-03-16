import CardsText from './cardsText'

// Esto genera las cartas disponibles
// Props: cardsList - lista de cartas
//        title - string
function GenerateCards(cardsList){
    return cardsList.map((card, key)=> <CardsText key={key} type={card.type} text={card.text}/>)
}

function Cards(props){
    let cardsList = []

    if (props.cardsList){
        cardsList = GenerateCards(props.cardsList)
    }

    return (
        <div className="card">
            <p className="card-title">{props.title}</p>
            {cardsList}
        </div>
    )
}

export default Cards