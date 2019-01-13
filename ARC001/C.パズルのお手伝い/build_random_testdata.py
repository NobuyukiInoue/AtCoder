import sys
import random
import numpy as np
import ARC001C

def get_rand_int_3():
    nums = [-1]*3
    for i in range(3):
        check_result = True
        while check_result:
            result =  np.random.randint(0, 8)
            if result in nums:
                continue
            else:
                check_result = False
        nums[i] = result
    return nums

def main():
    flds = ["........"]*8
    rows = get_rand_int_3()
    rows.sort()
    cols = get_rand_int_3()
    cols.sort()

    n = 0
    for i in range(0, 8):
        if n < 3:
            if i == rows[n]:
                if cols[n] == 0:
                    flds[i] = 'Q.......'
                elif cols[n] == 7:
                    flds[i] = '.......Q'
                else:
                    flds[i] = flds[i][0:cols[n]] + 'Q' + flds[i][cols[n] + 1:8]
                n += 1

    for temp in flds:
        print(temp)

if __name__ == "__main__":
    main()
