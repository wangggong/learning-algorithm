/*
 * 
 * // 朴素 DP
 * #include <iostream>
 * 
 * using namespace std;
 * typedef long long ll;
 * const ll N = 50;
 * const ll MOD = 998244353;
 * ll n, m, k, dp[N * N + 10], ans;
 * 
 * int main() {
 *     scanf("%lld %lld %lld", &n, &m, &k);
 *     dp[0] = 1;
 *     while (n--)
 *         for (ll i = k; i >= 0; i--) {
 *             dp[i] = 0;
 *             for (ll j = 1; j <= m && j <= i; j++)
 *                 dp[i] = (dp[i] + dp[i - j]) % MOD;
 *         }
 *     for (ll i = 1; i <= k; i++)
 *         ans = (ans + dp[i]) % MOD;
 *     printf("%lld", ans);
 *     return 0;
 * }
 */

// 参考: https://atcoder.jp/contests/abc248/submissions/45035045
//
// 后缀和优化 DP
#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 50;
const ll MOD = 998244353;
ll n, m, k, suf[N * N + 10], ans;

int main() {
    scanf("%lld %lld %lld", &n, &m, &k);
    suf[0] = 1;
    while (n--)
        for (ll i = k; i >= 0; i--)
            suf[i] = (suf[i + 1] - (suf[i] - suf[max(i - m, 0ll)]) % MOD + MOD) % MOD;
    printf("%lld", suf[0]);
    return 0;
}
