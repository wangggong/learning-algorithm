#include <iostream>
#include <cstring>
#include <algorithm>

using namespace std;
typedef long long ll;

const ll N = 3e5;
ll t, n, h[N + 10], tmp[N + 10];

bool check(ll target, ll k) {
    ll even = k / 2, odd = k - even;
    memcpy(tmp, h, n * sizeof(ll));
    for (ll i = 0; i < n; i++)
        if (even > 0) {
            ll cur = min((target - tmp[i]) / 2, even);
            even -= cur, tmp[i] += 2 * cur;
        }
    for (ll i = 0; i < n; i++)
        if (odd > 0) {
            ll cur = min(target - tmp[i], odd);
            odd -= cur, tmp[i] += cur;
        }
    for (ll i = 0; i < n; i++)
        if (tmp[i] < target)
            return false;
    return true;
}

ll get(ll target) {
    ll tot = 0;
    for (ll i = 0; i < n; i++)
        tot += target - h[i];
    ll p = 0, q = tot * 2 - 1;
    while (p < q) {
        ll mid = (p + q) >> 1;
        if (check(target, mid))
            q = mid;
        else
            p = mid + 1;
    }
    return p;
}

int main() {
    scanf("%lld", &t);
    while (t--) {
        scanf("%lld", &n);
        for (ll i = 0; i < n; i++)
            scanf("%lld", h + i);
        sort(h, h + n);
        printf("%lld\n", min(get(h[n - 1]), get(h[n - 1] + 1)));
    }
    return 0;
}
