// 参考: https://atcoder.jp/contests/arc136/submissions/45039336
//
// 十进制下的 SOS DP
#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 1e6;
ll n, a[N + 10], f[N + 10], ans;

bool carry(ll n) {
    for (; n; n /= 10)
        if ((n % 10) * 2 >= 10)
            return true;
    return false;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++) {
        scanf("%lld", a + i);
        f[a[i]]++;
    }
    for (ll p = 1; p < N; p *= 10)
        for (ll i = 0; i < N; i++)
            if ((i / p) % 10)
                f[i] += f[i - p];
    for (ll i = 0; i < n; i++)
        ans += f[N - 1 - a[i]] - (carry(a[i]) ? 0 : 1);
    printf("%lld", ans >> 1);
    return 0;
}
