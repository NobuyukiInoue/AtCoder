package main

import (
	"fmt"
	"sort"
)

func main() {
	x := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&x[i])
	}
	//x := []int{2, 6, 3}

	sort.Sort(sort.IntSlice(x))

	if (3*x[2]-mySum(x))%2 == 0 {
		fmt.Println((3*x[2] - mySum(x)) / 2)
	} else {
		fmt.Println((3*(1+x[2]) - mySum(x)) / 2)
	}
}

func mySum(x []int) int {
	sum := 0
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}
