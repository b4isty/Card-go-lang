package main 

import "fmt"

func main() {
	cards := newDeck()
	// cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("------------------")
	remainingCards.print()

	greeting := "Hi there!"
	fmt.Println([]byte(greeting))

}
     