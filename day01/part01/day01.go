package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("realInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")
	lenLines := len(lines) - 1

	var total int
	for i := 0; i < lenLines; i++ {
		currentLine := lines[i]
		fmt.Println(currentLine)
		lineTotal := make([]int, 0)
		for _, j := range currentLine {
			if j >= '0' && j <= '9' {
				lineTotal = append(lineTotal, int(j-'0'))
			}
		}
		total += int(lineTotal[0]*10 + lineTotal[len(lineTotal)-1])
	}
	fmt.Println(total)
}
