package main

import "fmt"

func main(){
    //cards := deck{"Ace of Diamonds", "Ace of Spades", newCard()}
    //cards = append(cards, "Six of Spades")

    //cards.print()
    fmt.Println("Dealing Cards!!")
    cards := newDeck()
    cards.shuffle()
    cards.print()
    //fmt.Println(cards.toString())
    //cards.saveToFile("my_cards")
    //cards := newDeckFromFile("my_cards")
    //cards.print()
}

func newCard() string {
    return "Five of Diamonds"
}
