package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var contents []string
var ranges string
var location = "12_2.txt"

func main() {
	contents = ReadContent(location)

	numbers := Split(contents)

	total := 0
	for _, t_number := range numbers {
		number := strings.Split(t_number, "-")
		total += valid(number)
	}
	fmt.Println(total)
}
func valid(number []string) int {
	invalid_number := 0

	head, tail := GetHeadAndTail(number)
	for i := head; i <= tail; i++ {
		s := strconv.Itoa(i)
		invalid_number += FindNumber(i, s)
	}

	return invalid_number
}

func FindNumber(i int, s string) int {
	number := 0
	if len(s)%2 == 0 { //个数符合要求
		mid := len(s) / 2
		a := s[:mid]
		b := s[mid:]
		if a == b {
			number += i
		}
	}
	return number

}

func ReadContent(location string) []string {
	var c []string
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		content := scanner.Text()
		c = append(c, content)
	}
	return c
}

func Split(c []string) []string {
	ranges = strings.Join(c, "")
	numbers := strings.Split(ranges, ",")
	return numbers
}
func GetHeadAndTail(n []string) (int, int) {
	head, err := strconv.Atoi(n[0])
	if err != nil {
		log.Fatal(err)
	}

	tail, err := strconv.Atoi(n[1])
	if err != nil {
		log.Fatal(err)
	}
	return head, tail

}
