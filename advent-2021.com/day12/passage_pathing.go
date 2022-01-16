package day12

import (
	"bufio"
	"os"
	"strings"
)

func PassagePathing(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	connections := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "-")
		if strings.Contains(line, "start") {
			child := lineSplit[0]
			if child == "start" {
				child = lineSplit[1]
			}
			connections["start"] = append(connections["start"], child)
		} else if strings.Contains(line, "end") {
			parent := lineSplit[0]
			if parent == "end" {
				parent = lineSplit[1]
			}
			connections[parent] = append(connections[parent], "end")
		} else {
			connections[lineSplit[0]] = append(connections[lineSplit[0]], lineSplit[1])
			connections[lineSplit[1]] = append(connections[lineSplit[1]], lineSplit[0])
		}
	}

	paths := make(map[string]int)
	findPaths("start", connections, paths, "")
	//fmt.Println(len(paths))
	return len(paths)
}

func findPaths(nextConnection string, connections map[string][]string, paths map[string]int, path string) string {
	if nextConnection != "end" {
		if children, ok := connections[nextConnection]; ok {
			if len(children) == 1 && strings.HasSuffix(path, "-"+children[0]) {
				if !isSmallCave(children[0]) ||
					(isSmallCave(children[0]) && !strings.Contains(path, "-"+children[0])) {
					findPaths(children[0], connections, paths, path+"-"+nextConnection)
				}
			} else {
				for _, child := range children {
					if !isSmallCave(child) ||
						(isSmallCave(child) && !strings.Contains(path, "-"+child)) {
						findPaths(child, connections, paths, path+"-"+nextConnection)
					}
				}
			}
		}
	} else {
		paths[path[1:]+"-"+nextConnection] = 7 // only key matters
	}
	return path + "-" + nextConnection
}

func PassagePathing_Part2(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	connections := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "-")
		if strings.Contains(line, "start") {
			child := lineSplit[0]
			if child == "start" {
				child = lineSplit[1]
			}
			connections["start"] = append(connections["start"], child)
		} else if strings.Contains(line, "end") {
			parent := lineSplit[0]
			if parent == "end" {
				parent = lineSplit[1]
			}
			connections[parent] = append(connections[parent], "end")
		} else {
			connections[lineSplit[0]] = append(connections[lineSplit[0]], lineSplit[1])
			connections[lineSplit[1]] = append(connections[lineSplit[1]], lineSplit[0])
		}
	}

	paths := make(map[string]int)
	findPaths_Part2("start", connections, paths, "")
	//fmt.Println(len(paths))
	return len(paths)
}

func findPaths_Part2(nextConnection string, connections map[string][]string, paths map[string]int, path string) string {
	if nextConnection != "end" {
		if children, ok := connections[nextConnection]; ok {
			for _, child := range children {
				if !isSmallCave(child) || canEnterCave(path+","+nextConnection, child) {
					findPaths_Part2(child, connections, paths, path+","+nextConnection)
				}
			}
		}
	} else {
		paths[path[1:]+","+nextConnection] = 7 // only key matters
	}
	return path + "," + nextConnection
}

func isSmallCave(connection string) bool {
	if connection != "start" && connection != "end" && strings.ToLower(connection) == connection {
		return true
	}
	return false
}

func canEnterCave(path string, id string) bool {
	if isSmallCave(id) {
		caves := make(map[string]int)
		hasTwo := false
		for _, item := range strings.Split(path, ",") {
			if isSmallCave(item) {
				caves[item]++
				if caves[item] == 2 {
					hasTwo = true
				}
			}
		}
		if caves[id] == 2 ||
			(hasTwo && caves[id] == 1) {
			return false
		}
	}
	return true
}
