package main

import (
	"fmt"
	"github.com/grosheth/Advent-of-code/golang/util"
	"strconv"
	"strings"
)

func main() {
	// answer1 := part1(util.ReadFile("input.txt"))
	// fmt.Println("Part1:", answer1)

	answer2 := part2(util.ReadFile("sample.txt"))
	fmt.Println("Part2:", answer2)

}

// func part1(file []string) int {
// 	var numbers_list []string
// 	answer := 0
//
// 	for _, each_ln := range file {
// 		var digits []int
// 		for _, char := range each_ln {
// 			if value, err := strconv.Atoi(string(char)); err == nil {
// 				digits = append(digits, value)
// 			}
// 		}
// 		numbers := fmt.Sprint(digits[0]) + fmt.Sprint(digits[len(digits)-1])
// 		numbers_list = append(numbers_list, fmt.Sprint(numbers))
// 	}
// 	for _, value := range numbers_list {
// 		num, err := strconv.Atoi(value)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		answer += num
// 	}
// 	return answer
// }

func part2(file []string) int {
	var numbers_list []string
	answer := 0

	for _, each_ln := range file {
		var digits []int
		var replacer = strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
		each_ln = replacer.Replace(each_ln)

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
