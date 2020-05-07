package main

import (
	"fmt"
)

const (
	// MOD is .
	MOD = 1000000000 + 7
	// MOD = 13
)

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)
	res := modPow(2, n) + 2*MOD - combMod(n, a) - combMod(n, b) - 1
	fmt.Println(res % MOD)
	// fmt.Println(combMod(n, a))
}

func combMod(n, k int) int {
	if k == 0 {
		return 1
	}
	res := n
	div := k
	for i := 1; i < k; i++ {
		res = res * (n - i) % MOD
		div = div * (k - i) % MOD
	}
	div = modPow(div, MOD-2)
	return res * div % MOD
}

func modPow(a, n int) int {
	if n == 0 {
		return 1
	}
	res := 1
	for p := n; p > 0; p >>= 1 {
		if p&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
	}
	return res
}

/*
func main() {
	var n, a, b int

	//fmt.Scan(&n, &a, &b)
	n, a, b = 1000000000, 141421, 173205
	sum, num := 0, 1
	for i := 0; i < n-2; {
		if num == a || num == b {
			num++
			continue
		}

		if num == 65 {
			fmt.Println(num)
		}
		sum += combination(n, num)
		num++
		i++
	}
	fmt.Println(sum % (1000000000 + 7))
}

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}
*/
