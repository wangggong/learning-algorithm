#include <iostream>

using namespace std;

typedef long long ll;
const ll N = 2e3, MOD = 1e9 + 7;

ll dp[N + 10][N + 10], n, m, a[N + 10], b[N + 10];

int main() {
    scanf("%lld %lld", &n, &m);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    for (ll i = 0; i < m; i++)
        scanf("%lld", b + i);
    for (ll i = 0; i <= n; i++)
        dp[i][0] = 1;
    for (ll i = 0; i <= m; i++)
        dp[0][i] = 1;
    for (ll i = 1; i <= n; i++)
        for (ll j = 1; j <= m; j++)
            dp[i][j] = ((dp[i - 1][j] + dp[i][j - 1]) % MOD + MOD - (a[i - 1] == b[j - 1] ? 0 : dp[i - 1][j - 1])) % MOD;
    cout << dp[n][m] << endl;
    return 0;
}
