package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/2

func safeScanPartOne(arr []int) int {
	safe := 0
	dValid := true
	checker := []int{}

	for i := 0; i < len(arr); i++ {
		if i+1 == len(arr) {
			break
		}

		if arr[i] > arr[i+1] {
			checker = append(checker, 1)
		}
		if arr[i] < arr[i+1] {
			checker = append(checker, 0)
		}

		diff := int(math.Abs(float64(arr[i]) - float64(arr[i+1])))

		if diff < 1 || diff > 3 {
			dValid = false
			break
		}
	}

	if len(checker) == 0 {
		return 0
	}
	base := checker[0]
	for _, v := range checker {
		if base != v {
			dValid = false
			break
		}
	}

	if dValid {
		safe++
	}
	return safe
}

func safeScanPartTwo(arr []int) int {
	dValid := true
	r := safeScanPartOne(arr)
	if r == 1 {
		return 1
	}

	// try replacing
	for i := 0; i < len(arr); i++ {
		n := 0
		dValid = true
		checker := []int{}
		for j := 0; j < len(arr); j++ {
			n = j + 1
			if i == j {
				continue
			}

			if n == len(arr) || (n == i && n+1 == len(arr)) {
				break
			} else if n == i && n+1 < len(arr) {
				n++
			}

			if arr[j] > arr[n] {
				checker = append(checker, 1)
			} else if arr[j] < arr[n] {
				checker = append(checker, 0)
			}

			diff := int(math.Abs(float64(arr[j]) - float64(arr[n])))

			if diff < 1 || diff > 3 {
				dValid = false
				break
			}
		}
		base := -1

		if len(checker) == 0 {
			dValid = false
		} else {
			base = checker[0]
		}
		for _, v := range checker {
			if base != v {
				dValid = false
				break
			}
		}

		if dValid {
			return 1
		}
	}
	return 0
}

func main() {
	file, err := os.Open("/home/joni/shugli/aoc24/day2/input.txt")
	// file, err := os.Open("/home/joni/shugli/aoc24/day2/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	s0 := 0
	s1 := 0
	for scanner.Scan() {
		x := scanner.Text()

		levels := strings.Split(x, " ")
		rpt := []int{}

		for _, v := range levels {
			x, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			rpt = append(rpt, x)
		}
		s0 += safeScanPartOne(rpt)
		s1 += safeScanPartTwo(rpt)
	}

	fmt.Printf("answer: %d\n", s0)
	fmt.Printf("answer: %d\n", s1)
}
