#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 3e5;
ll n, k, b[N + 10], a, d, dd[N + 10], ans;

int main() {
    scanf("%lld %lld", &n, &k);
    for (ll i = 0; i < n; i++)
        scanf("%lld", b + i);
    for (ll i = n - 1; i >= 0; i--) {
        d += dd[i];
        a += d;
        if (a < b[i]) {
            ll kk = min(i + 1, k);
            ll cur = (b[i] - a - 1) / kk + 1;
            ans += cur;
            a += cur * kk;
            if (i) {
                dd[i - 1] -= cur;
                if (i > kk)
                    dd[i - kk - 1] += cur;
            }
        }
    }
    printf("%lld\n", ans);
    return 0;
}
