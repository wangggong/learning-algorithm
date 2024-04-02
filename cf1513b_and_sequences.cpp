#include <iostream>
#include <climits>

using namespace std;
typedef long long ll;
const ll N = 2e5;
const ll MOD = 1e9 + 7;
ll t, n, a[N + 10], F[N + 10];

int main() {
    F[0] = 1;
    for (ll i = 1; i <= N; i++)
        F[i] = F[i - 1] * i % MOD;
    scanf("%lld", &t); while (t--) {
        ll x = -1, _min = LLONG_MAX, _c = 0;
        scanf("%lld", &n); for (ll i = 0; i < n; i++) {
            scanf("%lld", a + i);
            x &= a[i];
            if (a[i] < _min)
                _min = a[i], _c = 1;
            else if (a[i] == _min)
                _c++;
        }
        ll ans = 0;
        if (_min != x || _c < 2)
            printf("%lld\n", ans);
        else
            printf("%lld\n", _c * (_c - 1) % MOD * F[n - 2] % MOD);
    }
    return 0;
}
