package invasion

import "fmt"

type alienNames []string

func initAlienNames(numAliens int) []string {
	names := make(alienNames, numAliens)
	for i := 0; i < numAliens; i++ {
		alienName := fmt.Sprintf("alien %d", i+1)
		names[i] = alienName
	}
	return names
}

type alienMoves map[string]int // moves indexed by alien name

func initAlienMoves(alienNames []string) alienMoves {
	moves := make(alienMoves)
	for _, alienName := range alienNames {
		moves[alienName] = 0
	}
	return moves
}
