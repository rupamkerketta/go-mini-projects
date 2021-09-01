package main

import (
	"fmt"
    "time"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
    cards := deck{}

    cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
    cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

    for _, cardSuit := range cardSuits {
        for _, cardValue := range cardValues {
            cards = append(cards, fmt.Sprintf("%v of %v", cardValue, cardSuit))
        }
    }

    return cards
}

func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}

func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}

// 'toString' is gonna take a deck and return a complete string representation of it.
func (d deck) toString() string {
    return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
    bs, err := ioutil.ReadFile(filename)
    if err != nil {
        // Option #1 : log the error and return the call to newDeck()
        // Option #2 : log the error and entirely quit the program
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    //fmt.Println(bs)
    s := strings.Split(string(bs), ",")
    return deck(s)
}

func (d deck) shuffle() {
    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)

    for i := range d {
        newPosition := r.Intn(len(d) - 1)

        d[i], d[newPosition] = d[newPosition], d[i]
    }
}
