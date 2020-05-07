package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	cards := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cards[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cards)))

	pointA, pointB := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			pointA += cards[i]
		} else {
			pointB += cards[i]
		}
	}

	fmt.Println(pointA - pointB)
}
