A, B = map(int, input().split())
 
diff = abs(B - A)
value = {0: 0, 1: 1, 2: 2, 3: 3, 4: 2, 5: 1, 6: 2, 7: 3, 8: 3, 9: 2}
 
print(diff // 10 + value[diff % 10])
