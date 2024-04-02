#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;

const ll N = 5000;
ll a[N + 10], ps[N + 10], dp[2][N + 10], n, m, k;

int main() {
    cin >> n >> m >> k;
    for (int i = 0; i < n; i++)
        cin >> a[i], ps[i + 1] = ps[i] + a[i];
    for (int i = 1; i <= k; i++) {
        memset(dp[i % 2], 0, sizeof dp[i % 2]);
        for (int j = m; j <= n; j++)
            dp[i % 2][j] = max(dp[i % 2][j - 1], dp[(i - 1) % 2][j - m] + (ps[j] - ps[j - m]));
    }
    cout << dp[k % 2][n] << endl;
    return 0;
}
