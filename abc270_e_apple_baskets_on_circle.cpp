#include <iostream>

using namespace std;
typedef long long ll;

const ll N = 1e5;
ll n, k;
ll a[N + 10];

bool check(ll cnt) {
    ll left = k;
    for (ll i = 0; i < n; i++) {
        left -= min(cnt, a[i]);
        if (left <= 0)
            break;
    }
    return left <= 0;
}

int main() {
    scanf("%lld %lld", &n, &k);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    ll p = 0, q = k + 1;
    while (p < q) {
        ll mid = (p + q) >> 1;
        if (check(mid))
            q = mid;
        else
            p = mid + 1;
    }
    for (ll i = 0; i < n; i++) {
        ll cur = min(a[i], p - 1);
        k -= cur, a[i] -= cur;
    }
    for (ll i = 0; i < n && k; i++)
        if (a[i])
            a[i]--, k--;
    for (ll i = 0; i < n; i++)
        printf("%lld ", a[i]);
    return 0;
}
