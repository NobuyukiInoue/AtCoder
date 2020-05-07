package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int

	n = readInt()
	if n%2 == 1 {
		fmt.Println(n/2 + 1)
	} else {
		fmt.Println(n / 2)
	}
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
