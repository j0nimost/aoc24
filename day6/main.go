package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pathFinderPartOne(areaMap [][]string) (map[string]int, int) {

	var steps map[string]int = make(map[string]int)

	// find initial position
	isFound := false
	posy, posx := 0, 0
	face := ""
	for i := 0; i < len(areaMap); i++ {
		for j := 0; j < len(areaMap[i]); j++ {
			p := areaMap[i][j]
			if p == "^" || p == "<" || p == ">" || p == "v" {
				isFound = true
				posy = i
				posx = j
				face = p
				break
			}
		}

		if isFound {
			break
		}
	}

	// loop until out of the map and count the steps taken
	for {
		steps[fmt.Sprintf("%d,%d", posy, posx)] += 1

		if posx == 0 || posx == len(areaMap)-1 {
			break
		}
		if posy == 0 || posy == len(areaMap[posx])-1 {
			break
		}
		if face == "^" && posy-1 >= 0 && areaMap[posy-1][posx] != "#" {
			posy--
		} else if face == "^" && posy-1 >= 0 && areaMap[posy-1][posx] == "#" {
			face = ">"
		}

		if face == ">" && posx+1 < len(areaMap[posy]) && areaMap[posy][posx+1] != "#" {
			posx++
		} else if face == ">" && posx+1 < len(areaMap[posy]) && areaMap[posy][posx+1] == "#" {
			face = "v"
		}

		if face == "v" && posy+1 < len(areaMap) && areaMap[posy+1][posx] != "#" {
			posy++
		} else if face == "v" && posy+1 < len(areaMap) && areaMap[posy+1][posx] == "#" {
			face = "<"
		}

		if face == "<" && posx-1 >= 0 && areaMap[posy][posx-1] != "#" {
			posx--
		} else if face == "<" && posx-1 >= 0 && areaMap[posy][posx-1] == "#" {
			face = "^"
		}
	}
	return steps, len(steps)
}

func pathFinderPartTwo(areaMap [][]string, steps map[string]int) int {
	face := ""
	isFound := false
	initposy, initposx := 0, 0
	count := 0

	for i := 0; i < len(areaMap); i++ {
		for j := 0; j < len(areaMap[i]); j++ {
			p := areaMap[i][j]
			if p == "^" || p == "<" || p == ">" || p == "v" {
				isFound = true
				initposx = j
				initposy = i
				break
			}
		}

		if isFound {
			break
		}
	}

	for k := range steps {
		face = areaMap[initposy][initposx]
		// check starting pos add obstacle
		if k == face {
			continue
		}
		//create a copy of areaMap
		areaMapCopy := make([][]string, len(areaMap))

		for m := 0; m < len(areaMap); m++ {
			areaMapCopy[m] = make([]string, len(areaMap[m]))
			copy(areaMapCopy[m], areaMap[m])
		}

		visitedPath := make(map[string]int)

		pos := strings.Split(k, ",")
		i, _ := strconv.Atoi(pos[0])
		j, _ := strconv.Atoi(pos[1])

		areaMapCopy[i][j] = "#"
		posx := initposx
		posy := initposy
		for {
			visitedPath[fmt.Sprintf("%d,%d", posy, posx)] += 1
			if visitedPath[fmt.Sprintf("%d,%d", posy, posx)] >= 30 {
				count++
				break
			}

			if posx == 0 || posx == len(areaMapCopy)-1 {
				break
			}
			if posy == 0 || posy == len(areaMapCopy[posx])-1 {
				break
			}

			if face == "^" && posy-1 >= 0 && areaMapCopy[posy-1][posx] != "#" {
				posy--
			} else if face == "^" && posy-1 >= 0 && areaMapCopy[posy-1][posx] == "#" {
				face = ">"
			}

			if face == ">" && posx+1 < len(areaMapCopy[posy]) && areaMapCopy[posy][posx+1] != "#" {
				posx++
			} else if face == ">" && posx+1 < len(areaMapCopy[posy]) && areaMapCopy[posy][posx+1] == "#" {
				face = "v"
			}

			if face == "v" && posy+1 < len(areaMapCopy) && areaMapCopy[posy+1][posx] != "#" {
				posy++
			} else if face == "v" && posy+1 < len(areaMapCopy) && areaMapCopy[posy+1][posx] == "#" {
				face = "<"
			}

			if face == "<" && posx-1 >= 0 && areaMapCopy[posy][posx-1] != "#" {
				posx--
			} else if face == "<" && posx-1 >= 0 && areaMapCopy[posy][posx-1] == "#" {
				face = "^"
			}
		}
	}

	return count
}

func main() {
	// file, err := os.Open("/home/joni/shugli/aoc24/day6/test.txt")
	// file, err := os.Open("/home/joni/shugli/aoc24/day6/test2.txt")
	file, err := os.Open("/home/joni/shugli/aoc24/day6/input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var areaMap [][]string = [][]string{}
	for scanner.Scan() {
		txt := scanner.Text()
		row := []string{}
		for _, v := range txt {
			row = append(row, string(v))
		}
		areaMap = append(areaMap, row)
	}

	steps, r0 := pathFinderPartOne(areaMap)
	r1 := pathFinderPartTwo(areaMap, steps)
	fmt.Printf("answer : %d\n", r0)
	fmt.Printf("answer : %d\n", r1)
}
