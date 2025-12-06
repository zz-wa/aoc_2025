package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var location = "12_6.txt"

func main() {
	number := IndividualGrandAnswersTotals()
	fmt.Println(number)
}
func IndividualGrandAnswersTotals() int {
	allLines := ReadContent(location)
	if len(allLines) > 0 && allLines[len(allLines)-1] == "" {
		allLines = allLines[:len(allLines)-1]
	}

	if len(allLines) == 0 {
		return 0
	}

	bottomOperationLine := allLines[len(allLines)-1]
	operators := make([]byte, 0)
	operatorColumnStarts := make([]int, 0)

	for columnIndex := 0; columnIndex < len(bottomOperationLine); columnIndex++ {
		if bottomOperationLine[columnIndex] != ' ' {
			operators = append(operators, bottomOperationLine[columnIndex])
			operatorColumnStarts = append(operatorColumnStarts, columnIndex)
		}
	}
	operatorColumnStarts = append(operatorColumnStarts, len(bottomOperationLine))

	dataRowsCount := len(allLines) - 1
	if dataRowsCount <= 0 {
		return 0
	}

	totalAnswerSum := 0

	for operatorIdx := 0; operatorIdx < len(operators); operatorIdx++ {
		currentOperator := operators[operatorIdx]
		verticalColumnResult := 0

		if currentOperator == '*' {
			verticalColumnResult = 1
		}

		groupStartColumn := operatorColumnStarts[operatorIdx]
		groupEndColumn := operatorColumnStarts[operatorIdx+1]

		if operatorIdx < len(operators)-1 {
			groupEndColumn -= 1
		}

		for currentColumn := groupStartColumn; currentColumn < groupEndColumn; currentColumn++ {
			verticalNumber := 0

			for rowIndex := 0; rowIndex < dataRowsCount; rowIndex++ {
				character := allLines[rowIndex][currentColumn]
				if character != ' ' {
					verticalNumber = verticalNumber*10 + int(character-'0')
				}
			}

			if currentOperator == '*' {
				verticalColumnResult *= verticalNumber
			} else {
				verticalColumnResult += verticalNumber
			}
		}

		totalAnswerSum += verticalColumnResult
	}

	return totalAnswerSum
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
