#include <iostream>
#define __FAST_READ__ std::ios::sync_with_stdio(false), std::cin.tie(nullptr), std::cout.tie(nullptr);

using namespace std;
typedef long long ll;

const ll N = 2e5;
ll a[2][N + 10], n, c, dp[2][2];

int main() {
    __FAST_READ__;
    cin >> n >> c;
    for (int j = 0; j < 2; j++)
        for (int i = 2; i <= n; i++)
            cin >> a[j][i];
    dp[1][1] = INT_MAX;
    cout << min(dp[1][0], dp[1][1]) << " ";
    for (int i = 2; i <= n; i++)
        dp[i % 2][0] = min(dp[(i - 1) % 2][0], dp[(i - 1) % 2][1]) + a[0][i],
        dp[i % 2][1] = min(dp[(i - 1) % 2][0] + c, dp[(i - 1) % 2][1]) + a[1][i],
        cout << min(dp[i % 2][0], dp[i % 2][1]) << " ";
    return 0;
}
