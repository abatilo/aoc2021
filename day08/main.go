package main

import (
	"bufio"
	"fmt"
	"os"
)

// parseLine takes a string and fmt.Scanf("%s %s %s %s %s %s %s %s %s %s | %s %s %s %s") and returns two slices. []string{a, b, c, d, e, f, g, h, i, j} and []string{k, l, m, n}
func parseLine(line string) ([]string, []string) {
	var a, b, c, d, e, f, g, h, i, j, k, l, m, n string
	_, err := fmt.Sscanf(line, "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s", &a, &b, &c, &d, &e, &f, &g, &h, &i, &j, &k, &l, &m, &n)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return []string{a, b, c, d, e, f, g, h, i, j}, []string{k, l, m, n}
}

// stringToSlice takes a string and returns a slice of strings for each character
func stringToSlice(s string) []string {
	var slice []string
	for _, c := range s {
		slice = append(slice, string(c))
	}
	return slice
}

// destructureSlice takes a slice of strings and returns a [][]string by calling stringToSlice on each string
func destructureSlice(slice []string) [][]string {
	var slices [][]string
	for _, s := range slice {
		slices = append(slices, stringToSlice(s))
	}
	return slices
}

// SevenSegmentDecoder is a struct with a map[int][]string and an int
type SevenSegmentDecoder struct {
	knownDecodings map[int][]string
	decodeCount    int
}

// DecodeOne is a function on SevenSegmentDecoder takes a [][]string and looks for a []string of length 2 and sets knownDecodings[1] and increment decodeCount
func (s *SevenSegmentDecoder) DecodeOne(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 2 {
			s.knownDecodings[1] = slice
			s.decodeCount++
		}
	}
}

// DecodeTwo takes a [][]string and sets knownDecodings[10]
func (s *SevenSegmentDecoder) DecodeTwo(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 10 {
			s.knownDecodings[10] = slice
		}
	}
}

// DecodeThree takes a [][]string and sets knownDecodings[10]
func (s *SevenSegmentDecoder) DecodeThree(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 10 {
			s.knownDecodings[10] = slice
		}
	}
}

// DecodeFour takes a [][]string and sets knownDecodings[4] if slice has length 4 and increment decodeCount
func (s *SevenSegmentDecoder) DecodeFour(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 4 {
			s.knownDecodings[4] = slice
			s.decodeCount++
		}
	}
}

// DecodeFive takes a [][]string and sets knownDecodings[10]
func (s *SevenSegmentDecoder) DecodeFive(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 10 {
			s.knownDecodings[10] = slice
		}
	}
}

// DecodeSix takes a [][]string and sets knownDecodings[10]
func (s *SevenSegmentDecoder) DecodeSix(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 10 {
			s.knownDecodings[10] = slice
		}
	}
}

// DecodeSeven takes a [][]string and sets knownDecodings[7] if slice is length 3 and increment decodeCount
func (s *SevenSegmentDecoder) DecodeSeven(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 3 {
			s.knownDecodings[7] = slice
			s.decodeCount++
		}
	}
}

// DecodeEight takes a [][]string and sets knownDecodings[8] if slice is length 7 and increment decodeCount
func (s *SevenSegmentDecoder) DecodeEight(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 7 {
			s.knownDecodings[8] = slice
			s.decodeCount++
		}
	}
}

// DecodeNine takes a [][]string and sets knownDecodings[10]
func (s *SevenSegmentDecoder) DecodeNine(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 10 {
			s.knownDecodings[10] = slice
		}
	}
}

func main() {
	// Open input.txt
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Read lines into a slice
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Create slice of SevenSegmentDecoder
	var decoders []SevenSegmentDecoder

	// for line in lines
	//   fmt.print line
	//   _, time := parseLine(line)
	//   slices is destructureSlice(time)
	//   create a SevenSegmentDecoder
	//   DecodeOne(slices)
	//   DecodeFour(slices)
	//   DecodeSeven(slices)
	//   DecodeEight(slices)
	//   append SevenSegmentDecoder to decoders

	for _, line := range lines {
		_, time := parseLine(line)
		slices := destructureSlice(time)
		decoder := SevenSegmentDecoder{
			knownDecodings: make(map[int][]string),
		}
		decoder.DecodeOne(slices)
		decoder.DecodeFour(slices)
		decoder.DecodeSeven(slices)
		decoder.DecodeEight(slices)
		decoders = append(decoders, decoder)
	}

	// fmt.print divider
	fmt.Println("------------------------------------------------------")

	// sum is the number of keys across every decoder
	sum := 0

	// loop through decoders and sum the decodeCount
	for _, decoder := range decoders {
		sum += decoder.decodeCount
	}

	// fmt.print sum
	fmt.Println(sum)
}
