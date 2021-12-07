package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// parseNumbers parses a string of numbers separated by commas into a slice of ints
func parseNumbers(line string) []int {
	// split line by comma
	var numbers = strings.Split(line, ",")
	// create a new slice of ints
	var intNumbers []int
	// loop over all numbers
	for _, number := range numbers {
		// convert number to int
		var intNumber, _ = strconv.Atoi(number)
		// append int to intNumbers
		intNumbers = append(intNumbers, intNumber)
	}
	// return intNumbers
	return intNumbers
}

// sort sorts a slice of ints
func sort(numbers []int) {
	// create a new slice of ints
	var sortedNumbers = make([]int, len(numbers))
	// copy numbers to sortedNumbers
	copy(sortedNumbers, numbers)
	// sort numbers
	for i := len(sortedNumbers) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if sortedNumbers[j] > sortedNumbers[j+1] {
				sortedNumbers[j], sortedNumbers[j+1] = sortedNumbers[j+1], sortedNumbers[j]
			}
		}
	}
	// copy sortedNumbers to numbers
	copy(numbers, sortedNumbers)
}

// abs returns the absolute value of an int
func abs(number int) int {
	// if number is negative
	if number < 0 {
		// return number multiplied by -1
		return number * -1
	}
	// return number
	return number
}

// SumDisplacementOfNumbers takes an int slice and an int and returns the sum of the displacement of each number compared to the int
func SumDisplacementOfNumbers(numbers []int, intNumber int) int {
	// create a new slice of ints
	var sortedNumbers = make([]int, len(numbers))
	// copy numbers to sortedNumbers
	copy(sortedNumbers, numbers)
	// sort numbers
	sort(sortedNumbers)

	// create a new int
	var sum int
	// loop over all numbers
	for _, number := range numbers {
		// add number to sum
		sum += abs(number - intNumber)
	}
	// return sum
	return sum
}

func main() {
	// open input.txt and ignore error
	file, _ := os.Open("input.txt")
	// close file
	defer file.Close()

	// create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// store all lines
	var lines []string
	// loop over all lines
	for scanner.Scan() {
		// store line in lines
		lines = append(lines, scanner.Text())
	}

	// split first line by comma and parse to int slice
	var firstLine = lines[0]
	var firstLineNumbers = parseNumbers(firstLine)

	// range over firstLineNumbers and find the lowest sum of displacement
	var lowestSumOfDisplacement int

	// loop over all numbers
	for _, number := range firstLineNumbers {
		// find sum of displacement of firstLineNumbers and number
		var sumOfDisplacement = SumDisplacementOfNumbers(firstLineNumbers, number)
		// if sumOfDisplacement is lower than lowestSumOfDisplacement
		if sumOfDisplacement < lowestSumOfDisplacement || lowestSumOfDisplacement == 0 {
			// set lowestSumOfDisplacement to sumOfDisplacement
			lowestSumOfDisplacement = sumOfDisplacement
		}
	}

	// print lowestSumOfDisplacement
	println(lowestSumOfDisplacement)
}
