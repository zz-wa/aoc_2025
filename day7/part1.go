package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var location = "/home/butter-july/桌面/12_7.txt"

func main() {
	contents := ReadContent(location)
	total := Count(contents)
	fmt.Println(total)
}
func Count(contents []string) int {
	StartIndex := 0
	for index, value := range contents[0] {
		if string(value) == "S" {
			StartIndex = index
			break
		}
	}

	width := len(contents[0])
	curr := make([]int, width)
	curr[StartIndex] = 1

	total := 0
	for _, row := range contents[1:] {
		nextCurr := make([]int, width)
		for col := 0; col < width; col++ {
			if curr[col] > 0 {
				char := row[col]
				if char == '^' {
					total++
					if col > 0 {
						nextCurr[col-1] += curr[col]
					}
					if col < width-1 {
						nextCurr[col+1] += curr[col]
					}

				} else {
					nextCurr[col] += curr[col]
				}
			}
		}
		curr = nextCurr
	}
	return total
}

func ReadContent(location string) []string {
	var c []string
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		content := scanner.Text()
		c = append(c, content)
	}
	return c
}
