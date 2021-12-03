package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// readFile reads a file and returns a string slice of its contents
func readFile(filename string) []string {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// Close file on exit
	defer file.Close()

	// Read file into a string slice
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	// Read file contents to string slice
	lines := readFile("input.txt")

	// Convert lines to int slice
	var nums []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	// Count number of times nums increases
	var count int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			count++
		}
	}

	// Print result
	fmt.Println(count)
}
