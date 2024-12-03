package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ReadFile() ([]int, []int) {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close() // Ensure the file is closed when done

	scanner := bufio.NewScanner(file)

	left := []int{}
	right := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		line_split := strings.Split(line, "   ")

		left_int, _ := strconv.Atoi(line_split[0])
		right_int, _ := strconv.Atoi(line_split[1])
		left = append(left, left_int)
		right = append(right, right_int)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return left, right
}

func SolvePartA(left []int, right []int) {
	sort.Ints(left)
	sort.Ints(right)

	res := 0
	for i := range left {
		res += Abs(left[i] - right[i])
	}

	fmt.Printf("Part B result is %d\n", res)
}

func SolvePartB(left []int, right []int) {
	hashMap := make(map[int]int)

	for _, value := range right {
		_, exists := hashMap[value]

		if exists {
			hashMap[value] += 1
		} else {
			hashMap[value] = 1
		}
	}

	res := 0
	for _, value := range left {
		count, exists := hashMap[value]

		if exists {
			res += value * count
		}
	}

	fmt.Printf("Part B result is %d\n", res)
}

func main() {

	left, right := ReadFile()
	SolvePartA(left, right)
	SolvePartB(left, right)
}
