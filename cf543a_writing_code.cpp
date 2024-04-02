#include <iostream>
#include <unordered_map>

using namespace std;
typedef long long ll;

const ll N = 500;
ll n, m, b, mod, a[N + 10], ans;
ll dp[N + 10][N + 10];

int main() {
    cin >> n >> m >> b >> mod;
    for (ll i = 1; i <= n; i++)
        cin >> a[i];
    dp[0][0] = 1;
    for (ll i = 1; i <= n; i++)
        for (ll j = 1; j <= m; j++)
            for (ll k = a[i]; k <= b; k++)
                dp[j][k] = (dp[j][k] + dp[j - 1][k - a[i]]) % mod;
    for (ll i = 0; i <= b; i++)
        ans = (ans + dp[m][i]) % mod;
    cout << ans << endl;
    return 0;
}
