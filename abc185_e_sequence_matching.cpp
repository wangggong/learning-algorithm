#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;

int main() {
    ll n, m;
    scanf("%lld %lld", &n, &m);
    vector<ll> a(n + 1), b(m + 1);
    vector<vector<ll>> dp(n + 1, vector<ll>(m + 1));
    for (ll i = 1; i <= n; i++)
        scanf("%lld", &a[i]), dp[i][0] = i;
    for (ll i = 1; i <= m; i++)
        scanf("%lld", &b[i]), dp[0][i] = i;
    for (ll i = 1; i <= n; i++)
        for (ll j = 1; j <= m; j++)
            dp[i][j] = min(min(dp[i - 1][j], dp[i][j - 1]) + 1, dp[i - 1][j - 1] + (a[i] == b[j] ? 0 : 1));
    printf("%lld\n", dp[n][m]);
    return 0;
}
