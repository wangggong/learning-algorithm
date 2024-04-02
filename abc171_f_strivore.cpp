#include <iostream>
#include <unordered_map>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)

using namespace std;
typedef long long ll;

const ll N = 2e6, MOD = 1e9 + 7;

ll k, F[N + 1], IF[N + 1];
string s;

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

ll C(ll n, ll m) { return ((F[n] * IF[m]) % MOD * IF[n - m]) % MOD; }

int main() {
    __FAST__;
    init();
    cin >> k >> s;
    ll n = s.size(), ans = 0, p26 = 1, p25 = POW(25, k), inv25 = POW(25, MOD - 2);
    for (ll i = 0; i <= k; i++)
        ans = (ans + p26 * p25 % MOD * C(n - 1 + k - i, n - 1) % MOD) % MOD,
        p26 = (p26 * 26) % MOD,
        p25 = (p25 * inv25) % MOD;
    cout << ans << endl;
    return 0;
}
