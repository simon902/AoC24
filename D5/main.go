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

func isValidPageUpdate(counts map[int]mapset.Set[int], line string) (bool, int) {
	// page_numbers := strings.Split(line, ",")

	page_numbers := []int{}
	for _, page := range strings.Split(line, ",") {
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
	}
	middle_elem := page_numbers[(len(page_numbers)-1)/2]
	return true, middle_elem
}

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	counts := map[int]mapset.Set[int]{}
	sum := 0
	process_page_update := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			process_page_update = true
			continue
		}
		if !process_page_update {
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
		} else {
			is_valid, middle_elem := isValidPageUpdate(counts, line)
			if is_valid {
				sum += middle_elem
			}
		}
	}
	// fmt.Println(counts)
	fmt.Printf("Sum is %d\n", sum)

}
