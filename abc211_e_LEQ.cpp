#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;
typedef long long ll;

const ll N = 3e5;
const ll MOD = 998244353;
const ll MOD2 = (MOD + 1) >> 1;
ll n, k, ans, tr[N + 10];

ll lowbit(ll x) { return x & (-x); }

int main() {
    vector<ll> a, b;
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", &k), a.push_back(k), b.push_back(k);
    sort(b.begin(), b.end());
    ll p2 = 1, ip2 = 1;
    for (ll i = 0; i < n; i++) {
        ll k = lower_bound(b.begin(), b.end(), a[i]) - b.begin() + 1;
        for (ll v = k; v; v -= lowbit(v))
            ans = (ans + tr[v] * p2) % MOD;
        p2 = p2 * 2 % MOD, ip2 = ip2 * MOD2 % MOD;
        for (ll v = k; v <= n; v += lowbit(v))
            tr[v] = (tr[v] + ip2) % MOD;
    }
    printf("%lld\n", ans);
    return 0;
}
