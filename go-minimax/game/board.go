package board

import "fmt"

type Game struct {
	board [6][7]int
	Blue  int
	Red   int
}

func NewBoard() *Game {
	n := new(Game)
	n.Blue = 1
	n.Red = 2
	return n
}

func (b *Game) Drop(color, col int) (int, int) {
	for r := 5; r >= 0; r-- {
		if b.board[r][col] == 0 {
			b.board[r][col] = color
			return r, col
		}
	}
	return 255, 255
}

func (b *Game) DrawBoard() {
	fmt.Println("[1 2 3 4 5 6 7]")
	fmt.Println("[- - - - - - -]")
	for _, val := range b.board {
		fmt.Println(val)
	}
}

func (b *Game) checkConsec(row, col, drow, dcol, color int) int {
	counter := 0
	for {
		row += drow
		col += dcol
		if row < 0 || row >= 6 || col < 0 || col >= 7 || b.board[row][col] != color {
			break
		}
		counter++
	}
	return counter
}

func (b *Game) CheckWin(row, col, color int) bool {
	directions := [][2]int{
		{0, 1},  // Horizontal
		{1, 0},  // Vertical
		{1, 1},  // Diagonal \
		{-1, 1}, // Diagonal /
	}

	for _, dir := range directions {
		drow, dcol := dir[0], dir[1]
		count := 1
		count += b.checkConsec(row, col, drow, dcol, color)
		count += b.checkConsec(row, col, -drow, -dcol, color)
		if count >= 4 {
			return true
		}
	}
	return false
}
