package util

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(pathFromCaller string) []string {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	// for _, each_ln := range text {
	// 	fmt.Println(each_ln)
	// }
	return text
}
