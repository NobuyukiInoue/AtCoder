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

type myprimes struct {
	minPrime []int
	Primes   []int
	size     int
}

func Constructor(N int) *myprimes {
	this := new(myprimes)

	this.minPrime = make([]int, N+1)
	this.Primes = make([]int, N+1)
	this.size = 0

	this.minPrime[0], this.minPrime[1] = 1, 1
	even, odd := 2, 3
	for odd <= N {
		this.minPrime[even] = 2
		this.minPrime[odd] = odd
		even += 2
		odd += 2
	}

	if N&1 > 0 {
		this.minPrime[N] = 2
	}

	this.Primes[this.size] = 2
	this.size++

	var x, y int
	for x = 3; x <= N; x += 2 {
		if this.minPrime[x] == x {
			this.Primes[this.size] = x
			this.size++
		}
		for y = 0; y < this.size && this.Primes[y] <= this.minPrime[x] && x*this.Primes[y] <= N; y++ {
			this.minPrime[x*this.Primes[y]] = this.Primes[y]
		}
	}

	return this
}

func myPow(a int, n int) int {
	res := 1
	waitting := a
	for n > 0 {
		if n&1 != 0 {
			res *= waitting
			res %= mod
		}
		waitting *= waitting
		waitting %= mod
		n >>= 1
	}

	return res
}

var mod int
var M int

func main() {
	var n int
	mod = 1000000000 + 7
	M = mod - 2

	sc.Split(bufio.ScanWords)
	n = next()

	P := Constructor(n)
	Count := make(map[int]int, 0)

	var invsum int
	invsum = 0

	for i := 0; i < n; i++ {
		a := next()
		invsum += myPow(a, M)

		if invsum >= mod {
			invsum -= mod
		}

		for a-1 > 0 {
			fac := P.minPrime[a]
			cnt := 1
			a /= fac

			for a%fac == 0 {
				a /= fac
				cnt++
			}

			if Count[fac] < cnt {
				Count[fac] = cnt
			}
		}
	}

	lcm := 1

	for x := 0; x < n; x++ {
		if Count[x] != 0 {
			temp := myPow(x, Count[x]) % mod
			lcm *= temp
		}
	}

	res := lcm * invsum % mod
	fmt.Println(res)
}
