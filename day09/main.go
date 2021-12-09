package main

import (
	"bufio"
	"fmt"
	"os"
)

// isLocalMinimum takes a grid, x, y, and value and returns true if the value is less than its peers. If column is 0 don't check to the left. If column is len(grid[0]) dont check to the right. if row is 0 don't check above and if row is len(grid)-1 don't check below
func isLocalMinimum(grid [][]int, x, y, val int) bool {
	if x == 0 {
		if y == 0 {
			return val < grid[y+1][x] && val < grid[y][x+1]
		} else if y == len(grid)-1 {
			return val < grid[y-1][x] && val < grid[y][x+1]
		} else {
			return val < grid[y-1][x] && val < grid[y][x+1] && val < grid[y+1][x]
		}
	} else if x == len(grid[0])-1 {
		if y == 0 {
			return val < grid[y+1][x] && val < grid[y][x-1]
		} else if y == len(grid)-1 {
			return val < grid[y-1][x] && val < grid[y][x-1]
		} else {
			return val < grid[y-1][x] && val < grid[y][x-1] && val < grid[y+1][x]
		}
	} else {
		if y == 0 {
			return val < grid[y+1][x] && val < grid[y][x-1] && val < grid[y][x+1]
		} else if y == len(grid)-1 {
			return val < grid[y-1][x] && val < grid[y][x-1] && val < grid[y][x+1]
		} else {
			return val < grid[y-1][x] && val < grid[y][x-1] && val < grid[y][x+1] && val < grid[y+1][x]
		}
	}
}

// localMinimums takes a grid and returns a slice of all numbers that are lower than all of its peers
func localMinimums(grid [][]int) []int {
	var lowestPoints []int
	for y, row := range grid {
		for x, val := range row {
			if isLocalMinimum(grid, x, y, val) {
				lowestPoints = append(lowestPoints, val)
			}
		}
	}
	return lowestPoints
}

func main() {
	// Open input.txt for reading
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the file
	scanner := bufio.NewScanner(file)

	// Read all lines into slice
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// convert each line into slice of chars
	var grid [][]int
	for _, line := range lines {
		var row []int
		for _, char := range line {
			row = append(row, int(char)-48)
		}
		grid = append(grid, row)
	}

	// print grid with spaces
	for _, row := range grid {
		for _, char := range row {
			fmt.Printf("%d ", char)
		}
		fmt.Println()
	}

	// Find all lowest points
	lowestPoints := localMinimums(grid)

	// print lowest points
	fmt.Println("Lowest points:", lowestPoints)

	// Sum all lowest points
	sum := 0
	for _, point := range lowestPoints {
		sum += point
	}

	// Add length of lowest points to sum
	sum += len(lowestPoints)

	fmt.Println("Sum of all points:", sum)
}
