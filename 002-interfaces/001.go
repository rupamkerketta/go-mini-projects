package main

import "fmt"

type bot interface {
    getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main() {
    fmt.Println("Interfaces !!!")

    eb := englishBot{}
    sb := spanishBot{}

    printGreeting(eb)
    printGreeting(sb)
}

func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}

func (e englishBot) getGreeting() string {
    return "Hello, there!!!"
}

func (s spanishBot) getGreeting() string {
    return "Hola!!"
}
