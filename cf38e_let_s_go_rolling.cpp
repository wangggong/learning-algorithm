#include <iostream>
#include <algorithm>
#include <vector>
#include <climits>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(nullptr), cout.tie(nullptr)

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
ll n, x, c, ans = LLONG_MAX;
vector<pll> vec;

int main() {
    __FAST__;
    cin >> n;
    for (ll i = 0; i < n; i++) {
        cin >> x >> c;
        vec.push_back({x, c});
    }
    sort(vec.begin(), vec.end());
    vector<vector<ll>> dp(n, vector<ll>(n, 0));
    dp[0][0] = vec[0].second;
    for (ll i = 1; i < n; i++) {
        dp[i][i] = LLONG_MAX;
        for (ll j = 0; j < i; j++)
            dp[i][j] = dp[i - 1][j] + vec[i].first - vec[j].first;
        for (ll j = 0; j < i; j++)
            dp[i][i] = min(dp[i][i], dp[i - 1][j]);
        dp[i][i] += vec[i].second;
    }
    for (ll i = 0; i < n; i++)
        ans = min(ans, dp[n - 1][i]);
    cout << ans << endl;
    return 0;
}
