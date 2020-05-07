package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	N := readInt()
	A := readInt()
	B := readInt()
	C := readInt()
	S := make([]string, N)
	for i := 0; i < N; i++ {
		S[i] = readString()
	}

	T := make([]string, N)
	ans := "Yes"
	for i, s := range S {
		if s == "AB" {
			if A == 0 && B == 0 {
				ans = "No"
				break
			}
			if A < B {
				A++
				B--
				T[i] = "A"
			} else if B < A {
				A--
				B++
				T[i] = "B"
			} else if i == N-1 || S[i+1] != "BC" {
				A++
				B--
				T[i] = "A"
			} else {
				A--
				B++
				T[i] = "B"
			}
		} else if s == "AC" {
			if A == 0 && C == 0 {
				ans = "No"
				break
			}
			if A < C {
				A++
				C--
				T[i] = "A"
			} else if C < A {
				A--
				C++
				T[i] = "C"
			} else if i == N-1 || S[i+1] != "BC" {
				A++
				C--
				T[i] = "A"
			} else {
				A--
				C++
				T[i] = "C"
			}
		} else {
			if B == 0 && C == 0 {
				ans = "No"
				break
			}
			if B < C {
				B++
				C--
				T[i] = "B"
			} else if C < B {
				B--
				C++
				T[i] = "C"
			} else if i == N-1 || S[i+1] != "AC" {
				B++
				C--
				T[i] = "B"
			} else {
				B--
				C++
				T[i] = "C"
			}
		}
	}
	fmt.Println(ans)
	if ans == "No" {
		return
	}
	for _, t := range T {
		fmt.Println(t)
	}
}

func myAbs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
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

func readChar() rune {
	stdinScanner.Scan()
	temp := stdinScanner.Text()
	return rune(temp[0])
}

func readInt() int {
	result, err := strconv.Atoi(readString())
	if err != nil {
		panic(err)
	}
	return result
}
