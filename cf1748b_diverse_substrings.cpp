// 参考: https://codeforces.com/problemset/submission/1748/242085912
#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 1e5, D = 10, M = 100;
ll t, n, c[D + 10];
char s[N + 10];

bool valid() {
    ll k = 0;
    for (ll i = 0; i < D; i++)
        k += c[i] > 0;
    for (ll i = 0; i < D; i++)
        if (c[i] > k)
            return false;
    return true;
}

int main() {
    for (scanf("%lld", &t); t-- > 0; ) {
        scanf("%lld%s", &n, s);
        ll ans = 0;
        for (ll l = 1; l <= M && l <= n; l++) {
            memset(c, 0, sizeof c);
            for (ll i = 0; i < l; i++)
                c[s[i] - '0']++;
            ans += valid();
            for (ll i = l; i < n; i++) {
                c[s[i - l] - '0']--, c[s[i] - '0']++, ans += valid();
            }
        }
        printf("%lld\n", ans);
    }
    return 0;
}
