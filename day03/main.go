package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readFile reads a file and returns its contents as a string.
func readFile(filename string) string {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// Read the file
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Close the file
	err = file.Close()
	if err != nil {
		panic(err)
	}

	// Return the contents as a string
	return strings.Join(lines, "\n")
}

// findCommonChar takes a string slice and an index and counts how many 1s are at index in each string and counts how many 0s are at index in each string. If there are equal or more 1s return 1, else return 0
func findCommonChar(strings []string, index int) string {
	// Create a slice of ints to hold counts of 1s and 0s
	counts := []int{0, 0}

	// Loop through each string in the slice
	for _, s := range strings {
		// Get the character at the index
		char := string(s[index])

		// If the character at the index in the string is a 1, increment the count of 1s
		if char == "1" {
			counts[0]++
		}

		// If the character at the index in the string is a 0, increment the count of 0s
		if char == "0" {
			counts[1]++
		}
	}

	// If the count of 1s is equal or more than the count of 0s, return "1"
	if counts[0] >= counts[1] {
		return "1"
	}

	// Else return "0"
	return "0"
}

// findScrubberRating takes a string slice and an index and counts how many 1s are at index in each string and counts how many 0s are at index in each string. If there are equal or more 0s return 0, else return 1
func findScrubberRating(strings []string, index int) string {
	// Create a slice of ints to hold counts of 1s and 0s
	counts := []int{0, 0}

	// Loop through each string in the slice
	for _, s := range strings {
		// Get the character at the index
		char := string(s[index])

		// If the character at the index in the string is a 1, increment the count of 1s
		if char == "1" {
			counts[0]++
		}

		// If the character at the index in the string is a 0, increment the count of 0s
		if char == "0" {
			counts[1]++
		}
	}

	// If the count of 1s is equal or more than the count of 0s, return "1"
	if counts[0] >= counts[1] {
		return "0"
	}

	// Else return "0"
	return "1"
}

// findMatchingStrs takes a string, an index, and a string slice and returns items from the string slice that match the character at index in string
func findMatchingStrs(str string, index int, strings []string) []string {
	// Create a slice of strings to hold matching strings
	var matches []string

	// Loop through each string in the slice
	for _, s := range strings {
		// Get the character at the index
		char := string(s[index])

		// If the character at the index in the string matches the character at the index in the string in the slice, append the string to matches
		if char == string(str[index]) {
			matches = append(matches, s)
		}
	}

	// Return the slice of strings
	return matches
}

func main() {
	// Read input.txt
	input := readFile("input.txt")

	// Split input into lines
	lines := strings.Split(input, "\n")

	// Create string named mask
	mask := ""

	// Call findCommonChar for each line in lines with index i and append the result to mask
	for i := 0; i < len(lines[0]); i++ {
		mask += findCommonChar(lines, i)
	}

	// set matches to findMatchingStrs with mask, 0, and lines
	matches := findMatchingStrs(mask, 0, lines)

	// fmt.Print the matches
	for _, match := range matches {
		fmt.Println(match)
	}

	// set i to 0
	i := 0

	// create copy of lines called remaining
	remaining := lines

	// while remaining has more than one element
	// reset mask
	// call findCommonChar with remaining and i and append the result to mask
	// findMatchingStrs with mask, i, and remaining
	// set matches to the result
	// set remaining to matches
	// set i to i + 1
	// fmt.Print remaining
	// Print a divider
	for len(remaining) > 1 {
		mask = ""
		for i := 0; i < len(remaining[0]); i++ {
			mask += findCommonChar(remaining, i)
		}
		matches = findMatchingStrs(mask, i, remaining)
		remaining = matches
		i++
		fmt.Println(remaining)
		fmt.Println("-")
	}

	// fmt.Print the first element of remaining
	fmt.Println(remaining[0])

	// create copy of lines called remainingScrubber
	remainingScrubber := lines

	// set i to 0
	i = 0

	// while remainingScrubber has more than one element
	// reset mask
	// call findScrubberRating with remainingScrubber and i and append the result to mask
	// findMatchingStrs with mask, i, and remainingScrubber
	// set matches to the result
	// set remainingScrubber to matches
	// set i to i + 1
	// fmt.Print remainingScrubber
	// Print a divider
	for len(remainingScrubber) > 1 {
		mask = ""
		for i := 0; i < len(remainingScrubber[0]); i++ {
			mask += findScrubberRating(remainingScrubber, i)
		}
		matches = findMatchingStrs(mask, i, remainingScrubber)
		remainingScrubber = matches
		i++
		fmt.Println(remainingScrubber)
		fmt.Println("-")
	}

	// fmt.Print the first element of remainingScrubber
	fmt.Println(remainingScrubber[0])

	// Convert first element of remaining to binary integer named oxy and ignore err
	oxy, _ := strconv.ParseInt(remaining[0], 2, 64)

	// Convert first element of remainingScrubber to binary integer named scrubber and ignore err
	scrubber, _ := strconv.ParseInt(remainingScrubber[0], 2, 64)

	// fmt.Print the multiplication of oxy and scrubber
	fmt.Println(oxy * scrubber)
}

func oneStar() {
	// Read input.txt
	input := readFile("input.txt")

	// Split input into lines
	lines := strings.Split(input, "\n")

	// Loop over lines and count frequency of "1" per column
	var freq []int
	for i := 0; i < len(lines[0]); i++ {
		freq = append(freq, 0)
	}

	// Count frequency of "1" per column
	for _, line := range lines {
		for i, c := range line {
			if c == '1' {
				freq[i]++
			}
		}
	}

	// Check if freq is half or more than number of lines
	// If so print 1 else print 0 and store output as string
	var output string
	for _, f := range freq {
		if f >= len(lines)/2 {
			output += "1"
		} else {
			output += "0"
		}
	}

	// Convert output to binary integer
	outputInt, err := strconv.ParseInt(output, 2, 64)
	if err != nil {
		panic(err)
	}

	// Count length of first line
	firstLineLength := len(lines[0])

	// Make string of 1s the length of firstLineLength
	ones := strings.Repeat("1", firstLineLength)

	// Convert ones to binary integer
	onesInt, err := strconv.ParseInt(ones, 2, 64)
	if err != nil {
		panic(err)
	}

	// xor onesInt and outputInt and store in a variable
	xor := onesInt ^ outputInt

	// Print xor in binary format
	fmt.Printf("%b\n", xor)

	// Print outputInt multiplied by xor
	fmt.Printf("%d\n", outputInt*xor)
}
