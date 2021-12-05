package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// max returns the max of all ints from varargs
func findMax(integers ...int) int {
	var result int
	for _, v := range integers {
		if v > result {
			result = v
		}
	}
	return result
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var max int

	for _, line := range lines {
		fields := strings.Fields(line)
		a, b := fields[0], fields[len(fields)-1]

		x1, _ := strconv.Atoi(strings.Split(a, ",")[0])
		y1, _ := strconv.Atoi(strings.Split(a, ",")[1])

		x2, _ := strconv.Atoi(strings.Split(b, ",")[0])
		y2, _ := strconv.Atoi(strings.Split(b, ",")[1])

		max = findMax(max, x1, y1, x2, y2)
	}

	grid := [][]int{}
	for i := 0; i <= max; i++ {
		grid = append(grid, make([]int, max+1))
	}

	for _, line := range lines {
		fields := strings.Fields(line)
		a, b := fields[0], fields[len(fields)-1]

		x1, _ := strconv.Atoi(strings.Split(a, ",")[0])
		y1, _ := strconv.Atoi(strings.Split(a, ",")[1])

		x2, _ := strconv.Atoi(strings.Split(b, ",")[0])
		y2, _ := strconv.Atoi(strings.Split(b, ",")[1])

		fmt.Printf("Moving between (%d, %d) -> (%d, %d)\n", x1, y1, x2, y2)
		if x1 == x2 {
			fmt.Println("Moving vertically")
			if y1 < y2 {
				fmt.Println("Moving down")
				for i := y1; i <= y2; i++ {
					grid[i][x1]++
				}
			} else {
				fmt.Println("Moving up")
				for i := y2; i <= y1; i++ {
					grid[i][x1]++
				}
			}
		} else if y1 == y2 {
			fmt.Println("Moving horizontally")
			if x1 < x2 {
				fmt.Println("Moving right")
				for i := x1; i <= x2; i++ {
					grid[y1][i]++
				}
			} else {
				fmt.Println("Moving left")
				for i := x2; i <= x1; i++ {
					grid[y1][i]++
				}
			}
		} else {
			fmt.Println("Moving diagonally")
			var xDir, yDir int

			if x1 <= x2 {
				xDir = 1
			} else {
				xDir = -1
			}

			if y1 <= y2 {
				yDir = 1
			} else {
				yDir = -1
			}

			i, j := x1, y1

			for i != x2 || j != y2 {
				grid[j][i] += 1
				i += xDir
				j += yDir
			}
			grid[j][i] += 1
		}
		fmt.Println("----")
	}

	// Count numbers above 1 in grid
	var count int
	for _, row := range grid {
		for _, v := range row {
			if v > 1 {
				count++
			}
		}
	}

	// pretty print grid
	for i := 0; i <= max; i++ {
		for j := 0; j <= max; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}

	fmt.Println(count)
}
