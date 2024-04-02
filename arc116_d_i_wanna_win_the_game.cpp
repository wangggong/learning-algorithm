#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 5000, MOD = 998244353;
ll n, m, dp[N + 10];
ll F[N + 10], IF[N + 10];

ll POW(ll n, ll k) {
    ll ans = 1;
    for (; k; k >>= 1) {
        if (k & 1)
            ans = (ans * n) % MOD;
        n = (n * n) % MOD;
    }
    return ans;
}

void init() {
    F[0] = 1;
    for (ll i = 1; i <= N; i++)
        F[i] = (F[i - 1] * i) % MOD;
    IF[N] = POW(F[N], MOD - 2);
    for (ll i = N - 1; i >= 0; i--)
        IF[i] = (IF[i + 1] * (i + 1)) % MOD;
    return;
}

ll C(ll n, ll m) { return F[n] * IF[m] % MOD * IF[n - m] % MOD; }

int main() {
    init();
    scanf("%lld %lld", &n, &m);
    dp[0] = 1;
    for (ll i = 2; i <= m; i += 2)
        for (ll j = 0; j <= i && j <= n; j += 2)
            dp[i] = (dp[i] + dp[(i - j) / 2] * C(n, j)) % MOD;
    printf("%lld\n", dp[m]);
    return 0;
}
