package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, mochi int
	fmt.Scan(&n)
	m := make(map[int]int)

	for i := 0; i < n; i++ {
		fmt.Scan(&mochi)
		m[mochi]++
	}

	fmt.Println(len(m))
}

func main2() {
	var n int
	fmt.Scan(&n)
	mochi := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&mochi[i])
	}

	sort.Sort(sort.Reverse(sort.IntSlice(mochi)))

	count := 1
	for i := 1; i < n; i++ {
		if mochi[i] == mochi[i-1] {
			continue
		}
		count++
	}

	fmt.Println(count)
}

type Entry struct {
	size  int
	count int
}

type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	return (l[i].size < l[j].size)
}

func main3() {
	var n int

	fmt.Scan(&n)
	d := make([]int, n)

	mochi := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
		if _, exists := mochi[d[i]]; exists {
			mochi[d[i]]++
		} else {
			mochi[d[i]] = 1
		}
	}

	sortedMochi := List{}
	for k, v := range mochi {
		e := Entry{k, v}
		sortedMochi = append(sortedMochi, e)
	}

	sort.Sort(sortedMochi)

	/*
		for i := 0; i < len(sortedMochi); i++ {
			fmt.Printf("sortedMochi[%d] = %d\n", i, sortedMochi[i])
		}
	*/

	count := 1
	for i := len(sortedMochi) - 1; i > 0; i-- {
		if sortedMochi[i].size > sortedMochi[i-1].size {
			count++
		}
	}

	fmt.Println(count)
}
