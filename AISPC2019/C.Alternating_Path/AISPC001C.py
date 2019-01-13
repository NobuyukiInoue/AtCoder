import sys

class Solution:
    def __init__(self):
        self.pair = []
        self.MAX_N = 400
        self.di = [1, 0, -1, 0]
        self.dj = [0, 1, 0, -1]

    def input(self):
        flds = input().split(' ')
        self.H, self.W = int(flds[0]), int(flds[1])

        self.data = [""]*self.H
        for i in range(self.H):
            self.data[i] = input()

    def isin(self, i, j):
        return i >= 0 and i < self.H and j >= 0 and j < self.W

    def check(self, i, j, d):
        ni = i + self.di[d]
        nj = j + self.dj[d]
        if not self.isin(ni, nj):
            return False
        if self.data[i][j] == self.data[ni][nj]:
            return False
        return True

    def solve(self):
        ans = 0
        b, w = 0, 0
        MAX_N = 400
        self.used = [[0]*MAX_N]*MAX_N

        for i in range(self.H):
            for j in range(self.W):
                if self.used[i][j]:
                    continue
                b, w = 0, 0
                que = []
                self.used[i][j] = True
                temp_p = [i, j]
                que.append(temp_p)
                while len(que) > 0:
                    p = que[0]
                    que.pop()
                    ci, cj = p[0], p[1]
                    if self.data[i][j] == "#":
                        b += 1
                    else:
                        w += 1
                    for d in range(0, 4):
                        if self.check(ci, cj, d):
                            continue
                        ni = ci + self.di[d]
                        nj = cj + self.dj[d]

                        if self.used[ni][nj]:
                            continue
                        
                        self.used[ni][nj] = True
                ans += b * w
        return ans


    def print_dart(self):
        for i in range(len(self.data)):
            print(self.data[i])


def main():
    sl = Solution()
    sl.input()
    sl.print_dart()
    result = sl.solve()
    print(result)

if __name__ == "__main__":
    main()
