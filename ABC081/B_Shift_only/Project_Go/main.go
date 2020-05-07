package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	nums := make([]int, n)
	count := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
		if nums[i]%2 == 0 {
			count++
		}
	}
	/*
		n = 3
		data := []int{8, 12, 40}
		n = 4
		data := []int{5, 6, 8, 10}
		nums := make([]int, n)
		count := 0
		for i := 0; i < n; i++ {
			nums[i] = data[i]
			if nums[i]%2 == 0 {
				count++
			}
		}
	*/

	res := 0
	for count == n {
		count = 0
		for i := 0; i < n; i++ {
			nums[i] /= 2
			if nums[i]%2 == 0 {
				count++
			}
		}
		res++
	}

	fmt.Println(res)
}
