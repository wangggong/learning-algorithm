#include <iostream>

using namespace std;
typedef long long ll;

const ll N = 1e3, M = 10, MOD = 998244353;
int n, m, dp[N + 2][M + 2][M + 2][M + 2], ans;

int main() {
    cin >> n >> m;
    dp[0][m][m][m] = 1;
    for (ll i = 1; i <= n; i++)
        for (ll x = 0; x <= m; x++)
            for (ll y = x; y <= m; y++)
                for (ll z = y; z <= m; z++)
                    for (ll j = 0; j < m; j++) {
                        ll cur = dp[i - 1][x][y][z];
                        if (j <= x)
                            dp[i][j][y][z] = (dp[i][j][y][z] + cur) % MOD;
                        else if (j <= y)
                            dp[i][x][j][z] = (dp[i][x][j][z] + cur) % MOD;
                        else if (j <= z)
                            dp[i][x][y][j] = (dp[i][x][y][j] + cur) % MOD;
                        else;
                    }
    for (ll x = 0; x < m; x++)
        for (ll y = x + 1; y < m; y++)
            for (ll z = y + 1; z < m; z++)
                ans = (ans + dp[n][x][y][z]) % MOD;
    cout << ans << endl;
    return 0;
}
