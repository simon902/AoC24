package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile() [][]rune {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	field := [][]rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_arr := []rune(line)
		field = append(field, line_arr)
	}

	return field
}

func IsValid(field [][]rune, x int, y int) bool {
	y_max := len(field)
	x_max := len(field[0])

	return 0 <= x && x < x_max && 0 <= y && y < y_max
}

func CountXMAS(field [][]rune, x int, y int) int {
	indices := [][]struct{ X, Y int }{
		{{X: x - 1, Y: y - 1}, {X: x - 2, Y: y - 2}, {X: x - 3, Y: y - 3}},
		{{X: x + 1, Y: y - 1}, {X: x + 2, Y: y - 2}, {X: x + 3, Y: y - 3}},
		{{X: x + 1, Y: y + 1}, {X: x + 2, Y: y + 2}, {X: x + 3, Y: y + 3}},
		{{X: x - 1, Y: y + 1}, {X: x - 2, Y: y + 2}, {X: x - 3, Y: y + 3}},

		{{X: x, Y: y - 1}, {X: x, Y: y - 2}, {X: x, Y: y - 3}},
		{{X: x + 1, Y: y}, {X: x + 2, Y: y}, {X: x + 3, Y: y}},
		{{X: x, Y: y + 1}, {X: x, Y: y + 2}, {X: x, Y: y + 3}},
		{{X: x - 1, Y: y}, {X: x - 2, Y: y}, {X: x - 3, Y: y}},
	}

	count := 0

	for _, direction := range indices {
		valid := true
		for i, coord := range direction {
			if !IsValid(field, coord.X, coord.Y) {
				valid = false
				break
			}

			if i == 0 && field[coord.Y][coord.X] != 'M' {
				valid = false
				break
			} else if i == 1 && field[coord.Y][coord.X] != 'A' {
				valid = false
				break
			} else if i == 2 && field[coord.Y][coord.X] != 'S' {
				valid = false
				break
			}
		}
		if valid {
			count += 1
		}
	}
	return count
}

func isX_MAS(field [][]rune, x int, y int) bool {
	indices := [][]struct{ X, Y int }{
		{{X: x - 1, Y: y - 1}, {X: x + 1, Y: y + 1}},
		{{X: x + 1, Y: y - 1}, {X: x - 1, Y: y + 1}},
	}

	for _, direction := range indices {
		coord1 := direction[0]
		coord2 := direction[1]
		if !IsValid(field, coord1.X, coord1.Y) || !IsValid(field, coord2.X, coord2.Y) {
			return false
		}

		dir_string := string([]rune{field[coord1.Y][coord1.X], field[coord2.Y][coord2.X]})

		if !strings.Contains(dir_string, "M") || !strings.Contains(dir_string, "S") {
			return false
		}

	}
	return true
}

func main() {
	field := ReadFile()

	fmt.Println(field)

	sum := 0
	for y, row := range field {
		for x, val := range row {
			if val == 'X' {
				sum += CountXMAS(field, x, y)
			}
		}
	}

	fmt.Printf("XMAS count is %d\n", sum)

	sum = 0
	for y, row := range field {
		for x, val := range row {
			if val == 'A' {
				if isX_MAS(field, x, y) {
					sum += 1
				}
			}
		}
	}

	fmt.Printf("X-MAS count is %d\n", sum)

}
