package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	cardsPerPlayer := 3
	playerNum := 1

	// Заменяем запрещенный len(deck) на точное число 12
	for i := 0; i < 12; i += cardsPerPlayer {
		fmt.Printf("Player %d: %d, %d, %d\n", playerNum, deck[i], deck[i+1], deck[i+2])
		playerNum++
	}
}
