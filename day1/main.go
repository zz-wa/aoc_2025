package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var rotations []string
var location = "12_1.txt"
var initialNumber = 50
var currentNumber int
var passwords int

func main() {

	rotations = ReadContent(location)

	for _, rotation := range rotations {

		number := StrToInt(rotation)

		passwords += number / 100

		numberRotation := number % 100 //小于99的数字

		passwords = CalculatePWD(rune(rotation[0]), numberRotation)

		initialNumber = currentNumber
	}

	fmt.Println(passwords)

}

func ReadContent(name string) []string {
	var r []string
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		rotation := scanner.Text()
		r = append(r, rotation)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return r
}

func StrToInt(rotation string) int {
	str := rotation[1:]

	number, _ := strconv.Atoi(str)
	return number

}

func CalculatePWD(direction rune, numberRotation int) int {
	if direction == 'L' {

		currentNumber = (initialNumber - numberRotation) % 100
		if currentNumber <= 0 {
			if initialNumber != 0 {
				passwords += 1
			}
			if currentNumber != 0 {
				currentNumber += 100
			}

		}
	} else {
		currentNumber = (initialNumber + numberRotation) % 100
		if initialNumber != 0 && (numberRotation+initialNumber) >= 100 {
			passwords += 1
		}
	}
	return passwords
}
