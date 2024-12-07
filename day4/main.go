package main

import (
	"bufio"
	"fmt"
	"os"
)

var xmasWord []string = []string{"X", "M", "A", "S"}

func traversePartOne(a int, b int, xIndex int, puzzle [][]string, isLeft bool, isRight bool, isUp bool, isDown bool) int {

	for xIndex < len(xmasWord) && xmasWord[xIndex] == puzzle[a][b] {
		// go crazy
		xIndex++

		if isDown && isRight && a+1 < len(puzzle) && b+1 < len(puzzle[a]) {
			a++
			b++
		} else if isDown && isLeft && a+1 < len(puzzle) && b-1 >= 0 {
			a++
			b--
		} else if isUp && isRight && a-1 >= 0 && b+1 < len(puzzle[a]) {
			a--
			b++
		} else if isUp && isLeft && a-1 >= 0 && b-1 >= 0 {
			a--
			b--
		} else if isUp && !(isLeft || isRight || isDown) && a-1 >= 0 {
			a--
		} else if isDown && !(isLeft || isRight || isUp) && a+1 < len(puzzle) {
			a++
		} else if isRight && !(isLeft || isUp || isDown) && b+1 < len(puzzle[a]) {
			b++
		} else if isLeft && !(isUp || isRight || isDown) && b-1 >= 0 {
			b--
		}

	}
	if xIndex == len(xmasWord) {
		return 1
	}

	return 0
}

func traversePartTwo(a int, b int, puzzle [][]string) bool {
	// go crazy
	checker := []string{}

	// left and up
	if a-1 >= 0 && b-1 >= 0 && (puzzle[a-1][b-1] == "M" || puzzle[a-1][b-1] == "S") {
		checker = append(checker, puzzle[a-1][b-1])
	}

	// right and down
	if a+1 < len(puzzle) && b+1 < len(puzzle[a]) && (puzzle[a+1][b+1] == "M" || puzzle[a+1][b+1] == "S") {
		checker = append(checker, puzzle[a+1][b+1])
	}

	// left and down
	if a+1 < len(puzzle) && b-1 >= 0 && (puzzle[a+1][b-1] == "M" || puzzle[a+1][b-1] == "S") {
		checker = append(checker, puzzle[a+1][b-1])
	}

	// right and up
	if a-1 >= 0 && b+1 < len(puzzle[a]) && (puzzle[a-1][b+1] == "M" || puzzle[a-1][b+1] == "S") {
		checker = append(checker, puzzle[a-1][b+1])
	}

	if len(checker) != 4 {
		return false
	}
	m, s := 2, 2
	for _, v := range checker {
		if v == "M" {
			m--
		} else if v == "S" {
			s--
		}
	}

	// pair 'em up
	return m == 0 && s == 0 && checker[0] != checker[1]
}

func xmasWordSearchPartOne(puzzle [][]string) int {

	count := 0
	// traverse the whole array
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if xmasWord[0] == puzzle[i][j] {
				count += traversePartOne(i, j, 0, puzzle, true, false, false, false) // left
				count += traversePartOne(i, j, 0, puzzle, false, true, false, false) // right
				count += traversePartOne(i, j, 0, puzzle, false, false, true, false) // up
				count += traversePartOne(i, j, 0, puzzle, false, false, false, true) // down
				count += traversePartOne(i, j, 0, puzzle, true, false, true, false)  // left and up
				count += traversePartOne(i, j, 0, puzzle, true, false, false, true)  // left and down
				count += traversePartOne(i, j, 0, puzzle, false, true, true, false)  // right and up
				count += traversePartOne(i, j, 0, puzzle, false, true, false, true)  // right and down
			}
		}
	}
	return count
}

func xmasWordSearchPartTwo(puzzle [][]string) int {
	count := 0
	// traverse the whole array
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == "A" && traversePartTwo(i, j, puzzle) {
				count++
			}
		}
	}
	return count
}

func main() {
	// file, err := os.Open("/home/joni/shugli/aoc24/day4/test.txt")
	file, err := os.Open("/home/joni/shugli/aoc24/day4/input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// sum := 0
	var puzzle [][]string = [][]string{}
	for scanner.Scan() {
		txt := scanner.Text()
		row := []string{}
		for _, v := range txt {
			row = append(row, string(v))
		}
		puzzle = append(puzzle, row)
	}

	r0 := xmasWordSearchPartOne(puzzle)
	r1 := xmasWordSearchPartTwo(puzzle)
	fmt.Printf("answer : %d\n", r0)
	fmt.Printf("answer : %d\n", r1)
}
