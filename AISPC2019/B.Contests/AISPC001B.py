import sys

def main():
    N = int(input())

    flds = input().split(' ')
    pt_A, pt_B = int(flds[0]), int(flds[1])

    flds = input().split(' ')
    points = [-1]*N
    for i in range(N):
        points[i] = int(flds[i])

    e, m, h = 0, 0, 0
    for i in range(N):
        if points[i] <= pt_A:
            e += 1
        elif points[i] <= pt_B:
            m += 1
        else:
            h += 1
    ans = min(e, m, h)
    print(ans)

def main_work():
    N = int(input())

    flds = input().split(' ')
    pt_A, pt_B = int(flds[0]), int(flds[1])

    flds = input().split(' ')
    points = [-1]*N
    for i in range(N):
        points[i] = int(flds[i])

    count, isSet_A, isSet_B, isSet_C = 0, False, False, False
    for i in range(N):
        if points[i] < 0:
            continue

        if isSet_A == False and 0 <= points[i] and points[i] <= pt_A and pt_A < 20:
            isSet_A = True
            points[i] = -1
        if isSet_B == False and pt_A + 1 <= points[i] and points[i] <= pt_B and pt_B < 20:
            isSet_B = True
            points[i] = -1
        if isSet_C == False and pt_B + 1 <= points[i] and pt_B < 20:
            isSet_C = True
            points[i] = -1
        
        if isSet_A and isSet_B and isSet_C:
            count += 1
            isSet_A, isSet_B, isSet_C = False, False, False

    print(count)

if __name__ == "__main__":
    main()
