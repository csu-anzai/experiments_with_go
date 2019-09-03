package hashes

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
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
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//slice := numbers(scanner.Text())
	slice := []int{100, 4, 200, 1, 3, 2, 101, 102, 5, 78, 42, 12, 213, 32, 43, 44, 45}
	start := time.Now()
	fmt.Println(sortSolution(slice))
	elapsed := time.Since(start)
	fmt.Printf("sortSolution took %s\n", elapsed)

	start = time.Now()
	fmt.Println(mapSolution(slice))
	elapsed = time.Since(start)
	fmt.Printf("mapSolution took %s\n", elapsed)
}

func mapSolution(inputSlice []int) (maxSoFar int) {
	hash := map[int]bool{}
	res, starting := 0, 0
	for _, el := range inputSlice {
		hash[el] = true
	}
	maxSoFar = 1
	for k, _ := range hash {
		if !hash[k-1] {
			res = 1
			starting = k
			for hash[starting] {
				if res > maxSoFar {
					maxSoFar = res
				}
				starting += 1
				res += 1

			}
		}
	}
	return
}

func sortSolution(inputSlice []int) (maxSoFar int) {
	maxSoFar = 1
	maxEndingHere := 1
	sort.Ints(inputSlice)
	lastElement := inputSlice[0]
	for i := 1; i < len(inputSlice); i++ {
		if inputSlice[i]-lastElement == 1 {
			maxEndingHere++
			if maxSoFar < maxEndingHere {
				maxSoFar = maxEndingHere
			}
		} else if inputSlice[i] != lastElement {
			maxEndingHere = 1
		}
		lastElement = inputSlice[i]
	}
	return
}
