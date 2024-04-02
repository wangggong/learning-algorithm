#include <iostream>
#define __FAST__ std::ios::sync_with_stdio(false), std::cin.tie(nullptr), std::cout.tie(nullptr)

using namespace std;
typedef long long ll;

const ll N = 2e5, MOD = 998244353;
string s, t;
ll ps[N + 10], n, m, ans;

ll pow(ll n, ll k) {
    ll ans = 1;
    for (; k; k >>= 1) {
        if (k & 1)
            ans = ans * n % MOD;
        n = n * n % MOD;
    }
    return ans;
}

int main() {
    __FAST__;
    cin >> n >> m >> s >> t;
    for (ll i = 0; i < m; i++)
        ps[i + 1] = ps[i] + (t[i] - '0');
    for (ll i = 0; i < n; i++)
        if (s[i] == '1' && m - (n - 1 - i) >= 0)
            ans = (ans + pow(2, n - 1 - i) * ps[m - (n - 1 - i)] % MOD) % MOD;
    cout << ans << endl;
    return 0;
}
