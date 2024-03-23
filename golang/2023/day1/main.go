package main

import (
	// "bufio"
	"fmt"
	"github.com/grosheth/Advent-of-code/golang/util"
	"strconv"
)

func main() {
	answer := part1(util.ReadFile("input.txt"))

	fmt.Println("Reponse:", answer)

}

func part1(file []string) []string {
	for _, each_ln := range file {
		var temp_list []int 
    for char := range each_ln {
  		if value, err := strconv.Atoi(string(char); err == nil {
        fmt.Println(char)
      }
    }

	}

	return file
}

//
// func part1(file string) []int {
// 	// first digit + last digit = 2 digit number
// 	// if one digit, double it
// 	// sum up all these values
// 	var numbers []int
// 	var digits []int
// 	for x := range file {
// 		// Validating if its the end of line
// 		if string(file[x]) != "" {
//
// 			// Validating if value is a number and adding FIRST and LAST digits to list
// 			if value, err := strconv.Atoi(string(file[x])); err == nil {
// 				fmt.Println(value)
// 				if len(digits) == 0 {
// 					digits = append(digits, value)
// 					digits = append(digits, value)
// 				} else {
// 					digits[1] = value
// 				}
// 			}
// 			numbers = digits
// 		}
// 	}
//
// 	return numbers
// }
