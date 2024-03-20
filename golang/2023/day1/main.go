package main

import (
	"bufio"
	"fmt"
	"github.com/grosheth/Advent-of-code"
	"log"
	"os"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	fmt.Println(lines[0])

}
