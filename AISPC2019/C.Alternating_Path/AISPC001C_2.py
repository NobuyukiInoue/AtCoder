import sys

sys.setrecursionlimit(10 ** 9)

H, W = map(int, input().split())
inputs = [input() for _ in range(H)]

visited = [[False for _ in range(W)] for _ in range(H)]
vectors = [[0, -1], [-1, 0], [0, 1], [1, 0]]
ans = 0

def dfs(h, w):
    # 入力された値はindex out of rangeにならないvalidationが掛けられている
    if visited[h][w]:
        return (0, 0)
    
    visited[h][w] = True

    if inputs[h][w] == '#':
        black, white = 1, 0 
    else:
        black, white = 0, 1

    for vector in vectors:
        next_h = h + vector[0]
        next_w = w + vector[1]
        if not (0 <= next_h < H and 0 <= next_w < W):
            continue
        next_color = inputs[next_h][next_w]
        if inputs[h][w] != next_color:
            p_black, p_white = dfs(next_h, next_w)
            black += p_black
            white += p_white

    return (black, white)

for h in range(H):
    for w in range(W):
        black, white = dfs(h, w)
        ans += black * white  

print(ans)
