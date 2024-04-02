#include <iostream>
#include <algorithm>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll t, n, k, a[N + 10], b, prefix[N + 10];

int main() {
    scanf("%lld", &t);
    while (t--) {
        scanf("%lld %lld", &n, &k);
        for (ll i = 0; i < n; i++)
            scanf("%lld", a + i);
        for (ll i = 0; i < n; i++)
            scanf("%lld", &b);
        sort(a, a + n);
        ll ans = 0;
        for (ll l = 0, r = 0; r < n; r++) {
            for (; a[r] - a[l] > k; l++);
            ll m = r - l + 1;
            ans = max(ans, m + prefix[l]);
            prefix[r + 1] = max(prefix[r], m);
        }
        printf("%lld\n", ans);
    }
    return 0;
}
