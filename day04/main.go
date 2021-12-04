package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parseStringToIntSlice parses a string with commas into an int slice
func parseStringToIntSlice(input string) []int {
	// Split input string by comma
	inputSplit := strings.Split(input, ",")

	// Create int slice
	inputIntSlice := make([]int, 0)

	// Parse string slice to int slice
	for _, input := range inputSplit {
		inputInt, _ := strconv.Atoi(input)
		inputIntSlice = append(inputIntSlice, inputInt)
	}

	return inputIntSlice
}

// Remove duplicate whitespace and split string by whitespace
func parseToIntSliceWithSpaces(input string) []int {
	// Remove duplicate whitespace
	input = strings.Join(strings.Fields(input), " ")

	// Split string by whitespace
	inputSplit := strings.Split(input, " ")

	// Create int slice
	inputIntSlice := make([]int, 0)

	// Parse string slice to int slice
	for _, input := range inputSplit {
		inputInt, _ := strconv.Atoi(input)
		inputIntSlice = append(inputIntSlice, inputInt)
	}

	return inputIntSlice
}

// parseBingoGrid takes in a string slice if the line is empty then continue else parse into space separated 2d int slice
func parseBingoGrid(lines []string) [][]int {
	// Create 2d int slice
	bingoGrid := make([][]int, 0)

	// Loop through lines
	for _, line := range lines {
		// If line is empty then continue
		if line == "" {
			continue
		}

		// Parse line to int slice
		lineIntSlice := parseToIntSliceWithSpaces(line)

		// Append line int slice to bingo grid
		bingoGrid = append(bingoGrid, lineIntSlice)
	}

	return bingoGrid
}

// contains returns true if numbers contains number
func contains(numbers []int, number int) bool {
	for _, num := range numbers {
		if num == number {
			return true
		}
	}

	return false
}

// hasAllRowNumbers returns true if all numbers in row are in numbers
func hasAllRowNumbers(row []int, numbers []int) bool {
	for _, number := range row {
		if !contains(numbers, number) {
			return false
		}
	}

	return true
}

// hasAllColumnNumbers returns true if all numbers in column are in numbers
func hasAllColumnNumbers(column []int, numbers []int) bool {
	for _, number := range column {
		if !contains(numbers, number) {
			return false
		}
	}

	return true
}

// gridHasBingo returns true if bingo grid has bingo for rows and columns
func gridHasBingo(bingoGrid [][]int, drawnNumbers []int) bool {
	// Check rows
	for _, row := range bingoGrid {
		if hasAllRowNumbers(row, drawnNumbers) {
			return true
		}
	}

	// Check columns
	for i := 0; i < len(bingoGrid[0]); i++ {
		column := make([]int, 0)
		for _, row := range bingoGrid {
			column = append(column, row[i])
		}

		if hasAllColumnNumbers(column, drawnNumbers) {
			return true
		}
	}

	return false
}

// sumUnfoundNumbers returns the sum of all numbers in a bingo grid that are not found
func sumUnfoundNumbers(bingoGrid [][]int, drawnNumbers []int) int {
	// Create int slice
	unfoundNumbers := make([]int, 0)

	// Loop through bingo grid
	for _, row := range bingoGrid {
		for _, number := range row {
			if !contains(drawnNumbers, number) {
				unfoundNumbers = append(unfoundNumbers, number)
			}
		}
	}

	// Sum unfound numbers
	sum := 0
	for _, number := range unfoundNumbers {
		sum += number
	}

	return sum
}

func main() {
	// Open input.txt file for reading and ignore error
	file, _ := os.Open("input.txt")

	// Read file line by line into string slice lines
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// drawnNumbers is parseIntToIntSlice(lines[0])
	drawnNumbers := parseStringToIntSlice(lines[0])

	// Remove first line from lines
	lines = lines[1:]

	// Chunk lines into slices of 6
	chunks := make([][]string, 0)
	for i := 0; i < len(lines); i += 6 {
		chunks = append(chunks, lines[i:i+6])
	}

	// Convert chunks to slice of bingo grids
	bingoGrids := make([][][]int, 0)
	for _, chunk := range chunks {
		bingoGrids = append(bingoGrids, parseBingoGrid(chunk))
	}

	// set drawnCount to 5
	drawnCount := 5

	// store winningGrid
	winningGrid := make([][]int, 0)

	// winningDrawCount
	winningDrawCount := 0

	// print divider
	fmt.Println("------------------------------------------------------")

	// currentDrawnNumbers is drawnNumbers slice from 0 until drawnCount
	currentDrawnNumbers := drawnNumbers[:drawnCount]

	// track which grids have bingo
	bingoGridsFound := make([]bool, 0)

	// instantiate bingoGridsFound with number of bingo grids
	for i := 0; i < len(bingoGrids); i++ {
		bingoGridsFound = append(bingoGridsFound, false)
	}

	// while drawnCount is less than len(drawnNumbers)
	// check each bingo grid for bingo using currentDrawnNumbers
	// skip bingo grids that have already been found
	// if bingo is found then set winningGrid to bingo grid and set winningDrawCount and keep going
	// else increment drawnCount and set currentDrawnNumbers to drawnNumbers slice from 0 until drawnCount
	for drawnCount < len(drawnNumbers) {
		for i, bingoGrid := range bingoGrids {
			if bingoGridsFound[i] {
				continue
			}

			if gridHasBingo(bingoGrid, currentDrawnNumbers) {
				winningGrid = bingoGrid
				winningDrawCount = drawnCount
				bingoGridsFound[i] = true
			}
		}

		drawnCount++
		currentDrawnNumbers = drawnNumbers[:drawnCount]
	}

	// currentDrawnNumbers to drawnNumbers slice from 0 until winningDrawCount
	currentDrawnNumbers = drawnNumbers[:winningDrawCount]

	// Print currentDrawnNumbers
	fmt.Println("Current Drawn Numbers:", currentDrawnNumbers)

	// Print winningGrid
	fmt.Println("Winning Grid:")
	for _, row := range winningGrid {
		fmt.Println(row)
	}

	// store sum of unfound numbers of winning grid
	sum := sumUnfoundNumbers(winningGrid, currentDrawnNumbers)

	// Get last drawn number from currentDrawnNumbers
	lastDrawnNumber := currentDrawnNumbers[len(currentDrawnNumbers)-1]

	// Print sum of unfound numbers of winning grid
	fmt.Println("Sum of Unfound Numbers:", sum)

	// Print last drawn number
	fmt.Println("Last Drawn Number:", lastDrawnNumber)

	// Multiply sum of unfound numbers by last drawn number
	product := sum * lastDrawnNumber

	// Print product
	fmt.Println("Product:", product)
}
