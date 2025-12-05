package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var location = "12_5.txt"

type Range struct {
	Head int
	Tail int
}

func main() {
	file := ReadContent(location)
	index := FindEmptyLine(file)
	scopes := file[:index]
	ranges := ToRanges(scopes)
	numbers := MergeAndCount(ranges)
	fmt.Println(numbers)
}
func ToRanges(scopes []string) []Range {
	var ranges []Range
	for _, scope := range scopes {
		head, tail := Spilt(scope)
		ranges = append(ranges, Range{Head: head, Tail: tail})
	}
	return ranges
}
func MergeAndCount(ranges []Range) int {
	if len(ranges) == 0 {
		return 0
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Head < ranges[j].Head
	})

	var totalnumber int

	currentHead := ranges[0].Head
	currentTail := ranges[0].Tail

	for i := 1; i < len(ranges); i++ {
		nextRange := ranges[i]

		if nextRange.Head <= currentTail+1 {
			if nextRange.Tail > currentTail {
				currentTail = nextRange.Tail
			}
		} else {
			totalnumber += currentTail - currentHead + 1
			currentHead = nextRange.Head
			currentTail = nextRange.Tail
		}
	}

	totalnumber += currentTail - currentHead + 1

	return totalnumber
}
func Spilt(scope string) (int, int) {
	s := strings.Split(scope, "-")
	head, err := strconv.Atoi(s[0])
	if err != nil {
		log.Fatal(err)
	}
	tail, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}
	return head, tail
}
func FindEmptyLine(contents []string) int {

	for i, value := range contents {
		if value == "" {
			return i
		}
	}
	return -1
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
