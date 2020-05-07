package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	x := make([]int, n)
	sum, lenx := 0, len(x)
	for i := 0; i < lenx; i++ {
		fmt.Scan(&x[i])
		sum += x[i]
	}

	p := int((float64(sum)/(float64)(lenx) + 0.5))
	hp := 0
	for i := 0; i < lenx; i++ {
		hp += (x[i] - p) * (x[i] - p)
	}

	fmt.Println(hp)
}
