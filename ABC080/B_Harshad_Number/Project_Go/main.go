package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	fx := 0
	temp := n
	for temp > 0 {
		fx += temp % 10
		temp /= 10
	}

	if n%fx == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
