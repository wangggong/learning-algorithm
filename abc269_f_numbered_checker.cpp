#include <iostream>

using namespace std;
typedef long long ll;
const ll MOD = 998244353;
// const ll MOD2 = (MOD + 1) / 2;
ll n, m, q, r1, r2, c1, c2, ans;

ll f(ll r, ll c) {
    // odd
    ll rs = (r + 1) / 2;
    ll cs = (c + 1) / 2;
    ll line1 = cs * cs % MOD;
    ll Sodd = rs * ((line1 + (rs - 1) * cs % MOD * m % MOD) % MOD) % MOD;
    // even
    rs = r / 2;
    cs = c / 2;
    ll line2 = cs * ((2 + m % MOD + cs - 1) % MOD) % MOD;
    ll Seven = rs * ((line2 + (rs - 1) * cs % MOD * m % MOD) % MOD) % MOD;
    return (Sodd + Seven) % MOD;
}

int main() {
    scanf("%lld %lld %lld", &n, &m, &q);
    while (q--) {
        scanf("%lld %lld %lld %lld", &r1, &r2, &c1, &c2);
        ans = (((f(r2, c2) - f(r1 - 1, c2) + MOD) % MOD - f(r2, c1 - 1) + MOD) % MOD + f(r1 - 1, c1 - 1)) % MOD;
        printf("%lld\n", ans);
    }
    return 0;
}
