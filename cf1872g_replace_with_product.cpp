// 参考: https://codeforces.com/contest/1872/submission/244558745
#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 2e5;
ll t, n, a, sum[N + 10], mul[N + 10], pos[N + 10], idx, mx, l, r;
pll ans;

pll get() {
    memset(sum, 0, sizeof sum), memset(mul, 0, sizeof mul), mul[1] = 1;
    idx = 0, mx = 0, l = 1, r = 1;
    bool valid = false;
    for (ll i = 1; i <= n; i++) {
        scanf("%lld", &a), sum[i + 1] = sum[i] + a, mul[i + 1] = mul[i] * a;
        if (a == 1)
            continue;
        pos[idx++] = i;
        if (mul[i + 1] >= n * 2)
            valid = true;
    }
    if (valid)
        return {pos[0], pos[idx - 1]};
    for (ll i = 0; i < idx; i++)
        for (ll j = i; j < idx; j++) {
            ll c = mul[pos[j] + 1] / mul[pos[i]] + sum[pos[i]] + (sum[n] - sum[pos[j] + 1]);
            if (c > mx)
                l = pos[i], r = pos[j], mx = c;
        }
    return {l, r};
}

int main() {
    for (scanf("%lld", &t); t--; )
        scanf("%lld", &n), ans = get(), printf("%lld %lld\n", ans.first, ans.second);
    return 0;
}
