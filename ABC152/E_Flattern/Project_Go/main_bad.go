package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func next() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	var n int

	sc.Split(bufio.ScanWords)
	n = next()

	var a, b []int
	a, b = make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = next()
	}

	lcdA := a[0]
	for i := 1; i < n; i++ {
		lcdA = lcd(lcdA, a[i])
	}

	sumb := 0
	for i := 0; i < n; i++ {
		b[i] = lcdA / a[i]
		sumb += b[i]
	}

	fmt.Println(sumb % (1000000000 + 7))
}

func lcd(a int, b int) int {
	x := a * b
	if a < b {
		a, b = b, a
	}
	r := a % b
	for r != 0 {
		a, b = b, r
		r = a % b
	}
	return x / b
}
