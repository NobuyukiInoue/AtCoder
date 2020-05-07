package main

import (
	"fmt"
	"strings"
)

func main() {

	// erasedream yes
	// dreameraser yes
	// eraserdreameraser
	// dreamerer no
	// dream dreamer erase eraser
	var s string
	fmt.Scan(&s)

	s = strings.Replace(s, "eraser", "", -1)
	s = strings.Replace(s, "erase", "", -1)
	s = strings.Replace(s, "dreamer", "", -1)
	s = strings.Replace(s, "dream", "", -1)

	if len(s) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/*
var s string
var words []string

func main() {
	//fmt.Scan(&s)
	s = "dreamerer"
	words = []string{"dream", "dreamer", "erase", "eraser"}

	if dfs(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func dfs(s string) bool {
	for i := 0; i < len(words); i++ {
		if len(s) < len(words[i]) {
			continue
		}
		temp := strings.Replace(s, words[i], "", 1)

		if len(temp) == 0 {
			return true
		}

		for j := 0; j < len(words); j++ {
			if len(temp) >= len(words[j]) {
				return dfs(temp)
			}
		}
	}
	return false
}
*/
