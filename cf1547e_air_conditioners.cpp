// 参考: https://codeforces.com/contest/1547/submission/250445406
#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 3e5;
ll t, n, k, a[N + 10], x, dp[N + 10];

int main() {
    scanf("%lld", &t); while (t--) {
        memset(dp, 0x3f, sizeof dp);
        scanf("%lld%lld", &n, &k);
        for (ll i = 0; i < k; i++)
            scanf("%lld", a + i);
        for (ll i = 0; i < k; i++)
            scanf("%lld", &x), dp[a[i]] = x;
        for (ll i = 1; i <= n; i++)
            dp[i] = min(dp[i], dp[i - 1] + 1);
        for (ll i = n; i > 0; i--)
            dp[i] = min(dp[i], dp[i + 1] + 1);
        for (ll i = 1; i <= n; i++)
            printf("%lld ", dp[i]);
        puts("");
    }
    return 0;
}
