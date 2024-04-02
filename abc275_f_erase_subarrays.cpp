#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 3000;
ll n, m, a, dp[2][N + 10][2];

int main() {
    memset(dp, 0x3f, sizeof dp);
    scanf("%lld %lld", &n, &m);
    dp[0][0][1] = 0;
    for (ll i = 1; i <= n; i++) {
        scanf("%lld", &a);
        dp[i % 2][0][0] = 1, dp[i % 2][0][1] = 0x3f3f3f3f;
        for (ll j = 1; j <= m; j++) {
            dp[i % 2][j][0] = min(dp[(i - 1) % 2][j][0], dp[(i - 1) % 2][j][1] + 1);
            dp[i % 2][j][1] = 0x3f3f3f3f;
            if (j >= a)
                dp[i % 2][j][1] = min(dp[(i - 1) % 2][j - a][0], dp[(i - 1) % 2][j - a][1]);
        }
    }
    for (ll i = 1; i <= m; i++) {
        ll v = min(dp[n % 2][i][0], dp[n % 2][i][1]);
        printf("%lld\n", v > 1e9 ? -1 : v);
    }
    return 0;
}
