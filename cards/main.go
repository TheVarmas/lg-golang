package main

//import "fmt"

func main(){
    newDeck := createDeck()
    //fmt.Println("\n   This is a deck: \n")
    newDeck.showCards()
    hand, remainingDeck := newDeck.deal(26)
    //fmt.Println("\n   This is the hand: \n")
    hand.showCards()
    //fmt.Println("\n   This is the remaining deck: \n")
    remainingDeck.showCards()

    newDeck.saveToFile("Deck_1")

//fmt.Println("\n Reading from a file \n")

    readFromFile("My_Deck").showCards()
}
