package main

import (
	"fmt"

	board "github.com/PS-Wizard/goConnect/game"
)

// TODO:
// - Func to check win

func main() {
	b := board.NewBoard()
	var color, idx int
	for {
		b.DrawBoard()
		fmt.Print("Enter shit to place at: ")
		fmt.Scanf("%d,%d", &color, &idx)
		r, c := b.Drop(color, idx-1)
		if won := b.CheckWin(r, c, color); won {
			fmt.Println("GAME OVER _______ ", color, "WON")
			return
		}
	}
}
