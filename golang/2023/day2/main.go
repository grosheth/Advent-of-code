package main

import (
	"fmt"
	"github.com/grosheth/Advent-of-code/golang/util"
	"regexp"
	// "strings"
)

func main() {
	answer1 := part1(util.ReadFile("sample.txt"))
	fmt.Println("Part1:", answer1)

	// answer2 := part2(util.ReadFile("input.txt"))
	// fmt.Println("Part2:", answer2)

}

func part1(file []string) []string {
	answer := 0
	// maximum number cannot be obtained for one color
	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	green := 13
	red := 12
	blue := 14
	fmt.Println(answer, green, red, blue)
	// Split input by games
	for _, game := range file {
		fmt.Println(game)
		gameSplit, err := regexp.MatchString(game, "[0-9]")
		println(err)
		fmt.Println(gameSplit)
	}

	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	// count number of game that are possible
	// Sum the ID's of each possible game
	return file
}

// func part2() {
//
// }
