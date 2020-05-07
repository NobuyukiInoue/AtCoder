package main

import (
	"fmt"
	"math"
)

func main() {
	var h, n int

	fmt.Scan(&h, &n)
	res := make([]int, h+1)
	for k := 0; k <= h; k++ {
		res[k] = math.MaxInt64
	}

	for i := 0; i < n; i++ {
		var HP, MP int
		fmt.Scan(&HP, &MP)
		for k := 1; k <= h; k++ {
			if k <= HP {
				if res[k] > MP {
					res[k] = MP
				}
			} else {
				use := res[k-HP] + MP
				if res[k] > use {
					res[k] = use
				}
			}
		}
	}
	fmt.Println(res[h])
}
