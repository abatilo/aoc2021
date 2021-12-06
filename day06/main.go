package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// WrapAroundDecrementSlice takes an int slice and decrements each element by one. If an element becomes negative, set its value to 6 and add an 8 to the slice and return the new slice
func WrapAroundDecrementSlice(slice []int) []int {
	for i := range slice {
		slice[i]--
		if slice[i] < 0 {
			slice[i] = 6
			slice = append(slice, 8)
		}
	}
	return slice
}

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

	// loop 80 times
	// call WrapAroundDecrementSlice on days and re-assign days
	// print every element of days on the same line
	// print a divider
	for i := 0; i < 80; i++ {
		days = WrapAroundDecrementSlice(days)
		for _, d := range days {
			print(d)
		}
		println()
	}

	// print length of days
	println(len(days))
}
