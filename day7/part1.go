package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var location = "12_7.txt"

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
		}
	}
	beams := map[int]bool{StartIndex: true}
	total := 0
	for _, row := range contents[1:] {

		nextBeams := make(map[int]bool)
		split := make(map[int]bool)

		for beam := range beams {
			if beam >= 0 && beam < len(row) {
				char := row[beam]
				if string(char) == "^" {
					split[beam] = true
				} else {
					nextBeams[beam] = true
				}
			}
		}
		total += len(split)
		for spiltCol := range split {
			nextBeams[spiltCol+1] = true
			nextBeams[spiltCol-1] = true
		}
		beams = nextBeams
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
