package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read every line from a file and convert to slice of integers
	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read every line from a file and convert to slice of integers
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// Convert lines to integers
	ints := make([]int, 0)
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}

	// Chunk ints by 3
	chunks := make([][]int, 0)
	for i := 0; i < len(ints); i += 3 {
		chunks = append(chunks, ints[i:i+3])
	}

	// Add up each chunk

	// Count how many times sum increases
	count := 0
	for i := 0; i < len(sums)-1; i++ {
		if sums[i] < sums[i+1] {
			count++
		}
	}

	fmt.Println(count)
}
