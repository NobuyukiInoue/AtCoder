package main

import (
	"fmt"
)

func main() {
	var a, b, c, x int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&x)

	count := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if i*500+j*100+k*50 == x {
					count++
				}

			}
		}
	}

	fmt.Println(count)
}
