#include <iostream>
#include <algorithm>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll t, n, L, R, a[N + 10];

ll bisect_left(ll i, ll target) {
    ll p = i + 1, q = n;
    while (p < q) {
        ll mid = p + q >> 1;
        if (a[i] + a[mid] >= target)
            q = mid;
        else
            p = mid + 1;
    }
    return p;
}

int main() {
    scanf("%lld", &t); while (t--) {
        scanf("%lld%lld%lld", &n, &L, &R);
        for (ll i = 0; i < n; i++)
            scanf("%lld", a + i);
        sort(a, a + n);
        ll ans = 0;
        for (ll i = 0; i < n; i++) {
            ll l = bisect_left(i, L), r = bisect_left(i, R + 1) - 1;
            ans += r - l + 1;
        }
        printf("%lld\n", ans);
    }
    return 0;
}
