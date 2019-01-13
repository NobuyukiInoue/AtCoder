import sys

def main():
    N = int(input())
    results = input()
    counts = [0]*4
    for temp in results:
        if temp == '1':
            counts[0] += 1
        elif temp == '2':
            counts[1] += 1
        elif temp == '3':
            counts[2] += 1
        elif temp == '4':
            counts[3] += 1

#   print(max(counts))
    max = 0
    min = sys.maxsize
    target = -1
    for i in range(len(counts)):
        if counts[i] > max:
            max = counts[i]
        if counts[i] < min:
            min = counts[i]

    print("%d %d" %(max, min))

if __name__ == "__main__":
    main()
