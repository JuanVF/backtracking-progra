// Hashmap con las clases disponibles
let cardTextType = new Map([
    ["discarded", "cardText-discarded"],
    ["selected", "cardText-selected"],
    ["default", "cardText-default"],
    ["answer", "cardText-answer"],
    ["right", "cardText-right"]
])

// Este componente corresponde al texto de las cartas
// props: type - string
//        text - string
function CardsText(props){   
    let hasCardType = cardTextType.has(props.type)
    let cardType = cardTextType["default"]

    if (hasCardType) {
        cardType = cardTextType.get(props.type)
    }

    return (
        <div className={cardType}>
            <p>{props.text}</p>
        </div>
    )
}

export default CardsText