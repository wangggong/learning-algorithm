/*
 * #include <iostream>
 * #include <cstring>
 * 
 * using namespace std;
 * typedef long long ll;
 * const ll N = 2e5;
 * ll t, n, dp[N + 10][2];
 * char a[N + 10];
 * 
 * int main() {
 *     for (scanf("%lld", &t); t; t--) {
 *         scanf("%lld\n%s", &n, a);
 *         memset(dp, 0x3f, sizeof dp);
 *         ll c = 0;
 *         if (a[0] == a[1])
 *             dp[1][a[0] - '0'] = 1;
 *         else
 *             dp[1][0] = dp[1][1] = 1, c++;
 *         for (ll i = 1; (i << 1) < n; i++)
 *             if (a[i << 1] == a[(i << 1) | 1]) {
 *                 ll k = a[i << 1] - '0';
 *                 dp[i + 1][k] = min(dp[i][k], dp[i][1 - k] + 1);
 *             } else {
 *                 c++;
 *                 dp[i + 1][0] = min(dp[i][0], dp[i][1] + 1);
 *                 dp[i + 1][1] = min(dp[i][1], dp[i][0] + 1);
 *             }
 *         printf("%lld %lld\n", c, min(dp[n >> 1][0], dp[n >> 1][1]));
 *     }
 *     return 0;
 * }
 */

// 参考: https://codeforces.com/contest/1678/submission/227695878
#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll t, n;
char a[N + 10];

int main() {
    for (scanf("%lld", &t); t; t--) {
        scanf("%lld\n%s", &n, a);
        ll c = 0, seg = 0;
        char pre = 0;
        for (ll i = 0; i < n; i += 2)
            if (a[i] != a[i + 1])
                c++;
            else if (a[i] != pre)
                seg++, pre = a[i];
        printf("%lld %lld\n", c, max(seg, 1ll));
    }
    return 0;
}
