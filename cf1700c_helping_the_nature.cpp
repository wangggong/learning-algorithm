#include <iostream>

using namespace std;
typedef long long ll;
ll n, t;

int main() {
    scanf("%lld", &t);
    while (t--) {
        scanf("%lld", &n);
        ll ans = 0, a0, a, last;
        for (ll i = 0; i < n; i++, last = a) {
            scanf("%lld", &a);
            if (i == 0) {
                a0 = a;
                continue;
            }
            ll d = a - last;
            ans += abs(d);
            if (d < 0)
                a0 += d;
        }
        ans += abs(a0);
        printf("%lld\n", ans);
    }
    return 0;
}
