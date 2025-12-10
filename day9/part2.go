package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var location = "12_9.txt"

type Point struct {
	x, y int
}

func main() {
	input := ReadContent(location)
	coords := SpiltElements(input)
	part2 := Result(coords)
	fmt.Printf("Part 2 Answer: %d\n", part2)
}

func Result(points []Point) int {
	var maxArea int

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			x1 := MIN(p1.x, p2.x)
			x2 := MAX(p1.x, p2.x)
			y1 := MIN(p1.y, p2.y)
			y2 := MAX(p1.y, p2.y)

			width := x2 - x1 + 1
			height := y2 - y1 + 1
			area := width * height
			if area <= maxArea {
				continue
			}

			if isValidRect(x1, x2, y1, y2, points) {
				maxArea = area
			}
		}
	}

	return maxArea
}

func isValidRect(x1, x2, y1, y2 int, poly []Point) bool {

	//中心点坐标
	CenterPointX := x1 + x2
	CenterPointY := y1 + y2

	//检查中心点是否在多边形里面
	if !isPointInPoly(CenterPointX, CenterPointY, poly) {
		return false
	}

	//确保多边形的边界不穿过矩形的内部
	PointNumber := len(poly)
	for i := 0; i < PointNumber; i++ {
		u := poly[i]
		v := poly[(i+1)%PointNumber]

		if u.x == v.x {
			edgeX := u.x
			edgeYMIN := MIN(u.y, v.y)
			edgeYMAX := MAX(u.y, v.y)

			if edgeX > x1 && edgeX < x2 {
				overlapStart := MAX(y1, edgeYMIN)
				overlapEnd := MIN(y2, edgeYMAX)
				if overlapStart < overlapEnd {
					return false
				}
			}
		} else {
			edgeY := u.y
			edgeXMIN := MIN(u.x, v.x)
			edgeXMAX := MAX(u.x, v.x)

			if edgeY > y1 && edgeY < y2 {
				overlapStart := MAX(x1, edgeXMIN)
				overlapEnd := MIN(x2, edgeXMAX)
				if overlapStart < overlapEnd {
					return false
				}
			}
		}
	}

	return true
}

// 检查点是否在多边形里面/边上
func isPointInPoly(x, y int, poly []Point) bool { //point---所有的位于多边形上的点

	//1检查的是中心点的位置，是否在多边形的内部  是否恰好在边界上
	n := len(poly) //多边形点的个数  --边的数量

	for i := 0; i < n; i++ {
		u := poly[i]
		v := poly[(i+1)%n] //相邻的下一个点。由于列表是循环的，所以他们组成一条边

		//二倍处理边界
		ux_2, uy_2, vx_2, vy_2 := Double(u, v)

		if u.x == v.x { //如果的x相同，那么在同一列之上，找到行的边界
			if ux_2 == x {
				minY := MIN(uy_2, vy_2)
				maxY := MAX(uy_2, vy_2)
				if y >= minY && y <= maxY {
					return true //检查这个点是否在这个范围之内
				}
			}
		} else {
			if uy_2 == y { //同理
				minX := MIN(ux_2, vx_2)
				maxX := MAX(ux_2, vx_2)
				if x >= minX && x <= maxX {
					return true
				}
			}
		}
	}

	//发送射线，检查交点个数--如果不在边界上
	intersections := 0
	for i := 0; i < n; i++ {
		//遍历了每一条边

		u := poly[i]
		v := poly[(i+1)%n] //获取当前的边

		if u.x == v.x { //仅处理垂直边

			minY := MIN(u.y, v.y) * 2 //边的Y范围下限
			maxY := MAX(u.y, v.y) * 2 //上限
			ex := u.x * 2             //X的坐标 。这条边的x的二倍值
			//判断是否相交
			if y >= minY && y < maxY && ex > x { //x中心点的二倍坐标，如果ex>x那么他们不会遇到，y是要检查范围
				intersections++
			}
		}
	}

	return intersections%2 == 1
}
func MAX(u, v int) int {
	if u > v {
		return u
	}
	return v
}
func MIN(u, v int) int {
	if u < v {
		return u
	}
	return v
}
func Double(u, v Point) (int, int, int, int) {
	u0_2 := u.x * 2
	u1_2 := u.y * 2
	v0_2 := v.x * 2
	v1_2 := v.y * 2
	return u0_2, u1_2, v0_2, v1_2
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
