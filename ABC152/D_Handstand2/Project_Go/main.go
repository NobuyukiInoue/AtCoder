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

type pair struct {
	first int
	last  int
}

func getpair(num int) pair {
	first := 0
	last := num % 10
	for num > 0 {
		if num < 10 {
			first = num
			break
		}
		num /= 10
	}
	return pair{first, last}
}

func main() {
	var n, count int

	sc.Split(bufio.ScanWords)
	n = next()

	var freq map[pair]int

	freq = make(map[pair]int, 0)
	for A := 1; A <= n; A++ {
		p := getpair(A)
		freq[p]++
	}
	for B := 1; B <= n; B++ {
		p := getpair(B)
		q := pair{p.last, p.first}
		count += freq[q]
	}

	fmt.Println(count)
}

/*
func main() {
	var n, count int

	sc.Split(bufio.ScanWords)
	n = next()

	var stra, strb string
	var lena, lenb int
	for a := 1; a <= n; a++ {
		stra = strconv.Itoa(a)
		lena = len(stra)
		for b := 1; b <= n; b++ {
			strb = strconv.Itoa(b)
			lenb = len(strb)
			if stra[lena-1] == strb[0] && stra[0] == strb[lenb-1] {
				count++
			}
		}
	}

	fmt.Println(count)
}
*/
