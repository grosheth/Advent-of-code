package main

import (
	"fmt"
	"github.com/grosheth/Advent-of-code/golang/util"
	"regexp"
	"strings"
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
		// gameid, _ := regexp.compile("[0-9]+")
		// gamesplit := gameid.findstringsubmatch(string(game))
		// fmt.Println("gameId", gameSplit)

		// Get colors
		rgx := regexp.MustCompile(`:+`)
		gameSplit := rgx.Split(game, -1)
		fmt.Println("gameSplit", gameSplit)

		for _, x := range gameSplit {

			rgx := regexp.MustCompile(`;+`)
			colorSplit := rgx.Split(x, -1)
			fmt.Println("colorSplit", colorSplit)

			for _, y := range colorSplit {

				rgx := regexp.MustCompile(`,+`)
				resultSplit := rgx.Split(y, -1)
				fmt.Println("resultSplit", resultSplit)

				for _, w := range resultSplit {

					colorList := []string{"green", "red", "blue"}
					for _, color := range colorList {
						if strings.Contains(w, color) {
							switch check := color; check {
							case "green":
								rgx := regexp.MustCompile(`[0-9]`)
								number := rgx.Split(y, -1)
								fmt.Println(number)
							case "blue":
								rgx := regexp.MustCompile(`,+`)
								number := rgx.Split(y, -1)
								fmt.Println(number)
							case "red":
								rgx := regexp.MustCompile(`,+`)
								number := rgx.Split(y, -1)
								fmt.Println(number)
							}
						}
					}
				}
				// only 12 red cubes, 13 green cubes, and 14 blue cubes
				// count number of game that are possible
				// Sum the ID's of each possible game resultSplit[0]

			}
		}
	}
	return file
}

// func part2() {
//
// }
