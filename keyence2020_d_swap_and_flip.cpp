#include <iostream>
#include <algorithm>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 18;
ll n, a[N + 10], b[N + 10], c[N + 10], d[N + 10], p[N + 10], q[N + 10], ans = 999;

void backtrack(ll k, bool odd) {
    if (k < n) {
        c[k] = a[k], d[k] = a[k];
        backtrack(k + 1, odd);
        c[k] = b[k], d[k] = -b[k];
        backtrack(k + 1, !odd);
        return;
    }
    memcpy(p, c, sizeof c);
    sort(p, p + n);
    memcpy(q, d, sizeof d);
    ll cur = 0;
    for (ll i = 0; i < n; i++) {
        bool valid = false;
        ll j = i;
        for (; j < n; j++)
            if (abs(q[j]) == p[i] && (j - i) % 2 == (q[j] < 0)) {
                valid = true;
                break;
            }
        if (!valid)
            return;
        cur += j - i;
        for (; j > i; j--)
            q[j - 1] *= -1, swap(q[j], q[j - 1]);
    }
    ans = min(ans, cur);
    return;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    for (ll i = 0; i < n; i++)
        scanf("%lld", b + i);
    backtrack(0, false);
    printf("%lld\n", ans == 999 ? -1 : ans);
    return 0;
}
