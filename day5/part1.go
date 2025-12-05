package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var location = "12_5.txt"

func main() {
	file := ReadContent(location)

	index := FindEmptyLine(file)
	scopes := file[:index]
	value := file[index+1:]
	number := IDNumber(scopes, value)
	fmt.Println(number)
}

func IDNumber(scopes, value []string) int {
	AllIDsLocation := FindAllIDsLocation(scopes, value)
	uniqueLocation := make(map[int]bool)

	for _, IDLocation := range AllIDsLocation {
		uniqueLocation[IDLocation] = true
	}
	return len(uniqueLocation)
}
func FindAllIDsLocation(scopes, value []string) []int {
	var locations []int
  
	for _, scope := range scopes {
		head, tail := Spilt(scope)
		for i, Snumber := range value {
			number, err := strconv.Atoi(Snumber)
			if err != nil {
				log.Fatal(err)
			}
			if number >= head && number <= tail {
				locations = append(locations, i)
			}
		}
	}

	return locations
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
