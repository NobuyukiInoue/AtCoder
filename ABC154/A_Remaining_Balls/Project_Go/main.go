package main

import (
	"fmt"
)

func main() {
	var s, t, u string
	var a, b int

	fmt.Scan(&s, &t)
	fmt.Scan(&a, &b)
	fmt.Scan(&u)
	if u == s {
		a--
	} else {
		b--
	}
	fmt.Println(a, b)
}
