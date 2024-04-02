a, b, k = [int(x) for x in input().split(' ')]

def comb(n, k):
    return factorial(n) // factorial(k) // factorial(n - k)

def factorial(n):
    ans = 1
    for i in range(1, n + 1):
        ans *= i
    return ans

left_a, left_b = a, b
strs = []
left_a, left_b = a, b
for i in range(a + b):
    if left_a == 0:
        strs += ['b']
        left_b -= 1
        continue
    if left_b == 0:
        strs += ['a']
        left_a -= 1
        continue
    c = comb(left_a + left_b - 1, left_b)
    if k > c:
        strs += ['b']
        left_b -= 1
        k -= c
    else:
        strs += ['a']
        left_a -= 1

print(''.join(strs))
