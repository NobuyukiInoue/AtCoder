import sys

def main():
    N, K, Q = map(int, input().split())

    A = [int(input()) for _ in range(Q)]
#   print(A)

    pt = [K - Q]*(N + 1)
#   print(pt)

    for i in range(Q):
        for n in range(len(pt)):
            if n == A[i]:
                pt[n] += 1
#       print("{0}: {1}".format(i, pt[1:]))

#   print(pt)
    for n in range(1, len(pt)):
        if pt[n] > 0:
            print("Yes")
        else:
            print("No")

if __name__ == "__main__":
    main()
