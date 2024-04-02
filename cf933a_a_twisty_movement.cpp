#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 2000;
ll a[N + 10], n, dp[4 + 1];

int main() {
    cin >> n;
    for (ll i = 0; i < n; i++) {
        cin >> a[i];
        if (a[i] == 1)
            dp[1]++, dp[2] = max(dp[1], dp[2]), dp[3] = max(dp[2], dp[3] + 1), dp[4] = max(dp[3], dp[4]);
        else
            dp[2] = max(dp[1], dp[2] + 1), dp[3] = max(dp[2], dp[3]), dp[4] = max(dp[3], dp[4] + 1);
    }
    cout << dp[4] << endl;
    return 0;
}
