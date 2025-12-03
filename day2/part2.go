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
var location = "/home/butter-july/桌面/12_2.txt"

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



func FindNumber(s string) bool {
	lens := len(s)
	if lens < 2 {
		return false
	}
	for i := 1; i <= lens/2; i++ {
		if lens%i == 0 {
			PreString := s[:i]
			repeats := lens / i
			if strings.Repeat(PreString, repeats) == s {
				return true
			}
		}
	}
	return false
}
func valid(number []string) int {
	invalid_number := 0
	head, tail := GetHeadAndTail(number)
	for i := head; i <= tail; i++ {
		s := strconv.Itoa(i)
		if FindNumber(s) {
			invalid_number += i

		}
	}
	return invalid_number

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
