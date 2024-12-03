package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := [][]int{}
	for scanner.Scan() {
		report := []int{}

		line := scanner.Text()
		line_split := strings.Split(line, " ")

		for _, value := range line_split {
			conv_int, _ := strconv.Atoi(value)
			report = append(report, conv_int)
		}

		data = append(data, report)
	}

	return data
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkIfReportIsSafe(report []int) bool {

	is_decreasing := true
	if report[0] < report[1] {
		is_decreasing = false
	}
	for idx := 1; idx < len(report); idx++ {

		diff := Abs(report[idx-1] - report[idx])
		if report[idx-1] == report[idx] {
			return false
		} else if !is_decreasing && report[idx-1] > report[idx] {
			return false
		} else if is_decreasing && report[idx-1] < report[idx] {
			return false
		} else if !(1 <= diff && diff <= 3) {
			return false
		}
	}

	return true
}

func CountSafeReports(data [][]int, tolerate_one bool) {

	num_safe_reports := 0

	for _, report := range data {

		count_report := checkIfReportIsSafe(report)
		if count_report {
			num_safe_reports += 1
			continue
		}

		if tolerate_one {
			for idx := range report {

				new_report := append([]int(nil), report[:idx]...)
				new_report = append(new_report, report[idx+1:]...)
				count_report := checkIfReportIsSafe(new_report)
				if count_report {
					num_safe_reports += 1
					break
				}
			}
		}
	}

	fmt.Printf("Part A: Num Safe reports %d\n", num_safe_reports)
}

func main() {
	data := ReadFile()
	CountSafeReports(data, false)
	CountSafeReports(data, true)
}
