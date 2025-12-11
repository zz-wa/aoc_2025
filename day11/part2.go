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
type targetDevice struct {
	s        string
	fft, dac bool
}

type Graph map[string][]string

var location = "/home/butter-july/桌面/12_11.txt"

func main() {
	contents := ReadContent(location)
	pairs := GeneratePairs(contents)
	graph := builtGraph(pairs)
	cache := make(map[targetDevice]int)
	if nextNoodes, exists := graph["svr"]; exists {
		result := FindOutPathWithFlags(graph, nextNoodes, false, false, cache)
		fmt.Println(result)
	}
}
func FindOutPathWithFlags(graph Graph, nextNodes []string, fft, dac bool, cache map[targetDevice]int) (r int) {
	for _, currentDevice := range nextNodes {
		if currentDevice == "out" {
			if fft && dac {
				r += 1
			}
			continue
		}
		newfft, newdac := fft, dac
		switch currentDevice {
		case "fft":
			newfft = true
		case "dac":
			newdac = true
		default:
		}
		key := targetDevice{currentDevice, newfft, newdac}
		if v, ok := cache[key]; ok {
			r += v
			continue
		}
		if nextDevices, exist := graph[currentDevice]; exist {
			result := FindOutPathWithFlags(graph, nextDevices, newfft, newdac, cache)
			cache[key] = result
			r += result
		}
	}
	return r
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
