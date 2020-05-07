#include <cstdio>
#include <array>
#include <utility>

template <unsigned int N> struct myprimes {
    unsigned int minPrime[N + 1];
    unsigned int Primes[N + 1];
    unsigned int size_;

    constexpr myprimes(void): minPrime(), Primes(), size_(0U) {
        minPrime[0] = 1; minPrime[1] = 1;
        for (int even = 2, odd = 3; odd <= N; even += 2, odd += 2) {
            minPrime[even] = 2;
            minPrime[odd] = odd;
        }

        if (!(N & 1))
            minPrime[N] = 2;

        Primes[size_++] = 2;

        for (int x = 3; x <= N; x += 2) {
            if (minPrime[x] == x)
                Primes[size_++] = x;
            for (int y = 0; y < size_ and Primes[y] <= minPrime[x] and x * Primes[y] <= N; ++y)
                    minPrime[x * Primes[y]] = Primes[y];
        }
    }

    // 範囲for文用の begin() と end()
    const unsigned int *begin(void) const noexcept {
        return Primes;
    }

    const unsigned int *end(void) const noexcept {
        return Primes + size_;
    }

    const unsigned int size(void) const {
        return this.size_;
    }
};

using uint = unsigned int;
constexpr uint N = 1e6;
constexpr uint mod = 1e9 + 7;
constexpr uint M = mod - 2;
constexpr myprimes<N> P;
std::array<uint, N + 1> Count;

// a^nを返す関数
const uint power(uint a, uint n)noexcept{
    unsigned long long res = 1, waiting = a;

    while (n) {
        if (n & 1) {
            res *= waiting;
            res %= mod;
        }
        waiting *= waiting; waiting %= mod; 
        n >>= 1;
    }

    return res;
}

uint fastRead(void) {
    char c;

    while (true) {
        c = getchar();
        if('0' <= c and c <= '9')
            break;
    }

    uint res = c - '0';
    while (true) {
        c = getchar();
        if (c < '0' or c > '9')
            break;
        res = 10 * res + (c - '0');
    }

    return res;
}


int main(void) {
    uint n = fastRead();
    uint inv_sum = 0;

    for (uint _ = 0; _ < n; ++_){
        uint a = fastRead();
        inv_sum += power(a, M);

        if(inv_sum >= mod)
            inv_sum -= mod;

        while (a - 1U) {
            uint fac = P.minPrime[a];
            uint cnt = 1; a /= fac;

            while (a % fac == 0)
                a /= fac, cnt++;

            if (Count[fac] < cnt)
                Count[fac] = cnt; 
        }
    }

    unsigned long long int lcm = 1;

    for(unsigned int x : P)
        if(Count[x])
            (lcm *= power(x, Count[x])) %= mod;

    unsigned int res = lcm * inv_sum % mod;
    
    printf("%u\n", res);
    return 0;
}
