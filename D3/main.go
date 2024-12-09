package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func SolvePartA(content []byte) {
	r, err := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	if err != nil {
		log.Fatal(err)
	}

	matches := r.FindAllStringSubmatch(string(content), -1)

	res := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		res += a * b
	}

	fmt.Printf("Result: %d\n", res)
}

func SolvePartB(content []byte) {
	r, err := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)|don't\(\)|do\(\)`)
	if err != nil {
		log.Fatal(err)
	}

	matches := r.FindAllStringSubmatch(string(content), -1)

	res := 0
	do_count := true
	for _, match := range matches {
		if do_count && strings.HasPrefix(match[0], "mul") {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			res += a * b
		} else if match[0] == "don't()" {
			do_count = false
		} else if match[0] == "do()" {
			do_count = true
		}
	}

	fmt.Printf("Result: %d\n", res)
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	SolvePartA(content)
	SolvePartB(content)
}
