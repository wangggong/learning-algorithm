#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll t, n, a[N + 10], b[N + 10];

bool check(ll k) {
    for (ll i = 0, j = 0; i < n; i++) {
        if (a[i] >= k - j - 1 && b[i] >= j)
            j++;
        if (j == k)
            return true;
    }
    return false;
}

int main() {
    for (scanf("%lld", &t); t-- > 0; ) {
        scanf("%lld", &n);
        for (ll i = 0; i < n; i++)
            scanf("%lld %lld", a + i, b + i);
        ll p = 0, q = n + 1;
        while (p < q) {
            ll mid = (p + q) >> 1;
            if (!check(mid))
                q = mid;
            else
                p = mid + 1;
        }
        printf("%lld\n", p - 1);
    }
    return 0;
}
