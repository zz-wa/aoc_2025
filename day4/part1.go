package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var contents []string
var location = "/home/butter-july/桌面/12_4.txt"

func main() {
	contents = ReadContent(location)
	rolls := GetRolls(contents)
	fmt.Println(rolls)
}

func GetRolls(content []string) int {
	rows := len(contents)
	cols := len(contents[0])
	SumRolls := 0

	offsets := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if content[row][col] == '@' {
				NumberOfAT := 0
				for _, offset := range offsets {
					dr, dc := offset[0], offset[1]
					nr, nc := row+dr, col+dc
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
						if content[nr][nc] == '@' {
							NumberOfAT++
						}
					}
				}
				if NumberOfAT < 4 {
					SumRolls++
				}
			}
		}
	}
	return SumRolls
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
