package main

import (
	"fmt"
	"github.com/grosheth/Advent-of-code/golang/util"
)

func main() {
	answer1 := part1(util.ReadFile("sample.txt"))
	fmt.Println("Part1:", answer1)

	// answer2 := part2(util.ReadFile("input.txt"))
	// fmt.Println("Part2:", answer2)

}

func part1(file []string) []string {
	// answer := 0
	// Split input by games
	fmt.Println(file)
	// maximum number cannot be shown for one color
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	// count number of game that are possible
	// Sum the ID's of each possible game
	return file
}

// func part2() {
//
// }
