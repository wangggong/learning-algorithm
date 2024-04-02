// 参考: https://www.luogu.com.cn/blog/endlesscheng/solution-at-tenka1-2019-d
//
// 重点掌握 0-1 背包的变形: 至少装满 & 恰好装满
#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 300;
const ll MOD = 998244353;
ll n, v, S, pow3 = 1, f[N * N + 10], g[N * N + 10];

int main() {
    f[0] = g[0] = 3;
    scanf("%lld", &n);
    while (n--) {
        scanf("%lld", &v);
        S += v;
        for (ll i = S; i >= 0; i--) {
            f[i] = (f[i] * 2 + f[max(i - v, 0ll)]) % MOD;   // 至少装满
            if (i >= v)
                g[i] = (g[i] + g[i - v]) % MOD;             // 恰好装满
        }
        pow3 = pow3 * 3 % MOD;
    }
    ll dup = (S & 1) ? 0 : g[S >> 1];
    ll ans = (pow3 - (f[(S + 1) >> 1] - dup) % MOD + MOD) % MOD;
    printf("%lld", ans);
    return 0;
}
