package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)



var location = "12_8.txt"

type Edge struct {
	a, b     int
	distance int
}
type Point struct {
	x int
	y int
	z int
}
type DSU struct {
	parent []int
	size   []int
}

func main() {
	contents := ReadContent(location)
	Points := SpiltElements(contents)
	fmt.Println(Points)
	edges := GenerateEdges(Points)
	n := len(Points)

	result := Result(n, edges, Points)
	fmt.Println(result)
}

func Result(n int, edges []Edge, point []Point) int {
	dsu := NewDSU(n) //parent size

	components := n
	var last Edge

	for _, edge := range edges {
		if dsu.Union(edge.a, edge.b) {
			last = edge
			components--
			if components == 1 {
				break
			}
		}
	}

	return point[last.a].x * point[last.b].x
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent, size}
}
func (dsu *DSU) Find(x int) int {
	if dsu.parent[x] != x { //Locate
		dsu.parent[x] = dsu.Find(dsu.parent[x])
		x = dsu.parent[x]
	}
	return dsu.parent[x]
}
func (dsu *DSU) Union(x, y int) bool {
	rootX := dsu.Find(x)
	rootY := dsu.Find(y)
	if rootX == rootY {
		return false
	}
	if dsu.size[rootX] < dsu.size[rootY] {
		rootX, rootY = rootY, rootX
	}
	dsu.parent[rootY] = rootX
	dsu.size[rootX] += dsu.size[rootY]
	return true
}


func GenerateEdges(points []Point) []Edge {
	var edges []Edge
	n := len(points)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, Edge{i, j, distance(points[i], points[j])})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	return edges
}

func distance(a, b Point) int {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return dx*dx + dy*dy + dz*dz
}

func SpiltElements(contents []string) []Point {
	var Points []Point
	for _, value := range contents {
		contentLine := strings.Split(value, ",")
		x, _ := strconv.Atoi(contentLine[0])
		y, _ := strconv.Atoi(contentLine[1])
		z, _ := strconv.Atoi(contentLine[2])
		newPoint := Point{x, y, z}
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
