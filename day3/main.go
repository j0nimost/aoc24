package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//https://adventofcode.com/2024/day/3

func multiplyPartOne(s string) uint64 {
	r, err := regexp.Compile(`(mul)([\(][0-9]{1,3}[,\^][0-9]{1,3}[\)])`)
	n, _ := regexp.Compile(`[0-9]{1,3}`)

	if err != nil {
		panic(err)
	}

	matches := r.FindAllString(s, -1)
	var sum uint64 = 0
	for _, m := range matches {
		numbers := n.FindAllString(m, -1)

		if len(numbers) != 2 {
			break
		}
		// convert
		n0, _ := strconv.Atoi(numbers[0])
		n1, _ := strconv.Atoi(numbers[1])

		sum += uint64(n0 * n1)

	}
	return sum
}

func multiplyPartTwo(s string) uint64 {
	do, _ := regexp.Compile(`do\(\)`)
	dont, _ := regexp.Compile(`don't\(\)`)

	doMatches := do.FindAllStringSubmatchIndex(s, -1)
	dontMatches := dont.FindAllStringSubmatchIndex(s, -1)

	var sum uint64 = 0

	i, j := 0, 0
	// isMult := true
	for {
		if j >= len(dontMatches) && i < len(doMatches) {
			x := s[doMatches[i][1] : len(s)-1]
			fmt.Printf("%s\n", x)
			sum += multiplyPartOne(x)
			break
		}

		if i >= len(doMatches) && j < len(dontMatches) {
			break
		}

		l := j
		if j == len(dontMatches) {
			l--
		}
		for i < len(doMatches) && doMatches[i][0] < dontMatches[l][0] {
			// isMult = true
			i++
		}

		m := i
		if i == len(doMatches) {
			m--
		}

		for j < len(dontMatches) && dontMatches[j][0] < doMatches[m][0] {
			j++
			// isMult = false
		}

		if doMatches[m][0] < dontMatches[l][0] {
			x := s[doMatches[m][1]:dontMatches[l][0]]
			fmt.Printf("%s\n\n\n\n\n", x)

			sum += multiplyPartOne(x)
		}

		// fmt.Printf("%d %d\n", i, j)
	}

	if dontMatches[0][0] < doMatches[0][0] {
		fmt.Printf("%s\n\n\n\n\n", s[0:dontMatches[0][0]])
		sum += multiplyPartOne(s[0:dontMatches[0][0]])
	}

	fmt.Printf("%d %d\n", len(doMatches), len(dontMatches))

	return sum
}

func main() {
	// file, err := os.Open("/home/joni/shugli/aoc24/day3/test.txt")
	file, err := os.Open("/home/joni/shugli/aoc24/day3/input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// sum := 0
	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}

	sum0 := multiplyPartOne(sb.String())
	// failed pt2 regex is hard
	sum1 := multiplyPartTwo(sb.String())
	fmt.Printf("answer : %d\n", sum0)
	fmt.Printf("answer : %d\n", sum1)

}
