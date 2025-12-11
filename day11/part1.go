package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pair struct {
	Head string
	Tail []string
}

type Graph map[string][]string

var location = "12_11.txt"

func main() {
	contents := ReadContent(location)
	pairs := GeneratePairs(contents)

	graph := builtGraph(pairs)
	visited := make(map[string]bool)
	result := FindOutPath(graph, "you", visited)
	fmt.Println(result)
}

func GeneratePairs(contents []string) []Pair {
	var pairs []Pair
	for _, content := range contents {
		var pair Pair

		values := strings.SplitN(content, ":", 2)

		if len(values) != 2 {
			continue
		}

		pair.Head = strings.TrimSpace(values[0])
		pair.Tail = strings.Fields(values[1])

		pairs = append(pairs, pair)
	}
	return pairs
}

func builtGraph(pairs []Pair) Graph {
	graph := make(Graph)
	for _, p := range pairs {
		graph[p.Head] = p.Tail
	}
	return graph
}

func FindOutPath(graph Graph, currentDevice string, visited map[string]bool) int {
	if visited[currentDevice] {
		return 0
	}

	if currentDevice == "out" {
		return 1
	}

	nextDevices, exists := graph[currentDevice]
	if !exists || len(nextDevices) == 0 {
		return 0
	}

	visited[currentDevice] = true

	totalPaths := 0
	for _, nextDevice := range nextDevices {
		totalPaths += FindOutPath(graph, nextDevice, visited)
	}

	delete(visited, currentDevice)

	return totalPaths
}

func ReadContent(location string) []string {
	var c []string
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content := strings.TrimSpace(scanner.Text())
		if content != "" {
			c = append(c, content)
		}
	}
	return c
}
