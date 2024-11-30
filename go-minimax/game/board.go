package board

import "fmt"

var directions = [][2]int{
	{0, 1},  // Horizontal
	{1, 0},  // Vertical
	{1, 1},  // Diagonal \
	{-1, 1}, // Diagonal /
}

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
	if col < 0 || col >= 7 {
		return 255, 255
	}
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

func (b *Game) GetValidMoves() []int {
	moves := []int{}
	for c := 0; c < 7; c++ {
		if b.board[0][c] == 0 { // Check if column is not full
			moves = append(moves, c)
		}
	}
	return moves
}

func (b *Game) Undo(row, col int) {
	b.board[row][col] = 0
}

// Improved evaluation based on consecutive counts and blocking
func evaluatePlayer(count []int) int {
    fmt.Println(count)
	score := 0
	for _, val := range count {
		switch val {
		case 4: // Immediate win
			score += 10000
		case 3: // 3 consecutive pieces (dangerous)
			score += 1000
		case 2: // 2 consecutive pieces
			score += 500
		case 1: // 1 consecutive piece (building)
			score += 100
		case 0: // No consecutive pieces
			score += 10
		}
	}
	return score
}

func (b *Game) Evaluate(row, col int) int {
	blueMoves, redMoves := []int{}, []int{}
	// Blue moves evaluation
	for _, dir := range directions {
		drow, dcol := dir[0], dir[1]
		count := 1
		count += b.checkConsec(row, col, drow, dcol, 1)
		count += b.checkConsec(row, col, -drow, -dcol, 1)
		blueMoves = append(blueMoves, count)
	}
	// Red moves evaluation
	for _, dir := range directions {
		drow, dcol := dir[0], dir[1]
		count := 1
		count += b.checkConsec(row, col, drow, dcol, 2)
		count += b.checkConsec(row, col, -drow, -dcol, 2)
		redMoves = append(redMoves, count)
	}
	// Evaluate scores for both players
	blueScore := evaluatePlayer(blueMoves)
	redScore := evaluatePlayer(redMoves)

	// Debug prints for clarity
	fmt.Println("Blue, ", blueScore, "Red", redScore)

	return blueScore - redScore
}

