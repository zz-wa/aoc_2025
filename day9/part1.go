package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	col int
	row int
}

var location = "/home/butter-july/桌面/12_9.txt"

func main() {
	contents := ReadContent(location)
	Points := SpiltElements(contents)
	//进行计算
	area := MaxArea(Points)
	fmt.Println(area)
}
func MaxArea(points []Point) int {
	var maxArea int
	for i := 0; i < len(points); i++ {
		for j := i; j < len(points); j++ {
			currentArea := Calculate(i, j, points)
			if currentArea > maxArea {
				maxArea = currentArea
			}
		}
	}
	return maxArea
}
func Calculate(i, j int, points []Point) int {

	length := math.Abs(float64(points[i].col-points[j].col))+1
	width := math.Abs(float64(points[i].row-points[j].row) )+1
	return int(length * width)
}
func SpiltElements(contents []string) []Point {
	var Points []Point
	for _, value := range contents {
		contentLine := strings.Split(value, ",")
		x, _ := strconv.Atoi(contentLine[0])
		y, _ := strconv.Atoi(contentLine[1])
		newPoint := Point{x, y}
		Points = append(Points, newPoint)
	}
	return Points
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
