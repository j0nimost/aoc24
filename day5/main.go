package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func pageOrderingRulePartOne(rule map[int][]int, update []string) int {

	// loop string and convert to ints
	sum := 0
	for i := 0; i < len(update); i++ {
		// init array
		updateNums := []int{}
		un := strings.Split(update[i], ",")
		for _, v := range un {
			n, _ := strconv.Atoi(v)
			updateNums = append(updateNums, n)
		}

		// validate each number
		notValid := false

		for j := 0; j < len(updateNums); j++ {
			orderArr := rule[updateNums[j]]
			for k := j - 1; k >= 0; k-- {
				// check all prior numbers
				for _, v := range orderArr {
					if v == updateNums[k] {
						// break
						// not valid
						notValid = true
						break
					}
				}

				if notValid {
					break
				}
			}
		}

		if notValid {
			continue
		} else if !notValid {
			m := len(updateNums) / 2
			sum += updateNums[m]
		}
	}
	return sum
}

func pageOrderingRulePartTwo(rule map[int][]int, update []string) int {
	// loop string and convert to ints
	sum := 0
	// fmt.Println(update[1])
	for i := 0; i < len(update); i++ {
		// init array
		updateNums := []int{}
		un := strings.Split(update[i], ",")
		for _, v := range un {
			n, _ := strconv.Atoi(v)
			updateNums = append(updateNums, n)
		}

		// validate each number
		notValid := false

		for j := 0; j < len(updateNums); j++ {
			orderArr := rule[updateNums[j]]
			for k := j - 1; k >= 0; k-- {
				// check all prior numbers
				for _, v := range orderArr {
					if v == updateNums[k] {
						// break
						// not valid
						notValid = true
						break
					}
				}

				if notValid {
					break
				}
			}
		}

		if notValid {
			// pathetic sort using paging rule
			for m := 0; m < len(updateNums); m++ {
				orderArr := rule[updateNums[m]]
				for n := m; n >= 0; n-- {
					for _, v := range orderArr {
						if updateNums[n] == v {
							temp := updateNums[m]
							updateNums[m] = updateNums[n]
							updateNums[n] = temp
							m = n
						}
					}
				}
			}
			m := len(updateNums) / 2
			sum += updateNums[m]
		} else if !notValid {
			continue
		}
	}
	return sum
}

func main() {
	// file, err := os.Open("/home/joni/shugli/aoc24/day5/test.txt")
	file, err := os.Open("/home/joni/shugli/aoc24/day5/input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// sum := 0
	var order map[int][]int = make(map[int][]int)
	var update []string = []string{}
	isOrderInput := true
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			isOrderInput = false
			continue
		}

		if isOrderInput {
			o := strings.Split(txt, "|")
			n1, _ := strconv.Atoi(o[0])
			n2, _ := strconv.Atoi(o[1])

			arr := order[n1]
			arr = append(arr, n2)
			order[n1] = arr
		} else {
			update = append(update, txt)
		}
	}

	// fmt.Printf("%d %d\n", len(order), len(update))
	r0 := pageOrderingRulePartOne(order, update)
	r1 := pageOrderingRulePartTwo(order, update)
	fmt.Printf("answer : %d\n", r0)
	fmt.Printf("answer : %d\n", r1)
}
