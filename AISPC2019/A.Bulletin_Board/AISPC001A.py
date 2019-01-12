import sys

def main():
    N = int(input())
    H = int(input())
    W = int(input())

    count = 0
    for i in range(N):
        for j in range(N):
            if i + H <= N and j + W <= N:
                count += 1
    print(count)

if __name__ == "__main__":
    main()
