package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func isValidPageUpdate(counts map[int]mapset.Set[int], line []string, until int) (bool, int) {

	page_numbers := []int{}
	for _, page := range line {
		page_num, _ := strconv.Atoi(page)
		page_numbers = append(page_numbers, page_num)
	}
	subsequent_pages := mapset.NewSet(page_numbers...)
	for i, page := range page_numbers {
		subsequent_pages.Remove(page)
		page_order, ok := counts[page]

		// Has no successor and is not last element
		if !ok && i != len(page_numbers)-1 {
			return false, -1
		}
		if ok && !subsequent_pages.IsSubset(page_order) {
			return false, -1
		}

		if i == until {
			return true, -1
		}
	}
	middle_elem := page_numbers[(len(page_numbers)-1)/2]
	return true, middle_elem
}

func ReadPageUpdates() (*bufio.Scanner, *os.File, map[int]mapset.Set[int]) {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	counts := map[int]mapset.Set[int]{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		fields := strings.Split(line, "|")
		before, _ := strconv.Atoi(fields[0])
		after, _ := strconv.Atoi(fields[1])
		val, ok := counts[before]
		if ok {
			val.Add(after)
		} else {
			counts[before] = mapset.NewSet[int]()
			counts[before].Add(after)
		}

	}

	return scanner, file, counts
}

func PartA(scanner *bufio.Scanner, file *os.File, counts map[int]mapset.Set[int]) {
	defer file.Close()

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		line_split := strings.Split(line, ",")
		is_valid, middle_elem := isValidPageUpdate(counts, line_split, -1)
		if is_valid {
			sum += middle_elem
		}
	}

	fmt.Printf("Sum is %d\n", sum)
}

func correctPageOrdering(counts map[int]mapset.Set[int], line_split []string) int {

	for i := range line_split {
		for j := i; j < len(line_split); j++ {
			new_line_split := make([]string, len(line_split))
			copy(new_line_split, line_split)

			tmp := new_line_split[i]
			new_line_split[i] = new_line_split[j]
			new_line_split[j] = tmp

			is_valid, _ := isValidPageUpdate(counts, new_line_split, i)
			if is_valid {
				line_split = new_line_split
				break
			}
		}
	}

	is_valid, middle_elem := isValidPageUpdate(counts, line_split, -1)
	if !is_valid {
		log.Fatal("Could not correct page ordering")
	}

	return middle_elem
}

func PartB(scanner *bufio.Scanner, file *os.File, counts map[int]mapset.Set[int]) {
	defer file.Close()

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		line_split := strings.Split(line, ",")
		is_valid, _ := isValidPageUpdate(counts, line_split, -1)
		if is_valid {
			continue
		}

		sum += correctPageOrdering(counts, line_split)

	}

	fmt.Printf("Sum is %d\n", sum)
}

func main() {
	scanner, file, counts := ReadPageUpdates()
	PartA(scanner, file, counts)

	scanner, file, counts = ReadPageUpdates()
	PartB(scanner, file, counts)
}
