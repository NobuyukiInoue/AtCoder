package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := readInt()
	m := readInt()
	h := make([]int, n)
	good := make([]bool, n)
	for i := 0; i < n; i++ {
		h[i] = readInt()
		good[i] = true
	}
	for i := 0; i < m; i++ {
		a_i := readInt() - 1
		b_i := readInt() - 1
		if h[a_i] > h[b_i] {
			good[b_i] = false
		} else if h[b_i] > h[a_i] {
			good[a_i] = false
		} else {
			good[a_i] = false
			good[b_i] = false
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if good[i] {
			ans++
		}
	}

	fmt.Println(ans)
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
