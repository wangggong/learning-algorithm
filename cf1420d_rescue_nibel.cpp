#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;
typedef long long ll;

const ll N = 3e5, MOD = 998244353;
ll n, k, l, r, s, c, ans, F[N + 10], IF[N + 10];

ll C(ll n, ll m) {
    return n < m || m < 0 ? 0 : F[n] * IF[m] % MOD * IF[n - m] % MOD;
}

ll pow(ll n, ll k) {
    ll ans = 1;
    for (; k; k >>= 1) {
        if (k & 1)
            ans = ans * n % MOD;
        n = n * n % MOD;
    }
    return ans;
}

void init() {
    F[0] = 1;
    for (ll i = 1; i <= N; i++)
        F[i] = F[i - 1] * i % MOD;
    IF[N] = pow(F[N], MOD - 2);
    for (ll i = N - 1; i >= 0; i--)
        IF[i] = IF[i + 1] * (i + 1) % MOD;
    return;
}

int main() {
    init();
    vector<ll> a;
    scanf("%lld %lld", &n, &k);
    for (ll i = 0; i < n; i++)
        scanf("%lld %lld", &l, &r), a.push_back(l << 1 | 1), a.push_back((r + 1) << 1);
    sort(a.begin(), a.end());
    for (ll i = 0; i < 2 * n; i++) {
        s += a[i] & 1 ? 1 : -1;
        if (a[i] & 1) {
            c++;
            if (!(a[i + 1] & 1))
                ans = (ans + C(s, k) - C(s - c, k) + MOD) % MOD, c = 0;
        }
    }
    printf("%lld\n", ans);
    return 0;
}
