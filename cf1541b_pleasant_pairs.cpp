#include <iostream>
#include <algorithm>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 1e5;
ll t, n, v;
pll a[N + 10];

int main() {
    for (scanf("%lld", &t); t--; ) {
        scanf("%lld", &n);
        for (ll i = 1; i <= n; i++)
            scanf("%lld", &v), a[i] = {v, i};
        sort(a + 1, a + n + 1);
        ll ans = 0;
        for (ll i = 1; i <= n; i++)
            for (ll j = i + 1; j <= n && a[j].first <= (n << 2) / a[i].first; j++)
                if (a[i].first * a[j].first == a[i].second + a[j].second)
                    ans++;
        printf("%lld\n", ans);
    }
    return 0;
}
