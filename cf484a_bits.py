def bitCnt1(n):
    ans = 0
    while n:
        ans, n = ans + 1, n & (n - 1)
    return ans


def getBits(l, r):
    if bitCnt1(r + 1) == 1:
        return r
    d = len(bin(r)[2:])
    k = 1 << d - 1
    return k - 1 if l < k else k + getBits(l - k, r - k)


t = int(input())

for _ in range(t):
    l, r = [int(x) for x in input().split()]
    print(getBits(l, r))
