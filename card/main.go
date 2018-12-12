package main 

// import "fmt"

func main() {
	cards := deck{"Ace of spades", newCard()}
	cards = append(cards, "Six of spades")
	// fmt.Println(cards)

	cards.print()
}



func newCard() string {

	return "Five of Diamonds"
}