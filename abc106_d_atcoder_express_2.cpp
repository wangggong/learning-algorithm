#include <iostream>

typedef long long ll;
const ll N = 500;
ll n, m, q, l, r, dp[N + 10][N + 10], a[N + 10][N + 10];

int main() {
    scanf("%lld %lld %lld", &n, &m, &q);
    while (m--) {
        scanf("%lld %lld", &l, &r);
        a[l][r]++;
    }
    // 二维前缀和那种玩法了属于是...
    for (ll j = 1; j <= n; j++)
        for (ll i = j; i >= 1; i--)
            dp[i][j] = dp[i + 1][j] + dp[i][j - 1] - dp[i + 1][j - 1] + a[i][j];
    while (q--) {
        scanf("%lld %lld", &l, &r);
        printf("%lld\n", dp[l][r]);
    }
    return 0;
}
