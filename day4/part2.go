package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var location = "12_4.txt"

type State [][]byte
type RollPosition struct {
	Row int
	Col int
}

func main() {
	contents := ReadContent(location) //[]string
	state := ToByte(contents)         //[][]byte

	_, rolls := GetRolls(state)
	fmt.Println(rolls)
}

func FindAccessibleRolls(content State) (State, int) {
	rows := len(content)
	cols := len(content[0])
	nextState, count := CheckNeighbors(rows, cols, content)
	return nextState, count
}

func GetRolls(state State) (State, int) {
	sum := 0
	nextstate := state
	for {
		var count int
		nextstate, count = FindAccessibleRolls(nextstate)
		if count == 0 {
			return nextstate, sum
		}
		sum += count
	}
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
func CheckNeighbors(rows, cols int, content State) (nextState State, count int) {
	nextState = Clone(content)

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
					nextState[row][col] = '.'
					count++
				}
			}
		}
	}
	return nextState, count
}
func Clone(state State) [][]byte {
	next := make(State, len(state))
	for i := range state {
		next[i] = make([]byte, len(state[i]))
		copy(next[i], state[i])
	}
	return next
}
func ToByte(contents []string) State {
	var state [][]byte
	//变成[][]byte类型
	for i, line := range contents {
		state = append(state, []byte{})
		for _, v := range line {
			state[i] = append(state[i], byte(v))
		}
	}

	return state
}
