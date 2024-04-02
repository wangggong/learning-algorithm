#include <iostream>
#include <climits>

using namespace std;
typedef long long ll;
const ll N = 1e5;
ll t, n, a[N + 10], dp[N + 10];

ll get() {
    for (ll i = 0; i < n; i++)
        dp[i] = LLONG_MAX;
    for (ll i = 1; i + 1 < n; i++)
        dp[i] = (i <= 2 ? 0 : min(dp[i - 2], i % 2 ? LLONG_MAX : dp[i - 3]))
            + max(max(a[i - 1], a[i + 1]) + 1 - a[i], 0ll);
    return min(dp[n - 2], n % 2 ? LLONG_MAX : dp[n - 3]);
}

int main() {
    for (scanf("%lld", &t); t--; ) {
        scanf("%lld", &n);
        for (ll i = 0; i < n; i++)
            scanf("%lld", a + i);
        printf("%lld\n", get());
    }
    return 0;
}
