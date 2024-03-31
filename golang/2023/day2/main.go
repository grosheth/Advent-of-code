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
		// fmt.Println(game)

		// get Game ID
		// gameId, _ := regexp.Compile("[0-9]+")
		// gameSplit := gameId.FindStringSubmatch(string(game))
		// fmt.Println("gameId", gameSplit)

		// Get colors
		a := regexp.MustCompile(`:+`)
		test := a.Split(game, -1)
		fmt.Println("test", test)
		for _, x := range test {
			// x[0] == game number
			// x[*] == game results
			b := regexp.MustCompile(`;+`)
			test1 := b.Split(x, -1)
			// fmt.Println(test1)
			for _, y := range test1 {
				b := regexp.MustCompile(`,+`)
				test2 := b.Split(y, -1)
				fmt.Println(test2)
			}
		}
	}

	// only 12 red cubes, 13 green cubes, and 14 blue cubes
	// count number of game that are possible
	// Sum the ID's of each possible game
	return file
}

// func part2() {
//
// }
