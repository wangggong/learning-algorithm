// 参考: https://atcoder.jp/contests/dwacon5th-prelims/submissions/46029217

#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 1000;
const ll D = 40;
ll n, k, a[N + 10], ans;

int main() {
    scanf("%lld %lld", &n, &k);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    for (ll d = D; d >= 0; d--) {
        ll cnt = 0;
        for (ll i = 0; i < n; i++) {
            ll S = 0;
            for (ll j = i; j < n; j++) {
                S += a[j];
                if ((S & (1ll << d)) && ((S & ans) == ans))
                    cnt++;
            }
        }
        if (cnt >= k)
            ans |= 1ll << d;
    }
    printf("%lld", ans);
    return 0;
}
