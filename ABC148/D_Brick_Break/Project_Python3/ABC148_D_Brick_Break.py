N = int(input())
a = list(map(int, input().split()))

print(N)
print(a)

pos, v = 1, 0

for i in range(N):
   if pos == a[i]:
      v += 1
      pos += 1

if v == 0:
   print(-1)
else:
   print(N - v + 1)
