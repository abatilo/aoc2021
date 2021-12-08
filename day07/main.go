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

// Sum returns the sum of a slice of ints
func Sum(numbers []int) int {
	// create a new int
	var sum int
	// loop over all numbers
	for _, number := range numbers {
		// add number to sum
		sum += number
	}
	// return sum
	return sum
}

// sumOfRange returns the sum of the range of numbers between start and end exclusive
func sumOfRange(start, end int) int {
	// create a new slice of ints
	var sumOfRange []int
	// loop over all numbers between start and end
	for i := start; i <= end; i++ {
		// append i to sumOfRange
		sumOfRange = append(sumOfRange, i)
	}
	// return sum of sumOfRange
	return Sum(sumOfRange)
}

// sumDisplacement takes int slice numbers and int target
// if number is less than target, add sumOfRange(1, target-number) to sum
// else add sumOfRange(1, number-target) to sum
func SumDisplacement(numbers []int, target int) int {
	// create a new int
	var sum int
	// loop over all numbers
	for _, number := range numbers {
		// if number is less than target
		if number < target {
			// add sumOfRange(1, target-number) to sum
			sum += sumOfRange(1, target-number)
		} else {
			// add sumOfRange(1, number-target) to sum
			sum += sumOfRange(1, number-target)
		}
	}
	// return sum
	return sum
}

// sort sorts a slice of ints in ascending order
func sort(numbers []int) {
	// create a new int
	var min int
	// loop over all numbers
	for i := 0; i < len(numbers); i++ {
		// set min to numbers[i]
		min = numbers[i]
		// loop over all numbers
		for j := i + 1; j < len(numbers); j++ {
			// if numbers[j] < min
			if numbers[j] < min {
				// set min to numbers[j]
				min = numbers[j]
				// swap numbers[i] and numbers[j]
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
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

	// sorted first line numbers
	var sortedFirstLineNumbers = append([]int{}, firstLineNumbers...)
	// sort first line numbers
	sort(sortedFirstLineNumbers)

	// max of first line numbers
	var maxOfFirstLineNumbers = sortedFirstLineNumbers[len(sortedFirstLineNumbers)-1]

	// lowestSumOfDisplacement starts at 0
	var lowestSumOfDisplacement int

	// loop over 0 to max of first line numbers
	// if number == 5
	//	print number
	//  print sumDisplacement
	// 	print divider
	// if sumDisplacement < lowestSumOfDisplacement o lowestSumOfDisplacement == 0
	//	lowestSumOfDisplacement = sumDisplacement

	for i := 0; i <= maxOfFirstLineNumbers; i++ {
		// if i == 5
		//	print i
		//  print sumDisplacement
		// 	print divider
		// if sumDisplacement < lowestSumOfDisplacement o lowestSumOfDisplacement == 0
		//	lowestSumOfDisplacement = sumDisplacement
		if i == 5 {
			println(i)
			println(SumDisplacement(firstLineNumbers, i))
			println("---")
		}
		if SumDisplacement(firstLineNumbers, i) < lowestSumOfDisplacement || lowestSumOfDisplacement == 0 {
			lowestSumOfDisplacement = SumDisplacement(firstLineNumbers, i)
		}
	}

	// print lowestSumOfDisplacement
	println(lowestSumOfDisplacement)
}
