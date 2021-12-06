package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// openFile opens a file and returns a slice of strings
func openFile(filename string) ([]string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// Create a new scanner
	scanner := bufio.NewScanner(file)
	// Create a slice to hold the lines
	lines := make([]string, 0)
	// Loop through the lines
	for scanner.Scan() {
		// Append the line to the slice
		lines = append(lines, scanner.Text())
	}
	// Close the file
	file.Close()
	// Return the slice of strings
	return lines, nil
}

// SplitCommasToIntSlice takes a string and splits it on commas and returns a slice of ints
func SplitCommasToIntSlice(input string) []int {
	// Create a slice to hold the ints
	slice := make([]int, 0)
	// Split the string on commas
	for _, s := range strings.Split(input, ",") {
		// Convert the string to an int
		i, _ := strconv.Atoi(s)
		// Append the int to the slice
		slice = append(slice, i)
	}
	// Return the slice of ints
	return slice
}

func main() {
	// Open input.txt and ignore the error
	input, _ := openFile("input.txt")

	// print the first line of the input
	println(input[0])

	// days is the first line of input as a slice of ints
	days := SplitCommasToIntSlice(input[0])

	// Create int slice that can hold 9 elements
	// This is the size of the memory block
	memory := make([]int, 9)

	// Count frequency of each value in days and store in memory
	for _, d := range days {
		memory[d]++
	}

	// print each value in memory on the same line
	for _, m := range memory {
		print(m)
		print(",")
	}

	// print new line
	println()
	// print divider
	println("----------------------------------------")

	// loop 256 times
	// loop through values in memory
	// store memory[0] in temp
	// if j >= 1 set memory[j - 1] to memory[j]
	// set memory[8] to 0
	// if temp is greater than 0, add it to memory[6] and set memory[8] to temp
	// print each value in memory on the same line
	// print divider

	for i := 0; i < 256; i++ {
		temp := memory[0]
		for j := 1; j < 9; j++ {
			memory[j-1] = memory[j]
		}
		memory[8] = 0
		if temp > 0 {
			memory[6] += temp
			memory[8] = temp
		}
		for _, m := range memory {
			print(m)
			print(",")
		}
		println()
		println("----------------------------------------")
	}

	// sum is an unsigned int 64
	sum := uint64(0)
	// loop through values in memory
	// if value is greater than 0, add it to sum
	for _, m := range memory {
		if m > 0 {
			sum += uint64(m)
		}
	}

	// print sum
	println(sum)
}
