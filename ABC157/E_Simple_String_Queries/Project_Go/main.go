package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s []int
var c []int

type type1 struct {
	iq int
	cq rune
}

type type2 struct {
	lq int
	rq int
}

func main() {
	var n, q int
	var s string

	n = readInt()
	s = readString()
	q = readInt()
	//que1 := make([]type1, 0)
	//que2 := make([]type1, 0)
	if len(s) != n {
		fmt.Println("len(s) != n... Error")
		return
	}

	for i := 0; i < q; i++ {
		val1 := readInt()
		char2 := readString()
		//char2 := readChar()
		val2, err := strconv.Atoi(char2)
		if err != nil {
			//que1 = append(que1, type1{val1, char2})
			s = s[0:val1] + char2 + s[val1+1:]

		} else {
			//que2 = append(que2, type1{val1, val2})
			dic := make(map[rune]int, 0)
			for i := val1; i <= val2; i++ {
				dic[rune(s[i])]++
			}

			/*
				count := 0
				for k, v := range dic {
					count++
				}
			*/
			//fmt.Println(count)
			fmt.Println(len(dic))
		}
	}

	//fmt.Printf("n = %d, m = %d\n", n, m)
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
