package main

import (
	"fmt"
)

func main() {
	var n, a, b int

	fmt.Scan(&n, &a, &b)

	total := 0
	for i := 0; i <= n; i++ {
		sum := 0
		val := i
		for val > 0 {
			sum += val % 10
			val /= 10
		}

		if a <= sum && sum <= b {
			total += i
		}
	}
	fmt.Println(total)
}
