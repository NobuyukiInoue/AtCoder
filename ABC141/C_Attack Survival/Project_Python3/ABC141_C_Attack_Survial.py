n,k,q = map(int,input().split())
ans = [0]*n
for i in range(q):
    idx = int(input()) -1
    ans[idx] += 1
for i in ans:
    print('Yes' if q-k<i else 'No')
