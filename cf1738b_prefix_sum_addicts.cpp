#include <iostream>
#include <climits>

using namespace std;
typedef long long ll;
const ll N = 1e5;
ll t, n, k, a[N + 10];

bool check() {
    if (k == 1)
        return true;
    ll cur = 0, last = LLONG_MAX;
    for (ll i = k - 2; i >= 0; i--) {
        cur = a[i + 1] - a[i];
        if (cur > last)
            return false;
        last = cur;
    }
    return cur * (n - k + 1) >= a[0];
}

int main() {
    for (scanf("%lld", &t); t--; ) {
        scanf("%lld%lld", &n, &k);
        for (ll i = 0; i < k; i++)
            scanf("%lld", a + i);
        puts(check() ? "YES" : "NO");
    }
    return 0;
}
