package main

import (
	"fmt"
	"github.com/grosheth/Advent-of-code/golang/util"
	"unicode"
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

	for pos, mid := range file {
		if mid == file[0] {
			down := file[pos+1]
			table := []string{
				mid,
				down,
			}
			getPartNumber(table)
		} else {
			if mid == file[len(file)-1] {
				up := file[pos-1]
				table := []string{
					up,
					mid,
				}
				getPartNumber(table)
			} else {
				up := file[pos-1]
				down := file[pos+1]
				table := []string{
					up,
					mid,
					down,
				}
				getPartNumber(table)
			}
		}
	}
	return answer
}

func getPartNumber(table []string) {
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
	for _, row := range table {
		for pos, char := range row {
			if unicode.IsDigit(char) {
				if pos > 0 {
					fmt.Println("|current char:", string(char), "| char before: ", string(table[0][pos]), "| current row |")
					if re[string(row[pos-1])] {
						fmt.Println("|current char:", string(char), "| char before: ", string(table[0][pos]), "| current row |")
					}
				} else {
					// check after current row
					if pos != len(row) {
						if re[string(row[pos+1])] {
							fmt.Println("check char after | current row |", string(char), string(table[0][pos]))
							fmt.Println(string(char), string(table[0][pos]))
						}
					}
				}
			}
		}
	}
}

// Part 2
// func part2(file []string) int {
// 	answer := 0
//
// 	return answer
// }

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
