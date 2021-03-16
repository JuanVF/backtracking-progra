import Cards from './cards'

// Genera las cartas
function CardsGenerator(cardsList){
    return cardsList.map((card, key)=><Cards key={key} title={card.title} cardsList={card.cardsList}/>)
}

function CardsContainer(props){
    let cardsList = []
    
    if (props.cardsList){
        cardsList = CardsGenerator(props.cardsList)
    }

    return (
        <div className="cards-container">
            {cardsList}
        </div>
    )
}

export default CardsContainer