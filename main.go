package main

import "fmt"

func main() {

	cards := newDeck()

	cards.saveToFile("playing-cards.txt")

	fmt.Print(readFromFile("playing-cardsss.txt").toString())
}
