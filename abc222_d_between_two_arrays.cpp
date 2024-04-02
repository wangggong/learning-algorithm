#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 3000;
const ll MOD = 998244353;
ll n, a[N + 10], b[N + 10], dp[2][N + 10];

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    for (ll i = 0; i < n; i++)
        scanf("%lld", b + i);
    for (ll j = 0; j <= N; j++)
        if (a[0] <= j && j <= b[0])
            dp[0][j] = 1;
    for (ll i = 1; i < n; i++) {
        memset(dp[i % 2], 0, sizeof dp[i % 2]);
        ll tot = 0;
        for (ll j = 0; j <= N; j++) {
            tot = (tot + dp[(i - 1) % 2][j]) % MOD;
            if (a[i] <= j && j <= b[i])
                dp[i % 2][j] = tot;
        }
    }
    ll tot = 0;
    for (ll j = 0; j <= N; j++)
        tot = (tot + dp[(n - 1) % 2][j]) % MOD;
    printf("%lld", tot);
    return 0;
}
