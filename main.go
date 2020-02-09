package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var SIZE = 6

type Stone byte

const (
	Blank Stone = '*'
	Black Stone = 'O'
	White Stone = 'X'
)

type Player struct {
}

type Board [6][6]Point
type Line []*Point
type Game struct {
	board Board
	lines []Line
	white Player
	black Player
	count int
}

func (game *Game) init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			//if rand.Intn(2) == 0 {
			//	game.board[i][j] = Point {
			//		y: i,
			//		x: j,
			//		stone: Black,
			//	}
			//} else {
			//	game.board[i][j] = Point {
			//		y: i,
			//		x: j,
			//		stone: White,
			//	}
			//}
			game.board[i][j] = Point{
				y:     i,
				x:     j,
				stone: Blank,
			}
		}
	}
	//game.lines = game.board.createLines()
}

func (board *Board) print() {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			fmt.Printf("%3s", string(board[i][j].stone))
		}
		fmt.Println()
	}
}

type Point struct {
	x     int
	y     int
	stone Stone
}

func (line Line) hasWinner() (bool, []*Point) {
	points := []*Point{line[0]}

	for _, point := range line[1:] {
		if len(points) == 5 && points[len(points)-1].stone != point.stone {
			return true, points
		}
		//fmt.Printf("line %d, points: %d\n", len(line), len(points))
		if point.stone == Blank || points[len(points)-1].stone != point.stone {
			points = points[:0]
		}
		points = append(points, point)
	}

	if len(points) == 5 {
		return true, points
	}
	return false, nil
}

//func (line Line) has3x3() (bool, []*Point) {
//
//}

func (game *Game) findPattern() bool {
	for _, line := range game.lines {
		result, points := line.hasWinner()
		if result {
			for _, point := range points {
				fmt.Printf("(%d, %d)  ", point.y, point.x)
			}
			fmt.Printf("\n%s가 승리하였습니다.\n", string(points[0].stone))
			return true
		}
	}
	return false
}

func (board *Board) findLines(y, x int) []Line {
	stone := board[y][x].stone
	var lines []Line
	var start int

	for i := y - 1; i > -1; i-- {
		if board[i][x].stone != stone {
			break
		}
		start = i
	}
	var line Line
	for i := start; i < SIZE; i++ {
		line = append(line, &board[i][x])
	}
	lines = append(lines, line)

	for i := x - 1; i > -1; i-- {
		if board[y][i].stone != stone {
			break
		}
		start = i
	}
	line = Line{}
	for i := start; i < SIZE; i++ {
		line = append(line, &board[y][i])
	}
	lines = append(lines, line)

	for i := 0; y-i > -1 && x-i > -1; i++ {
		if board[y-i][x-i].stone != stone {
			break
		}
		start = i
	}
	line = Line{}
	for i := -start; y+i < SIZE && x+i < SIZE; i++ {
		line = append(line, &board[y+i][x+i])
	}
	lines = append(lines, line)

	for i := 0; y-i > -1 && x+i < SIZE; i++ {
		if board[y-i][x+i].stone != stone {
			break
		}
		start = i
	}
	line = Line{}
	for i := -start; y+i < SIZE && x-i > -1; i++ {
		line = append(line, &board[y+i][x-i])
	}
	lines = append(lines, line)

	return lines
}

func (game *Game) putStone(y, x int) bool {
	stone := Black

	game.board[y][x].stone = stone
	lines := game.board.findLines(y, x)
	for _, line := range lines {
		result, points := line.hasWinner()
		if result {
			for _, point := range points {
				fmt.Printf("(%d, %d)  ", point.y, point.x)
			}
			fmt.Printf("\n%s가 승리하였습니다.\n", string(points[0].stone))
			return true
		}
	}
	return false
}

func main() {
	for true {
		var game Game
		game.init()
		fmt.Println("Start Game")
		var input string
		for true {
			fmt.Print("Enter Y: ")
			_, _ = fmt.Scanln(&input)
			y, _ := strconv.Atoi(input)
			fmt.Print("Enter X: ")
			_, _ = fmt.Scanln(&input)
			x, _ := strconv.Atoi(input)

			result := game.putStone(y, x)
			//if stone == Black {
			//	stone = White
			//} else {
			//	stone = Black
			//}
			game.board.print()
			if result {
				break
			}
		}
	}
}
