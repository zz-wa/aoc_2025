package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var location = "/home/butter-july/桌面/12_10.txt"

func main() {
	contents := ReadContent(location)
	newContent := SplitElement(contents)
	totalMinPresses := 0

	for _, value := range newContent {
		numberOfLights := NumberOfLights(value)
		buttons := SplitButtons(value, numberOfLights+3)
		lights := value[:numberOfLights+3]
		number := Initialise(lights, numberOfLights)

		minPresses := CheckAllCombination(number, buttons, len(buttons)-1)

		if minPresses != -1 {
			totalMinPresses += minPresses
		} else {
		}
	}

	fmt.Println(totalMinPresses)

}

// 先给. #设置一个值 。0 # 1

func Initialise(lights string, NUmberOfLights int) []int {
	var totalNumber []int
	for i := 1; i <= NUmberOfLights; i++ {
		if lights[i] == '.' {
			totalNumber = append(totalNumber, 0)
		} else if lights[i] == '#' {
			totalNumber = append(totalNumber, 1)
		}
	}
	return totalNumber
}

func NumberOfLights(content string) int {
	lights := strings.Split(content, "(")
	number := len(lights[0]) - 3
	return number
}

func ButtonToVector(buttonStr string, NumberOfLights int) []int {
	vector := make([]int, NumberOfLights)
	buttonStr = strings.Trim(buttonStr, "()")
	nums := strings.Split(buttonStr, ",")
	for _, numStr := range nums {

		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			continue
		}
		if num >= 0 && num < NumberOfLights {
			vector[num] = 1
		}
	}
	return vector
}
func CountTwo(a []int, b []int) []int {
	n := len(a)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = (a[i] + b[i]) % 2
	}
	return result
}
func Equal(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func SplitButtons(lines string, n int) []string {
	contents := lines[n:]
	button := strings.Split(contents, " ")
	return button
}
func checkCombination(targetK int, startIndex int, currentResult []int, targetLights []int, buttonVectors [][]int) bool {

	// 1:要求的按钮已经足够了
	if targetK == 0 {
		return Equal(currentResult, targetLights)
	}

	// 2:还不够
	remainingButtons := len(buttonVectors) - startIndex
	if targetK > remainingButtons || startIndex >= len(buttonVectors) {
		return false
	}

	// 1. 选择 buttonVectors[startIndex]（将当前按钮加入组合）：
	// targetK-1 (数量减 1), startIndex+1 (指针前进）
	newResult := CountTwo(currentResult, buttonVectors[startIndex])
	if checkCombination(targetK-1, startIndex+1, newResult, targetLights, buttonVectors) {
		return true
	}

	// 2. 不选择 buttonVectors[startIndex]（跳过当前按钮）：
	// targetK 不变, startIndex+1 (指针前进)
	if checkCombination(targetK, startIndex+1, currentResult, targetLights, buttonVectors) {
		return true
	}

	return false
}

func CheckAllCombination(LightsInNumber []int, buttons []string, NumberOfButtons int) int {
	nLights := len(LightsInNumber)

	buttonVectors := [][]int{}
	for _, button := range buttons {
		vector := ButtonToVector(button, nLights)
		isZero := true
		for _, v := range vector {
			if v == 1 {
				isZero = false
				break
			}
		}
		if !isZero {
			buttonVectors = append(buttonVectors, vector)
		}
	}
	m := len(buttonVectors)

	isTargetZero := true
	for _, v := range LightsInNumber {
		if v == 1 {
			isTargetZero = false
			break
		}
	}
	if isTargetZero {
		return 0
	}

	for k := 1; k <= m; k++ {
		initialZeroVector := make([]int, nLights)

		if checkCombination(k, 0, initialZeroVector, LightsInNumber, buttonVectors) {
			return k
		}
	}
	return -1
}
func SplitElement(contents []string) []string {
	var newContents []string
	for _, content := range contents {
		lines := strings.Split(content, "{")
		newContents = append(newContents, lines[0])
	}

	return newContents

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
