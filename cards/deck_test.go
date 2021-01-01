package main

import (
    "testing"
    "strings"
)    

// func testToString(t *testing.T){
//     //testDeck := []string {
//     //    "Ace of Diamonds",
//     //    "Ace of Hearts",
//         "Ace of Clubs",
//         "Ace of Spades",
//         "2 of Diamonds",
//         "2 of Hearts",
//         "2 of Clubs",
//         "2 of Spades",
//         "3 of Diamonds",
//         "3 of Hearts",
//         "3 of Clubs",
//         "3 of Spades",
//     }

//     //testToString := testDeck    
    

// }

func TestCreateDeck(t *testing.T) {
    d := createDeck()
    
    if len(d) != 12 {
        t.Errorf("Expected 12 cards found %v cards instead", len(d))
    }

    testDeck := []string {
        "Ace of Diamonds",
        "Ace of Hearts",
        "Ace of Clubs",
        "Ace of Spades",
        "2 of Diamonds",
        "2 of Hearts",
        "2 of Clubs",
        "2 of Spades",
        "3 of Diamonds",
        "3 of Hearts",
        "3 of Clubs",
        "3 of Spades",
    }    

    if strings.Join(d, ",") != strings.Join(testDeck, ",") {
        t.Errorf("Something is wrong!")
    }
}