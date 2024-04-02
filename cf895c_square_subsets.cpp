// 参考: https://codeforces.com/problemset/submission/895/249221937
//
// 状压 DP, 但是需要压缩为 Counter.
// 同时学会一个组合恒等式:
// `C_n^0 + C_n^2 + C_n^4 + ... + C_n^{n / 2 * 2} = C_n^1 + C_n^3 + C_n^5 + ... + C_n^{(n - 1) / 2 * 2 + 1} = 2 ^ (n - 1)`
// 证明可以参考 `(1 + x)^n` 的二项式展开, 取 `x = 1, -1` 并求和, 也就是左边的部分了.
#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll B = 19, M = 70, N = 1e5, MOD = 1e9 + 7;
ll primes[] = {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67};
ll n, a, c[M + 10], mask[M + 10], pow2[N + 10], dp[(1 << B) + 10], ndp[(1 << B) + 10];

int main() {
    pow2[0] = 1; for (ll i = 0; i < N; i++)
        pow2[i + 1] = pow2[i] * 2 % MOD;
    for (ll i = 2; i <= M; i++)
        for (ll k = i, j = 0; j < B; j++)
            for (; k % primes[j] == 0; k /= primes[j])
                mask[i] ^= 1ll << j;
    scanf("%lld", &n); for (ll i = 0; i < n; i++)
        scanf("%lld", &a), c[a]++;
    dp[0] = 1; for (ll i = 1; i <= M; i++)
        if (c[i]) {
            memset(ndp, 0, sizeof(ndp));
            for (ll j = 0; j < 1 << B; j++)
                ndp[j] += dp[j] * pow2[c[i] - 1] % MOD,
                ndp[j] %= MOD,
                ndp[j ^ mask[i]] += dp[j] * pow2[c[i] - 1] % MOD,
                ndp[j] %= MOD;
            swap(dp, ndp);
        }
    printf("%lld", (dp[0] - 1 + MOD) % MOD);
    return 0;
}
