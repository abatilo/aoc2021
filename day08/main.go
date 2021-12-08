package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// sliceIntersection takes two slices and returns a slice of the intersection
func sliceIntersection(slice1, slice2 []string) []string {
	var intersection []string
	for _, s1 := range slice1 {
		for _, s2 := range slice2 {
			if s1 == s2 {
				intersection = append(intersection, s1)
			}
		}
	}
	return intersection
}

// sliceDifference takes two slices and returns a slice of what's in slice1 but not in slice2 and what's in slice2 but not in slice1 combined into a single slice
func sliceDifference(slice1, slice2 []string) []string {
	var difference []string
	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
			}
		}
		if !found {
			difference = append(difference, s1)
		}
	}
	for _, s2 := range slice2 {
		found := false
		for _, s1 := range slice1 {
			if s1 == s2 {
				found = true
			}
		}
		if !found {
			difference = append(difference, s2)
		}
	}
	return difference
}

// sliceEquals takes two unsorted slices and returns true if they have the same elements
func sliceEquals(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for _, s1 := range slice1 {
		found := false
		for _, s2 := range slice2 {
			if s1 == s2 {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

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

// DecodeZero takes a [][]string
// find all slices of length 6
// do a slice intersection with knownDecodings[1]
// do a slice intersection with knownDecodings[4]
// if intersection1 sliceEquals knownDecodings[1] and not intersection2 sliceEqual knownDecodings[4] then set knownDecodings[0]
func (s *SevenSegmentDecoder) DecodeZero(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 6 {
			intersection1 := sliceIntersection(slice, s.knownDecodings[1])
			intersection2 := sliceIntersection(slice, s.knownDecodings[4])
			if sliceEquals(intersection1, s.knownDecodings[1]) && !sliceEquals(intersection2, s.knownDecodings[4]) {
				s.knownDecodings[0] = slice
				s.decodeCount++
			}
		}
	}
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

// DecodeTwo takes a [][]string and sets knownDecodings[2]
// find all slices of length 5
// do a set intersection with knownDecodings[1]
// do a slice difference with knownDecodings[6]
// if the length of difference > 2 set knownDecodings of 2 and intersection !sliceEquals knownDecodings[1]
func (s *SevenSegmentDecoder) DecodeTwo(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 5 {
			intersection := sliceIntersection(slice, s.knownDecodings[1])
			difference := sliceDifference(slice, s.knownDecodings[6])
			if len(difference) > 2 && !sliceEquals(intersection, s.knownDecodings[1]) {
				s.knownDecodings[2] = slice
				s.decodeCount++
			}
		}
	}
}

// DecodeThree takes a [][]string and sets knownDecodings[10]
// find all slices of length 5
// do a slice intersection with knownDecodings[1] and if the intersection is sliceEqual knownDecodings[1] then set knownDecodings[3]
func (s *SevenSegmentDecoder) DecodeThree(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 5 {
			intersection := sliceIntersection(slice, s.knownDecodings[1])
			if sliceEquals(intersection, s.knownDecodings[1]) {
				s.knownDecodings[3] = slice
				s.decodeCount++
			}
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

// DecodeFive takes a [][]string and sets knownDecodings[5]
// find all slices of length 5
// do a slice difference with knownDecodings[6] and if the length of difference is 1 then set knownDecodings[5]
func (s *SevenSegmentDecoder) DecodeFive(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 5 {
			difference := sliceDifference(slice, s.knownDecodings[6])
			if len(difference) == 1 {
				s.knownDecodings[5] = slice
				s.decodeCount++
			}
		}
	}
}

// ecodeSix takes a [][]string and sets knownDecodings[6]
// find all slices of length 6
// do a slice intersection with knownDecodings[1]
// do a slice intersection with knownDecodings[4]
// if intersection1 != knownDecodings[1] and intersection2 != knownDecodings[4] then set knownDecodings[6]
func (s *SevenSegmentDecoder) DecodeSix(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 6 {
			intersection1 := sliceIntersection(slice, s.knownDecodings[1])
			intersection2 := sliceIntersection(slice, s.knownDecodings[4])
			if !sliceEquals(intersection1, s.knownDecodings[1]) && !sliceEquals(intersection2, s.knownDecodings[4]) {
				s.knownDecodings[6] = slice
				s.decodeCount++
			}
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

// DecodeNine takes a [][]string
// find all slices of length 6
// do a slice intersection with knownDecodings[4]
// if intersection == knownDecodings[4] then set knownDecodings[9]
func (s *SevenSegmentDecoder) DecodeNine(slices [][]string) {
	for _, slice := range slices {
		if len(slice) == 6 {
			intersection := sliceIntersection(slice, s.knownDecodings[4])
			if sliceEquals(intersection, s.knownDecodings[4]) {
				s.knownDecodings[9] = slice
				s.decodeCount++
			}
		}
	}
}

// DecodeAll calls DecodeOne, DecodeFour, DecodeSeven, DecodeEight, DecodeNine, DecodeZero, DecodeSix, DecodeThree, DecodeFive, DecodeTwo
func (s *SevenSegmentDecoder) DecodeAll(slices [][]string) {
	s.DecodeOne(slices)
	s.DecodeFour(slices)
	s.DecodeSeven(slices)
	s.DecodeEight(slices)
	s.DecodeNine(slices)
	s.DecodeZero(slices)
	s.DecodeSix(slices)
	s.DecodeThree(slices)
	s.DecodeFive(slices)
	s.DecodeTwo(slices)
}

// FindMatchingSlice takes a a slice and checks it against each value in knownDecodings
// if the slice is sliceEqual to a knownDecoding then return the key of the knownDecoding
func (s *SevenSegmentDecoder) FindMatchingSlice(slice []string) int {
	for key, known := range s.knownDecodings {
		if sliceEquals(slice, known) {
			return key
		}
	}
	return -1
}

// JoinCodes takes a [][]string and returns a string of FindMatchingSlice of each string concatenated together
func (s *SevenSegmentDecoder) JoinCodes(slices [][]string) string {
	var codes []string
	for _, slice := range slices {
		code := s.FindMatchingSlice(slice)
		if code != -1 {
			codes = append(codes, strconv.Itoa(code))
		}
	}
	return strings.Join(codes, "")
}

// parseInt takes a string and returns an int
func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
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

	// Create slice of strings called codes
	var codes []string

	// for line in lines
	//   fmt.print line
	//   raw, code := parseLine(line)
	//   slices is destructureSlice(raw)
	//   create new SevenSegmentDecoder with empty map and decodeCount 0
	//   DecodeAll(slices)
	//   append SevenSegmentDecoder to decoders
	//   destructuredCodes is destructureSlice(code)
	//   append decoder.JoinCodes(destructuredCodes) to codes

	for _, line := range lines {
		raw, code := parseLine(line)
		slices := destructureSlice(raw)
		decoder := SevenSegmentDecoder{make(map[int][]string), 0}
		decoder.DecodeAll(slices)
		decoders = append(decoders, decoder)
		destructuredCodes := destructureSlice(code)
		codes = append(codes, decoder.JoinCodes(destructuredCodes))
	}

	// fmt.print divider
	fmt.Println("------------------------------------------------------")

	// convert codes to ints and sum
	for _, code := range codes {
		fmt.Println(code)
	}

	// parse codes to ints
	// sum the ints
	// print sum
	sum := 0
	for _, code := range codes {
		sum += parseInt(code)
	}
	fmt.Println(sum)
}
