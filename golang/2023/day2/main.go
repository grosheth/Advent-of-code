package main

import (
	"fmt"
	"github.com/grosheth/Advent-of-code/golang/util"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	answer1 := part1(util.ReadFile("input.txt"))
	fmt.Println("Part1:", answer1)

	answer2 := part2(util.ReadFile("input.txt"))
	fmt.Println("Part2:", answer2)

}

// Part 1
func part1(file []string) int {
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

// PART 2
func part2(file []string) int {
	answer := 0

	for _, game := range file {
		redNombre := 0
		greenNombre := 0
		blueNombre := 0
		rgx := regexp.MustCompile(`:+`)
		gameSplit := rgx.Split(game, -1)

		for _, x := range gameSplit {
			rgx := regexp.MustCompile(`;+`)
			colorSplit := rgx.Split(x, -1)

			for _, y := range colorSplit {
				rgx := regexp.MustCompile(`,+`)
				resultSplit := rgx.Split(y, -1)

				for _, w := range resultSplit {
					colorList := []string{"green", "red", "blue"}
					rgx := regexp.MustCompile(`[0-9]+`)
					number := rgx.FindStringSubmatch(w)
					nombre, _ := strconv.Atoi(number[0])

					for _, color := range colorList {
						if strings.Contains(w, color) {
							switch check := color; check {
							case "green":
								if nombre > greenNombre {
									greenNombre = nombre
								}
							case "blue":
								if nombre > blueNombre {
									blueNombre = nombre
								}
							case "red":
								if nombre > redNombre {
									redNombre = nombre
								}
							}
						}
					}
				}
			}
		}
		power := greenNombre * blueNombre * redNombre
		answer += power
	}
	return answer
}
