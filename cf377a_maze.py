directions = [[0, 1], [0, -1], [1, 0], [-1, 0]]
n, m, k = [int(x) for x in input().split()]
ss = []
for i in range(n):
    ss += [[x for x in input()]]

Q = []
tot = 0
for i, s in enumerate(ss):
    for j, c in enumerate(s):
        if c == '.':
            tot += 1
            ss[i][j] = 'X'
            if len(Q) == 0:
                Q += [(i, j)]
k = tot - k
while len(Q) > 0:
    x, y = Q[0]
    Q = Q[1:]
    if ss[x][y] == '.':
        continue
    ss[x][y] = '.'
    k -= 1
    if k == 0:
        break
    for dx, dy in directions:
        nx, ny = x + dx, y + dy
        if 0 <= nx < n and 0 <= ny < m and ss[nx][ny] == 'X':
            Q += [(nx, ny)]
print('\n'.join([''.join(s) for s in ss]))
