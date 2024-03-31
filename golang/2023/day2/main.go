package main

import (
	"fmt"
	"github.com/grosheth/Advent-of-code/golang/util"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// answer1 := part1(util.ReadFile("input.txt"))
	// fmt.Println("Part1:", answer1)

	answer2 := part2(util.ReadFile("sample.txt"))
	fmt.Println("Part2:", answer2)

}

// PART 2
// Get the biggest number of each color
// The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together.
// The power of the minimum set of cubes in game 1 is 48. In games 2-5 it was 12, 1560, 630, and 36, respectively.
// Adding up these five powers produces the sum 2286
func part2(file []string) int {
	valid := true
	answer := 0
	gameId := 0

	for index, game := range file {
		gameId = index + 1
		rgx := regexp.MustCompile(`:+`)
		gameSplit := rgx.Split(game, -1)
		// fmt.Println("gameSplit", gameSplit)

		for _, x := range gameSplit {
			rgx := regexp.MustCompile(`;+`)
			colorSplit := rgx.Split(x, -1)
			// fmt.Println("colorSplit", colorSplit)

			for _, y := range colorSplit {
				rgx := regexp.MustCompile(`,+`)
				resultSplit := rgx.Split(y, -1)
				// fmt.Println("resultSplit", resultSplit)

				for _, w := range resultSplit {
					colorList := []string{"green", "red", "blue", "Game"}
					rgx := regexp.MustCompile(`[0-9]+`)
					number := rgx.FindStringSubmatch(w)
					nombre, _ := strconv.Atoi(number[0])

					for _, color := range colorList {
						if strings.Contains(w, color) {
							switch check := color; check {
							case "green":
								if nombre > 13 {
									gameId = 0
								}
							case "blue":
								if nombre > 14 {
									gameId = 0
								}
							case "red":
								if nombre > 12 {
									gameId = 0
								}
							}
						}
					}
				}
			}
		}
		// Sum of Power
		if valid == true {
			// fmt.Println("Adding", gameId, "to", answer)
			answer += gameId
		}
		valid = true
	}
	return answer
}

func part1(file []string) int {
	valid := true
	answer := 0
	gameId := 0

	for index, game := range file {
		gameId = index + 1
		rgx := regexp.MustCompile(`:+`)
		gameSplit := rgx.Split(game, -1)

		for _, x := range gameSplit {
			rgx := regexp.MustCompile(`;+`)
			colorSplit := rgx.Split(x, -1)

			for _, y := range colorSplit {
				rgx := regexp.MustCompile(`,+`)
				resultSplit := rgx.Split(y, -1)

				for _, w := range resultSplit {
					colorList := []string{"green", "red", "blue", "Game"}
					rgx := regexp.MustCompile(`[0-9]+`)
					number := rgx.FindStringSubmatch(w)
					nombre, _ := strconv.Atoi(number[0])

					for _, color := range colorList {
						if strings.Contains(w, color) {
							switch check := color; check {
							case "green":
								if nombre > 13 {
									fmt.Println("game", gameId, "is not valid")
									gameId = 0
								}
							case "blue":
								if nombre > 14 {
									fmt.Println("game", gameId, "is not valid")
									gameId = 0
								}
							case "red":
								if nombre > 12 {
									fmt.Println("game", gameId, "is not valid")
									gameId = 0
								}
							}
						}
					}
				}
			}
		}
		if valid == true {
			answer += gameId
		}
		valid = true
	}
	return answer
}
