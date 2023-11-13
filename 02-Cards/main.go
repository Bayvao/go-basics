package main

func main() {

	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	//card := newCard()

	cards := newDeck()
	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")

	cards.shuffle()
	cards.print()

}
