#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll t, n, m, v, c[N + 10];

bool check(ll k) {
    ll ans = 0;
    for (ll i = 1; i <= n; i++) {
        ll t = min(c[i], k);
        ans += t + (k - t) / 2;
    }
    return ans >= m;
}

int main() {
    for (scanf("%lld", &t); t; t--) {
        memset(c, 0, sizeof c);
        scanf("%lld %lld", &n, &m);
        for (ll i = 0; i < m; i++)
            scanf("%lld", &v), c[v]++;
        ll p = 0, q = (m + 1) << 1;
        while (p < q) {
            ll mid = (p + q) >> 1;
            if (check(mid))
                q = mid;
            else
                p = mid + 1;
        }
        printf("%lld\n", p);
    }
    return 0;
}
