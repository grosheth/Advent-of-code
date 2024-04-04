package main

import (
	"fmt"
	"github.com/grosheth/Advent-of-code/golang/util"
	"log"
	"regexp"
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

	// find symbols
	re, err := regexp.Compile(`[-!$%^&*()#_+|~={}\[\]:";'<@>?,\/]`)
	if err != nil {
		log.Fatal(err)
	}

	for _, x := range file {
		x = re.ReplaceAllString(x, " ")
		fmt.Println(x)
	}
	return answer
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
