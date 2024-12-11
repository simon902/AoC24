package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Field rune
type Game [][]Field

const (
	Empty Field = '.'
	Wall  Field = '#'
	Guard Field = '^'
)

type Orientation = int

const (
	Up Orientation = iota
	Right
	Down
	Left
)

type Coord struct {
	R int
	C int
}

func ReadFile() Game {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	game := Game{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		game = append(game, []Field(line))
	}

	return game
}

func FindGuardPos(game Game) (int, int) {
	for i, row := range game {
		for j, field := range row {
			if field == Guard {
				return i, j
			}
		}
	}

	return -1, -1
}

func IsValid(game Game, row int, col int) bool {
	return 0 <= row && row < len(game[0]) && 0 <= col && col < len(game)
}

func MoveForward(row int, col int, orientation Orientation) (int, int) {
	switch orientation {
	case Up:
		return row - 1, col
	case Right:
		return row, col + 1
	case Down:
		return row + 1, col
	case Left:
		return row, col - 1
	default:
		return -1, -1
	}
}

func TurnRight(orientation Orientation) Orientation {
	switch orientation {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		return Up
	}
}

func MakeNextMove(game Game, g_row int, g_col int, orientation Orientation) (int, int, Orientation) {

	next_row, next_col := MoveForward(g_row, g_col, orientation)

	if !IsValid(game, next_row, next_col) {
		return -1, -1, Up
	}

	if game[next_row][next_col] == Wall {
		next_orientation := TurnRight(orientation)
		next_row, next_col = MoveForward(g_row, g_col, next_orientation)
		return next_row, next_col, next_orientation
	}

	return next_row, next_col, orientation
}

func CountDistinctEntries(arr [][]int) int {
	sum := 0
	for _, row := range arr {
		for _, val := range row {
			if val > 0 {
				sum += 1
			}
		}
	}

	return sum
}

func PartA() {

	game := ReadFile()
	g_row, g_col := FindGuardPos(game)
	g_orientation := Up

	visited := make([][]int, len(game))
	for i := range visited {
		visited[i] = make([]int, len(game[0]))
	}

	for {
		visited[g_row][g_col] += 1
		g_row, g_col, g_orientation = MakeNextMove(game, g_row, g_col, g_orientation)

		if g_row == -1 && g_col == -1 {
			break
		}

	}

	num_distinct_pos := CountDistinctEntries(visited)
	fmt.Printf("Number of distinct position: %d\n", num_distinct_pos)
}

func PrintGame(game Game) {
	for _, row := range game {
		for _, val := range row {
			fmt.Print(string(val))
		}
		fmt.Println("")
	}
}

func Contains(slice []Orientation, element Orientation) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func TryObstaclePosition(game Game, g_row int, g_col int, g_orientation Orientation, o_row int, o_col int) bool {
	new_game := make(Game, len(game))
	for i := range new_game {
		new_game[i] = make([]Field, len(game[0]))
		copy(new_game[i], game[i])
	}

	visited := map[Coord][]Orientation{}

	new_game[o_row][o_col] = Wall
	for {

		curr_coord := Coord{g_row, g_col}
		val, ok := visited[curr_coord]
		if ok {
			// Cycle detected
			if Contains(val, g_orientation) {
				return true
			}
			visited[curr_coord] = append(val, g_orientation)
		} else {
			visited[curr_coord] = []Orientation{g_orientation}
		}

		g_row, g_col, g_orientation = MakeNextMove(new_game, g_row, g_col, g_orientation)

		if g_row == -1 && g_col == -1 {
			break
		}
	}

	return false
}

func PartB() {

	game := ReadFile()
	g_row, g_col := FindGuardPos(game)
	g_orientation := Up

	barriers := make([][]int, len(game))
	for i := range barriers {
		barriers[i] = make([]int, len(game[0]))
	}

	for o_row := range len(game[0]) {
		for o_col := range len(game) {
			if game[o_row][o_col] == Guard {
				continue
			}
			if TryObstaclePosition(game, g_row, g_col, g_orientation, o_row, o_col) {
				barriers[o_row][o_col] += 1
			}
		}
	}

	num_distinct_pos := CountDistinctEntries(barriers)
	fmt.Printf("Number of distinct position: %d\n", num_distinct_pos)
}

func main() {
	PartA()
	PartB()
}
