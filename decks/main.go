package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()

	cards.print()
}

func print() {
	fmt.Println("test")
}
