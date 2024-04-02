#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;
ll n, m, k, v;

int main() {
    scanf("%lld %lld", &n, &m);
    vector<ll> dp(m + 1);
    while (n--) {
        scanf("%lld", &k);
        vector<ll> S(k + 1), mx(k + 1);
        for (ll i = 0; i < k; i++)
            scanf("%lld", &v), S[i + 1] = S[i] + v;
        for (ll i = 1; i <= k; i++)
            for (ll j = 0; j <= i; j++)
                mx[i] = max(mx[i], S[j] + S[k] - S[k - i + j]);
        for (ll i = m; i >= 0; i--)
            for (ll j = k; j > 0; j--)
                if (i + j <= m)
                    dp[i + j] = max(dp[i + j], dp[i] + mx[j]);
    }
    printf("%lld\n", dp[m]);
    return 0;
}
