package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/grosheth/Advent-of-code/golang/util"
)

func main() {
	answer1 := part1(util.ReadFile("input.txt"))
	fmt.Println("Part1:", answer1)

	answer2 := part2(util.ReadFile("input.txt"))
	fmt.Println("Part2:", answer2)
}

func part1(file []string) int {
	var numbers_list []string
	answer := 0

	for _, each_ln := range file {
		var digits []int
		for _, char := range each_ln {
			if value, err := strconv.Atoi(string(char)); err == nil {
				digits = append(digits, value)
			}
		}
		numbers := fmt.Sprint(digits[0]) + fmt.Sprint(digits[len(digits)-1])
		numbers_list = append(numbers_list, fmt.Sprint(numbers))
	}
	for _, value := range numbers_list {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
		}
		answer += num
	}
	return answer
}

func part2(file []string) int {
	var numbers_list []string
	answer := 0
	for _, each_ln := range file {

		var digits []int
		replace_list := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		for _, x := range replace_list {
			if strings.Contains(each_ln, x) {
				switch check := x; check {
				case "one":
					each_ln = strings.Replace(each_ln, "one", "o1ne", -1)
				case "two":
					each_ln = strings.Replace(each_ln, "two", "t2wo", -1)
				case "three":
					each_ln = strings.Replace(each_ln, "three", "t3hree", -1)
				case "four":
					each_ln = strings.Replace(each_ln, "four", "f4our", -1)
				case "five":
					each_ln = strings.Replace(each_ln, "five", "f5ive", -1)
				case "six":
					each_ln = strings.Replace(each_ln, "six", "s6ix", -1)
				case "seven":
					each_ln = strings.Replace(each_ln, "seven", "s7even", -1)
				case "eight":
					each_ln = strings.Replace(each_ln, "eight", "ei8ght", -1)
				case "nine":
					each_ln = strings.Replace(each_ln, "nine", "n9ine", -1)
				}
			}
		}
		for _, char := range each_ln {
			if value, err := strconv.Atoi(string(char)); err == nil {
				digits = append(digits, value)
			}
		}
		numbers := fmt.Sprint(digits[0]) + fmt.Sprint(digits[len(digits)-1])
		numbers_list = append(numbers_list, fmt.Sprint(numbers))
	}

	for _, value := range numbers_list {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
		}
		answer += num
	}
	return answer
}
