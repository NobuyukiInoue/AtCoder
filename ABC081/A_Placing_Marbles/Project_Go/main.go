package main

import (
	"fmt"
)

func main() {
	var grid string

	fmt.Scan(&grid)
	count := 0
	for i := 0; i < len(grid); i++ {
		if grid[i] == '1' {
			count++
		}
	}

	fmt.Printf("%d\n", count)
}
