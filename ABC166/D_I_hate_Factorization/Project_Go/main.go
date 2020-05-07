package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calc(x int) []int {
	for a := -118; a <= 119; a++ {
		for b := -119; b <= 118; b++ {
			if a*a*a*a*a-b*b*b*b*b == x {
				return []int{a, b}
			}
		}
	}
	return []int{}
}

func main() {
	x := readInt()
	ans := calc(x)
	fmt.Printf("%d %d\n", ans[0], ans[1])
}

const (
	ioBufferSize = 1 * 1024 * 1024 // 1 MB
)

var stdinScanner = func() *bufio.Scanner {
	result := bufio.NewScanner(os.Stdin)
	result.Buffer(make([]byte, ioBufferSize), ioBufferSize)
	result.Split(bufio.ScanWords)
	return result
}()

func readString() string {
	stdinScanner.Scan()
	return stdinScanner.Text()
}

func readInt() int {
	result, err := strconv.Atoi(readString())
	if err != nil {
		panic(err)
	}
	return result
}
