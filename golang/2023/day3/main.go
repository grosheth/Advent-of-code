package main

import (
	"fmt"
	"unicode"

	"github.com/grosheth/Advent-of-code/golang/util"
)

func main() {
	answer1 := part1(util.ReadFile("sample.txt"))
	fmt.Println("Part1:", answer1)

	// answer2 := part2(util.ReadFile("input.txt"))
	// fmt.Println("Part2:", answer2)
}

// Part 1
// Add the part numbers
// part numbers are adjacent to symbols
func part1(file []string) int {
	answer := 0

	getPartNumber(file)

	return answer
}

func getPartNumber(table []string) {
	part := []int{}
	re := map[string]bool{
		"[": true,
		"-": true,
		"!": true,
		"$": true,
		"*": true,
		"%": true,
		"^": true,
		"&": true,
		"(": true,
		")": true,
		"#": true,
		"_": true,
		"+": true,
		// |~={}\[\]:";'<@>?,\/]`}
	}
	// check with same row
	for position, row := range table {
		for charPos, char := range row {
			if unicode.IsDigit(char) {
				// Build number before analysing other values
				// Get all numbers next to char when char is a number
				// Get all numbers next to char when char is a number
				// Get all numbers next to char when char is a number
				// Get all numbers next to char when char is a number
				// Get all numbers next to char when char is a number
				if charPos > 0 {
					if ok := re[string(row[charPos-1])]; ok {
						fmt.Println("|current char:", string(char), "| char before: ", string(row[charPos-1]))
						part = append(part, int(char))
						break
					}
				}
				// Checking char after on same row
				if charPos != len(row) {
					if ok := re[string(row[charPos+1])]; ok {
						fmt.Println("|current char:", string(char), "| char after: ", string(row[charPos+1]))
						part = append(part, int(char))
						break
					}
				}
				// Checking for char on upper row
				if position > 0 {
					if ok := re[string(table[position-1][charPos])]; ok {
						fmt.Println("|current char:", string(char), "| char over: ", string(table[position-1][charPos]))
						part = append(part, int(char))
						break
					}
					if charPos > 0 {
						if ok := re[string(table[position-1][charPos-1])]; ok {
							fmt.Println("|current char:", string(char), "| char over left: ", string(table[position-1][charPos-1]))
							part = append(part, int(char))
							break
						}
					}
					if charPos != len(row) {
						if ok := re[string(table[position-1][charPos+1])]; ok {
							fmt.Println("|current char:", string(char), "| char over right: ", string(table[position-1][charPos+1]))
							part = append(part, int(char))
							break
						}
					}
				}
				// Checking for char on lower row
				if position != len(table)-1 {
					if ok := re[string(table[position+1][charPos])]; ok {
						fmt.Println("|current char:", string(char), "| char under: ", string(table[position+1][charPos]))
						part = append(part, int(char))
						break
					}
					if charPos > 0 {
						if ok := re[string(table[position][charPos-1])]; ok {
							fmt.Println("|current char:", string(char), "| char under left: ", string(table[position+1][charPos-1]))
							part = append(part, int(char))
							break
						}
					}
					if charPos != len(row) {
						if ok := re[string(table[position][charPos+1])]; ok {
							fmt.Println("|current char:", string(char), "| char under right: ", string(table[position+1][charPos+1]))
							part = append(part, int(char))
							break
						}
					}
				}
			}
		}
	}
}

// remove all symbols
// func part1(file []string) int {
// 	answer := 0
//
// 	re, err := regexp.Compile(`[^\w]`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, x := range file {
// 		x = re.ReplaceAllString(x, "")
// 		fmt.Println(x)
// 	}
// 	return answer
// }

// Part 2
// func part2(file []string) int {
// 	answer := 0
//
// 	return answer
// }
