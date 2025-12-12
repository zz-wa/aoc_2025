package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var location = "/home/butter-july/桌面/12_12.txt"

type Region struct {
	size          int
	presentNumber int
}

func main() {
	contents := ReadContent(location)
	/*	var Num [][]int
	 */
	Allregion := GetRegions(contents)

	result := GetResult(Allregion)
	fmt.Println(result)
}
func GetResult(allRegions []Region) int {
	result := 0
	for _, value := range allRegions {
		if value.size < value.presentNumber*9 {
			result += 0
		} else {
			result += 1
		}
	}
	return result
}
func GetRegions(contents []string) []Region {
	Allines := contents[30:]
	var regions []Region
	for _, value := range Allines {
		region := Region{
			size:          0,
			presentNumber: 0,
		}
		n := strings.Replace(value, ":", " ", 1)
		line := strings.Fields(n)

		for j, number := range line {
			if j == 0 {
				region.size = CalculateSize(number)
			}
			currentPresentNumber, _ := strconv.Atoi(number)
			region.presentNumber += currentPresentNumber
		}
		regions = append(regions, region)

	}
	return regions
}
func CalculateSize(number string) int {
	n := strings.Split(number, "x")
	length, _ := strconv.Atoi(n[0])
	width, _ := strconv.Atoi(n[1])
	return width * length
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
