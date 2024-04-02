#include <iostream>
#include <climits>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)

using namespace std;
typedef long long ll;

const ll N = 3e5;
ll n, a[N + 10], dp[2][2], ans;

// 打 家 劫 舍
int main() {
    __FAST__;
    cin >> n;
    for (ll i = 0; i < n; i++)
        cin >> a[i];

    dp[0][0] = INT_MAX, dp[0][1] = a[0];
    for (ll i = 1; i < n; i++)
        dp[i % 2][0] = dp[(i - 1) % 2][1], dp[i % 2][1] = min(dp[(i - 1) % 2][0], dp[(i - 1) % 2][1]) + a[i];
    ans = dp[(n - 1) % 2][0];

    dp[0][0] = a[n - 1], dp[0][1] = a[n - 1] + a[0];
    for (ll i = 1; i < n; i++)
        dp[i % 2][0] = dp[(i - 1) % 2][1], dp[i % 2][1] = min(dp[(i - 1) % 2][0], dp[(i - 1) % 2][1]) + a[i];
    ans = min(ans, min(dp[(n - 2) % 2][0], dp[(n - 2) % 2][1]));

    cout << ans << endl;
    return 0;
}
