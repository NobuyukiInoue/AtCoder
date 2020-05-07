package main

import (
	"fmt"
	"strconv"
)

func solve(c []int, n int) string {
	if c[0] == 0 && n == 1 {
		return "0"
	} else if c[0] == 0 {
		return "-1"
	}

	if c[0] == -1 && n == 1 {
		return "0"
	}

	answer := ""
	if c[0] == -1 {
		answer = "1"
	} else {
		answer = strconv.Itoa(c[0])
	}
	for i := 1; i < n; i++ {
		if c[i] > 0 {
			answer += strconv.Itoa(c[i])
		} else {
			answer += "0"
		}
	}
	return answer
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = -1
	}

	result := true
	for i := 0; i < m; i++ {
		var p, q int
		fmt.Scan(&p, &q)
		if c[p-1] == -1 {
			c[p-1] = q
		} else if c[p-1] != q {
			result = false
			break
		}
	}

	if result == false {
		fmt.Println("-1")
	} else {
		fmt.Println(solve(c, n))
	}

}
