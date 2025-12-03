package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var contents []string
var location = "12_3.txt"

func main() {
	contents = ReadContent(location)
	var totalMaxJolts int
	for _, line := range contents {
		totalMaxJolts += FindJolts(line, 2)
	}
	fmt.Println(totalMaxJolts)
}
func FindJolts(bank string, digits int) int {
	output := 0
	index := -1
	for i := digits; i > 0; i-- {
		var jolts int
		jolts, index = FindMax(bank, index+1, len(bank)-i+1)
		output = output*10 + jolts
	}
	return output
}
func FindMax(list string, start, end int) (int, int) {
	var maxJolts uint8
	var index int
	for i := start; i < end; i++ {
		if i < 0 || i >= len(list) {
			continue
		}
		if list[i] > maxJolts {
			maxJolts = list[i]
			index = i
		}
	}
	if maxJolts == 0 {
		return 0, -1
	}
	return int(maxJolts - '0'), index
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
