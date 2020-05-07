package main

import (
	"fmt"
)

func main() {
	// 1492ms
	var n int

	fmt.Scan(&n)
	a := make([]int, n)
	dic := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if _, exist := dic[a[i]]; exist {
			fmt.Println("NO")
			return
		}
		dic[a[i]]++
	}

	fmt.Println("YES")
}

/*
	for k, v := range dic {
		fmt.Printf("dic[%d] = %d\n", k, v)
	}

	for _, v := range dic {
		if v > 1 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
*/
