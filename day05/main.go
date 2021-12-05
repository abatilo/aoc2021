package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Vector2D represents a 2D vector
type Vector2D struct {
	X int
	Y int
}

// Vector2DPairs represents two Vector2D objects
type Vector2DPairs struct {
	A Vector2D
	B Vector2D
}

// VectorPath takes a VectorPairs and returns a slice of Vector2D objects for each point bewteen them and prints the pairs
func VectorPath(pairs Vector2DPairs) []Vector2D {
	// Create a slice of Vector2D objects
	path := make([]Vector2D, 0)

	// Add the first pair
	path = append(path, pairs.A)

	// Add the rest of the pairs
	for x := pairs.A.X; x != pairs.B.X; {
		if x < pairs.B.X {
			x++
		} else {
			x--
		}
		path = append(path, Vector2D{x, pairs.A.Y})
	}
	for y := pairs.A.Y; y != pairs.B.Y; {
		if y < pairs.B.Y {
			y++
		} else {
			y--
		}
		path = append(path, Vector2D{pairs.B.X, y})
	}

	return path
}

// VectorFrequencies takes a slice of Vector2D objects and returns a map of Vector2D objects and their frequency
func VectorFrequencies(path []Vector2D) map[Vector2D]int {
	frequencies := make(map[Vector2D]int)
	for _, vector := range path {
		frequencies[vector]++
	}
	return frequencies
}

// openFile opens a file and returns a file pointer and an error
func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// parseInputLine takes a string of format "X1,Y1 -> X2,Y2" and returns a Vector2DPairs and converts each coordinate
// to an int
func parseInputLine(line string) Vector2DPairs {
	// Split line into two strings
	splitLine := strings.Split(line, " -> ")

	// Split first string into two coordinates
	coordinates1 := strings.Split(splitLine[0], ",")

	// Split second string into two coordinates
	coordinates2 := strings.Split(splitLine[1], ",")

	// Convert coordinates to ints
	x1, _ := strconv.Atoi(coordinates1[0])
	y1, _ := strconv.Atoi(coordinates1[1])
	x2, _ := strconv.Atoi(coordinates2[0])
	y2, _ := strconv.Atoi(coordinates2[1])

	// Create Vector2DPairs
	pairs := Vector2DPairs{Vector2D{x1, y1}, Vector2D{x2, y2}}

	return pairs
}

// println is a convenience function for printing to stdout
func println(a ...interface{}) (int, error) {
	return fmt.Println(a...)
}

// PrettyPrintGrid takes a slice of Vector2D objects and prints them in a 9x9 grid
func PrettyPrintGrid(path []Vector2D) {
	// Find the min and max X and Y values
	var minX, maxX, minY, maxY int
	for _, vector := range path {
		if vector.X < minX {
			minX = vector.X
		}
		if vector.X > maxX {
			maxX = vector.X
		}
		if vector.Y < minY {
			minY = vector.Y
		}
		if vector.Y > maxY {
			maxY = vector.Y
		}
	}

	// Print the grid
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			var found bool
			for _, vector := range path {
				if vector.X == x && vector.Y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// PrettyPrintFrequencies takes a map of Vector2D objects and their frequencies and prints the frequency in a grid
func PrettyPrintFrequencies(frequencies map[Vector2D]int) {
	// Find the min and max X and Y values
	var minX, maxX, minY, maxY int
	for vector, _ := range frequencies {
		if vector.X < minX {
			minX = vector.X
		}
		if vector.X > maxX {
			maxX = vector.X
		}
		if vector.Y < minY {
			minY = vector.Y
		}
		if vector.Y > maxY {
			maxY = vector.Y
		}
	}

	// Print the grid
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			var found bool
			for vector, _ := range frequencies {
				if vector.X == x && vector.Y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print(frequencies[Vector2D{x, y}])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	// Open file input.txt and ignore errors
	file, _ := openFile("input.txt")
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)

	// Create a slice of Vector2DPairs by reading each line
	var pairs []Vector2DPairs
	for scanner.Scan() {
		pairs = append(pairs, parseInputLine(scanner.Text()))
	}

	// Keep pairs that are either horizontal or vertical
	var horizontalPairs, verticalPairs []Vector2DPairs
	for _, pair := range pairs {
		if pair.A.X == pair.B.X {
			verticalPairs = append(verticalPairs, pair)
		} else if pair.A.Y == pair.B.Y {
			horizontalPairs = append(horizontalPairs, pair)
		}
	}

	// combine horizontalpairs and verticalparis into a single slice of Vector2DPairs
	var combinedPairs []Vector2DPairs
	for _, pair := range horizontalPairs {
		combinedPairs = append(combinedPairs, pair)
	}
	for _, pair := range verticalPairs {
		combinedPairs = append(combinedPairs, pair)
	}

	// print the combined pairs
	println("Combined pairs:")
	for _, pair := range combinedPairs {
		println(pair.A, " -> ", pair.B)
	}

	// print vector path of first horizontal pair
	println("First horizontal pair:")
	println(VectorPath(combinedPairs[0]))

	// Create a slice of Vector2D objects by combining all the horizontal and vertical paths
	var path []Vector2D
	for _, pair := range horizontalPairs {
		path = append(path, VectorPath(pair)...)
	}
	for _, pair := range verticalPairs {
		path = append(path, VectorPath(pair)...)
	}

	// print all paths
	println("All paths:")
	for _, vector := range path {
		println(vector)
	}

	// print line divider
	println()

	// calculate vector frequencies
	frequencies := VectorFrequencies(path)

	// count the number of times 2 or more vectors appear
	var count int
	for _, frequency := range frequencies {
		if frequency >= 2 {
			count++
		}
	}

	// print the number of times 2 or more vectors appear
	println("Number of times 2 or more vectors appear:", count)
}
