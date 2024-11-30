package main

import (
	"fmt"

	board "github.com/PS-Wizard/goConnect/game"
)

func main() {
	b := board.NewBoard()
	var color, idx int
	for {
		fmt.Print("Enter shit to place at: ")
		fmt.Scanf("%d,%d", &color, &idx)
		r, c := b.Drop(color, idx-1)
		fmt.Println("Eval says", b.Evaluate(r,c))
		if won := b.CheckWin(r, c, color); won {
			fmt.Println("GAME OVER _______ ", color, "WON")
			return
		}
		b.DrawBoard()
	}
}
