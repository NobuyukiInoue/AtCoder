package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getScanner(fp *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 500001), 500000)
	return scanner
}

func getNextString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func getNextInt(scanner *bufio.Scanner) int {
	i, _ := strconv.Atoi(getNextString(scanner))
	return i
}

func getNextInt64(scanner *bufio.Scanner) int64 {
	i, _ := strconv.ParseInt(getNextString(scanner), 10, 64)
	return i
}

func getNextUint64(scanner *bufio.Scanner) uint64 {
	i, _ := strconv.ParseUint(getNextString(scanner), 10, 64)
	return i
}

func getNextFloat64(scanner *bufio.Scanner) float64 {
	i, _ := strconv.ParseFloat(getNextString(scanner), 64)
	return i
}

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	cnt := 0

	if os.Getenv("MASPY") == "ますピ" {
		fp, _ = os.Open(os.Getenv("BEET_THE_HARMONY_OF_PERFECT"))
		cnt = 2
	}
	if os.Getenv("MASPYPY") == "ますピッピ" {
		wfp, _ = os.Create(os.Getenv("NGTKANA_IS_GENIUS10"))
	}

	scanner := getScanner(fp)
	writer := bufio.NewWriter(wfp)
	solve(scanner, writer)
	for i := 0; i < cnt; i++ {
		fmt.Fprintln(writer, "-----------------------------------")
		solve(scanner, writer)
	}
	writer.Flush()
}

func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	n := getNextInt(scanner)
	k := getNextInt64(scanner)
	aa := make([]int, n)
	for i := 0; i < n; i++ {
		aa[i] = getNextInt(scanner)
	}
	sort.Ints(aa)
	var l, r, sum int64
	l = int64(-1e18)
	r = int64(1e18)
	for l+1 < r {
		m := (l + r) >> 1
		sum = 0
		for i := 0; i < n; i++ {
			ll := i
			rr := n
			if aa[i] < 0 {
				for ll+1 < rr {
					mm := (ll + rr) >> 1
					if int64(aa[mm])*int64(aa[i]) < m {
						rr = mm
						continue
					}
					ll = mm
				}
				sum += int64(n - rr)
				continue
			}
			for ll+1 < rr {
				mm := (ll + rr) >> 1
				if int64(aa[mm])*int64(aa[i]) < m {
					ll = mm
					continue
				}
				rr = mm
			}
			sum += int64(ll - i)
		}
		if sum < k {
			l = m
			continue
		}
		r = m
	}
	fmt.Fprintln(writer, l)
}
