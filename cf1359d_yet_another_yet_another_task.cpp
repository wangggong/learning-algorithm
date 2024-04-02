#include <iostream>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 1e5, MIN = -30, MAX = 30;
ll n, a[N + 10], ans;

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    for (ll k = MIN; k <= MAX; k++)
        for (ll p = 0, q = 0; p < n; p = q) {
            for (; p < n && a[p] > k; p++);
            if (p == n)
                break;
            ll cur = 0, mx = k;
            for (q = p; q < n && a[q] <= k; q++)
                cur = max(cur, 0ll) + a[q], mx = max(mx, cur);
            ans = max(ans, mx - k);
        }
    printf("%lld\n", ans);
    return 0;
}
