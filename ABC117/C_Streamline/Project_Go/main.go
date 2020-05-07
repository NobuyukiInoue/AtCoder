package main

import (
	"fmt"
	"sort"
)

func main() {
	// 448ms
	var n, m int

	fmt.Scan(&n, &m)
	x := make([]int, m)
	distance := make([]int, m-1)

	for i := 0; i < m; i++ {
		fmt.Scan(&x[i])
	}

	sort.Sort(sort.IntSlice(x))

	for i := 0; i < m; i++ {
		//fmt.Printf("x[%d] = %d\n", i, x[i])
		if i > 0 {
			distance[i-1] = x[i] - x[i-1]
		}
	}

	/*
		for i := 0; i < m-1; i++ {
			fmt.Printf("distance[%d] = %d\n", i, distance[i])
		}
	*/

	sort.Sort(sort.IntSlice(distance))
	ans := 0
	for j := 0; j < m-1-n+1; j++ {
		ans += distance[j]
	}

	fmt.Println(ans)
}
