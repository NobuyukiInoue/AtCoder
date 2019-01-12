import sys

class Tree:
    def __init__(self, val, depth):
        self.val = val
        self.m1 = None
        self.m5 = None
        self.m10 = None
        self.p1 = None
        self.depth = depth

class Calc_Tree:
    def __init__(self):
        self.resultDepth = []

    def set_node(self, num, depth):
        if num == 0:
            self.resultDepth.append(depth)
            node = Tree(num, depth)
            return node

        node = Tree(num, depth)
        if num - 1 >= 0 and (num - 1) % 5 < 3:
            node.m1 = self.set_node(num - 1, depth + 1)
        if num - 5 >= 0 and num % 10 == 5:
            node.m5 = self.set_node(num - 5, depth + 1)
        if num - 10 >= 0:
            node.m10 = self.set_node(num - 10, depth + 1)
        if num % 10 > 7 or num % 5 > 2:
            node.p1 = self.set_node(num + 1, depth + 1)
        return node

def main():
    flds = input().split(" ")
    s_temp, d_temp = int(flds[0]), int(flds[1])
    #print("%d %d" %(s_temp, d_temp))
    least_temp = abs(d_temp - s_temp)
    calc_t = Calc_Tree()
    root = calc_t.set_node(least_temp, 0)
    print(min(calc_t.resultDepth))

if __name__ == "__main__":
    main()
