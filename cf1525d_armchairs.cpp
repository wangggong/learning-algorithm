#include <iostream>
#include <vector>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 5000;
ll n, x, dp[N + 10];
vector<ll> a[2];

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", &x), a[x].push_back(i);
    memset(dp, 0x3f, sizeof dp);
    dp[0] = 0;
    ll m = a[1].size();
    for (ll v : a[0])
        for (ll i = m - 1; i >= 0; i--)
            dp[i + 1] = min(dp[i + 1], abs(v - a[1][i]) + dp[i]);
    printf("%lld\n", dp[m]);
    return 0;
}
