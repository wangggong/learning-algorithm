#include <algorithm>
#include <iostream>
#include <unordered_map>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 3e5 + 1, MOD = 998244353;
ll n, tot = 1, cf = 1, cs = 1, cp = 1, s, t;
pll a[N + 10];

int main() {
    cin >> n;
    unordered_map<ll, ll> Cf, Cs, Cp;
    for (ll i = 1; i <= n; i++) {
        cin >> s >> t;
        Cf[s]++, Cs[t]++, Cp[s * N + t]++, a[i] = {s, t};
        tot = (tot * i) % MOD;
        cf = (cf * Cf[s]) % MOD;
        cs = (cs * Cs[t]) % MOD;
        cp = (cp * Cp[s * N + t]) % MOD;
    }
    sort(a + 1, a + n + 1);
    bool can_be_ordered = true;
    for (ll i = 2; i <= n; i++)
        if (a[i].second < a[i - 1].second)
            can_be_ordered = false;
    tot = ((tot - cf) % MOD + MOD) % MOD;
    tot = ((tot - cs) % MOD + MOD) % MOD;
    if (can_be_ordered)
        tot = (tot + cp) % MOD;
    cout << max(tot, 0ll) << endl;
    return 0;
}
