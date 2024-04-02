#include <iostream>
#include <algorithm>
 
using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 1e5;
ll n, a, b, ans, dp[N + 10];
pll pos[N + 10];
 
int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld %lld", &a, &b), pos[i] = {a, b};
    sort(pos, pos + n);
    for (ll i = 0; i < n; i++) {
        ll p = 0, q = i, k = pos[i].first - pos[i].second;
        while (p < q) {
            ll mid = (p + q) >> 1;
            if (pos[mid].first >= k)
                q = mid;
            else
                p = mid + 1;
        }
        dp[i] = (p > 0 ? dp[p - 1] : 0) + 1;
        ans = max(ans, dp[i]);
    }
    printf("%lld", n - ans);
    return 0;
}
