package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func main() {
	// call solution with user inputted array
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	slice := numbers(scanner.Text())
	fmt.Println(solution(slice))
}

func solution(v []int) (maxSoFar int) {
	maxSoFar = 1
	maxEndingHere := 1
	sort.Ints(v)
	lastElement := v[0]
	for i := 1; i < len(v); i++ {
		if v[i]-lastElement == 1 {
			maxEndingHere++
			if maxSoFar < maxEndingHere {
				maxSoFar = maxEndingHere
			}
		} else if v[i] != lastElement {
			maxEndingHere = 1
		}
		lastElement = v[i]
	}
	return
}
