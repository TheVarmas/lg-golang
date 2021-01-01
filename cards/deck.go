package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

type deck []string

func createDeck() deck {
    values := [] string {
        "Ace",
        "2",
        "3",
        // "4",
        // "5",
        // "6",
        // "7",
        // "8",
        // "9",
        // "10",
        // "Jack",
        // "Queen",
        // "King",
    }

    suits := []string {
        "Diamonds",
        "Hearts",
        "Clubs",
        "Spades",
    }
    
    var someDeck deck


    for _, value := range values{
        for _, suit := range suits{
            someDeck = append(someDeck, value + " of " + suit)
        }
    }

    return someDeck
}

func (d deck) showCards() {
    for _, card := range d {
        fmt.Println(card)
    }
}

func (d deck) deal (handSize int) (deck, deck){
    return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
    return strings.Join([]string(d), ",")
}

func (d deck) saveToFile (fileName string) {
    ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func toDeck (ds string) deck {
    return deck(strings.Split(ds, ","))
    
}

// First attempt that has gone wrong!
// [TODO] Retry with mutiple lines and reverse saveToFile

// func readFromFileBad(nameFile string) deck {
//     return deck(string(ioutil.ReadFile(nameFile)))
// }

// 2nd attempt
func readFromFile (fileName string) deck {
    bs, err := ioutil.ReadFile(fileName)

    if err != nil {
        fmt.Println("Error: ", err)
    }

    return toDeck(string(bs))
}