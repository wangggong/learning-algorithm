// 参考: https://codeforces.com/contest/1582/submission/165737186
//
// 多少是超纲了 (
#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;

const ll N = 1e5;
ll a[N + 10], dp[2][N + 10], n, k, t;

int main() {
    cin >> t;
    while (t--) {
        memset(a, 0, sizeof a);
        memset(dp, 0, sizeof dp);
        cin >> n;
        for (int i = 0; i < n; i++)
            cin >> a[i];
        for (int i = n - 1; i >= 0; i--)
            dp[0][i] = max(dp[0][i + 1], a[i]), a[i] += a[i + 1];
        k = 2;
        for (ll l = n - 1; l >= k; l -= k - 1, k++) {
            dp[(k - 1) % 2][l-k+1] = 0;
            for (ll i = l - k; i >= 0; i--) {
                ll s = a[i] - a[i + k];
                dp[(k - 1) % 2][i] = dp[(k - 1) % 2][i + 1] < s && s < dp[k % 2][i + k] ? s : dp[(k - 1) % 2][i + 1];
            }
            if (dp[(k - 1) % 2][0] == 0)
                break;
        }
        cout << k - 1 << endl;
    }
    return 0;
}
