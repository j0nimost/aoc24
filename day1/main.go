package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/1

func merge(arr []int, l int, m int, r int) {
	lArr := make([]int, m-l+1)
	rArr := make([]int, r-m)

	i, j := 0, 0
	for i = 0; i < len(lArr); i++ {
		lArr[i] = arr[l+i]
	}

	for j = 0; j < len(rArr); j++ {
		rArr[j] = arr[m+j+1]
	}

	i, j = 0, 0
	k := l
	for i < len(lArr) && j < len(rArr) {
		if lArr[i] <= rArr[j] {
			arr[k] = lArr[i]
			i++
		} else {
			arr[k] = rArr[j]
			j++
		}
		k++
	}

	// copy remaining
	for i < len(lArr) {
		arr[k] = lArr[i]
		i++
		k++
	}

	for j < len(rArr) {
		arr[k] = rArr[j]
		j++
		k++
	}

}

func mergeSort(arr []int, l int, r int) {
	if l < r {
		m := l + (r-l)/2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func processArraysPartOne(arr1 []int, arr2 []int) int {
	// sort
	mergeSort(arr1, 0, len(arr1)-1)
	mergeSort(arr2, 0, len(arr2)-1)

	sum := 0
	for i := 0; i < len(arr1); i++ {
		sum += int(math.Abs(float64(arr1[i]) - float64(arr2[i])))
	}
	return sum
}

func processArraysPartTwo(arr1 []int, arr2 []int) int {
	// sort
	mergeSort(arr1, 0, len(arr1)-1)
	mergeSort(arr2, 0, len(arr2)-1)

	sum := 0

	hashArr := make(map[int]int)
	for i := 0; i < len(arr2); i++ {
		if _, ok := hashArr[arr2[i]]; ok {
			hashArr[arr2[i]] += 1
		} else {
			hashArr[arr2[i]] = 1
		}
	}

	for j := 0; j < len(arr1); j++ {
		if v, ok := hashArr[arr1[j]]; ok {
			sum += arr1[j] * v
		}
	}

	return sum
}

func main() {
	file, err := os.Open("/home/joni/shugli/aoc24/day1/input.txt")
	// file, err := os.Open("/home/joni/shugli/aoc24/day1/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	arr1 := []int{}
	arr2 := []int{}
	for scanner.Scan() {
		x := scanner.Text()
		columns := strings.Split(x, "   ")
		n1, err := strconv.Atoi(columns[0])
		if err != nil {
			panic(err)
		}

		n2, err := strconv.Atoi(columns[1])
		if err != nil {
			panic(err)
		}
		arr1 = append(arr1, n1)
		arr2 = append(arr2, n2)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	r := processArraysPartOne(arr1, arr2)
	fmt.Printf("answer %d\n", r)
	r = processArraysPartTwo(arr1, arr2)
	fmt.Printf("answer %d\n", r)
}
