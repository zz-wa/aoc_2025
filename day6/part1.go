package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var location = "12_6.txt"

func main() {
	contents := ReadContent(location)
	fmt.Println(contents)
	Location := FindOperatorLocation(contents)
	Number := contents[:Location]
	operators := contents[Location:]
	//sum := 0
	Numbers := ToInt(Number)
	operator := GetOperator(operators)

	sum := Calculate(Numbers, operator)
	fmt.Println(sum)

}

func GetOperator(operatorsLine []string) []string {
	var operators []string
	line := operatorsLine[0]
	for _, char := range line {
		op := string(char)
		if op == "*" || op == "+" {
			operators = append(operators, op)
		}
	}
	return operators
}

func Calculate(Numbers [][]int, operators []string) int {
	numberOfProblem := len(Numbers[0])
	numbers := make([][]int, numberOfProblem)
	for i := 0; i < numberOfProblem; i++ {
		for j := 0; j < len(Numbers); j++ {
			numbers[i] = append(numbers[i], Numbers[j][i])
		}
	}
	total := 0
	for i, number := range numbers {
		operator := operators[i]
		currentResult := number[0]
		for j := 1; j < len(number); j++ {
			n := number[j]
			if operator == "*" {
				currentResult *= n
			} else if operator == "+" {
				currentResult += n
			}
		}
		total += currentResult
	}
	return total
}

func ToInt(Number []string) [][]int {
	var totalNumbers [][]int
	for _, SNumber := range Number {
		numbers := Spilt(SNumber)
		totalNumbers = append(totalNumbers, numbers)
	}
	return totalNumbers
}

func Spilt(SNumbers string) []int {
	n := strings.Split(SNumbers, " ")
	var numbers []int
	for _, SNumber := range n {
		number, _ := strconv.Atoi(SNumber)
		if number != 0 {
			numbers = append(numbers, number)
		}

	}
	return numbers
}

func FindOperatorLocation(contents []string) int {
	for i, content := range contents {
		for _, value := range content {
			if string(value) == "*" || string(value) == "+" {
				return i
			}
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
