#include <iostream>

using namespace std;
typedef long long ll;

const int N = 1e5;
ll a[N + 10], dp[N + 10][2], n;

int main() {
    scanf("%lld", &n);
    for (int i = 1; i <= n; i++)
        scanf("%lld", a + i);
    ll ans = LLONG_MIN;
    dp[2][1] = abs(a[1] - a[2]);
    ans = max(ans, dp[2][1]);
    for (int i = 3; i <= n; i++)
        dp[i][(i - 1) % 2] = abs(a[i - 1] - a[i]),
        dp[i][0] = max(dp[i][0], dp[i - 1][0] + abs(a[i - 1] - a[i]) * (i % 2 ? 1 : -1)),
        dp[i][1] = max(dp[i][1], dp[i - 1][1] + abs(a[i - 1] - a[i]) * (i % 2 ? -1 : 1)),
        ans = max(ans, max(dp[i][0], dp[i][1]));
    printf("%lld\n", ans);
    return 0;
}
