package main

import (
	"fmt"
)

type Position struct {
	time int
	x    int
	y    int
}

func main() {
	var n int

	fmt.Scan(&n)
	pst := make([]Position, n+1)
	pst[0] = Position{0, 0, 0}
	for i := 1; i < n+1; i++ {
		fmt.Scan(&pst[i].time, &pst[i].x, &pst[i].y)
	}
	/*
		n = 1
		pst := make([]Position, n+1)
		pst[1] = Position{2, 100, 100}
		for i := 0; i < n+1; i++ {
			fmt.Printf("%d = (%d, %d)\n", pst[i].time, pst[i].x, pst[i].y)
		}
	*/

	for i := 1; i < n+1; i++ {
		ts := pst[i].time - pst[i-1].time
		dxy := myAbs(pst[i].x-pst[i-1].x) + myAbs(pst[i].y-pst[i-1].y)
		if ts < dxy || ts%2 != dxy%2 {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

func myAbs(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}
