package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Scan(&s)
	slen := len(s)
	res := make([]rune, slen)
	for i := 0; i < slen; i++ {
		res[i] = 'x'
	}

	fmt.Println(string(res))
}
