// 参考: https://www.luogu.com.cn/blog/endlesscheng/solution-at-diverta2019-e
//
// DP + "延迟" 计算优化, 属实厉害.
#include <iostream>

using namespace std;
typedef long long ll;

const ll D = 20;
const ll MOD = 1e9 + 7;
ll n, v, x, cnt0 = 1, f0[1 << D], f1[1 << D], pre0[1 << D];

ll pow(ll n, ll k) {
    ll ans = 1;
    for (; k; k >>= 1)
        ans = (k & 1) ? (ans * n) % MOD : ans, n = (n * n) % MOD;
    return ans;
}

int main() {
    for (ll i = 0; i < 1 << D; i++)
        f0[i] = 1;
    scanf("%lld", &n);
    while (n--) {
        scanf("%lld", &v);
        x ^= v;
        if (!x)
            cnt0++;
        else {
            f0[x] = (f0[x] + f1[x] * (cnt0 - pre0[x])) % MOD;
            f1[x] = (f0[x] + f1[x]) % MOD;
            pre0[x] = cnt0;
        }
    }
    ll ans = f0[x];
    if (!x) {
        ans = pow(2, cnt0 - 2);
        for (ll f : f1)
            ans = (ans + f) % MOD;
    }
    printf("%lld\n", ans);
    return 0;
}
