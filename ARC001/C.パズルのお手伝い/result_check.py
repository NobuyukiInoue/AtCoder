import sys
import ARC001C

def main():
    flds = [""]*8

    for i in range(len(flds)):
        flds[i] = input()
        if flds[i] == "No Answer":
            print(flds[i])
            exit(1)

    qs = ARC001C.QueenSet(flds)
    qs.flds = qs.org_flds[:]
    if qs.isStartEnable():
        print("OK.")
    else:
        print("BAD Answer")
    
    for temp in flds:
        print(temp)

if __name__ == "__main__":
    main()
