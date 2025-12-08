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

const MAX = 188_000_000

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
	edges := GenerateEdges(Points)
	n := len(Points)

	result := Result(n, edges)
	fmt.Println(result)
}

func Result(n int, edges []Edge) int {
	dsu := NewDSU(n)
	limit := min(1000, len(edges))
	for i := 0; i < limit; i++ {
		dsu.Union(edges[i].a, edges[i].b)
	}
	circuitMap := make(map[int]int)
	for i := 0; i < n; i++ {
		root := dsu.FInd(i)
		circuitMap[root]++
	}
	var circuits []int
	for _, size := range circuitMap {
		circuits = append(circuits, size)
	}
	sort.Slice(circuits, func(i, j int) bool {
		return circuits[i] > circuits[j]
	})
	topCount := min(3, len(circuits))
	product := 1
	for i := 0; i < topCount; i++ {
		product *= circuits[i]
	}

	return product
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = i
	}
	return &DSU{parent, size}
}

func (dsu *DSU) FInd(x int) int {
	if dsu.parent[x] != x {
		dsu.parent[x] = dsu.FInd(dsu.parent[x])
	}
	return dsu.parent[x]
}
func (dsu *DSU) Union(x, y int) {
	rootX := dsu.FInd(x)
	rootY := dsu.FInd(y)
	if rootX == rootY {
		return
	}
	if dsu.size[rootX] < dsu.size[rootY] {
		rootX, rootY = rootY, rootX
	}
	dsu.parent[rootY] = rootX
	dsu.size[rootX] += dsu.size[rootY]
}
func GenerateEdges(points []Point) []Edge {
	var edges []Edge
	n := len(points)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			d := distance(points[i], points[j])
			if d < MAX {
				edges = append(edges, Edge{i, j, d})
			}
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
