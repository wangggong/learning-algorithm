#include <iostream>
#include <cstring>

typedef long long ll;
const ll N = 800, MAXN = 1e9;
ll n, k, a[N + 10][N + 10], S[N + 10][N + 10];

bool check(ll m) {
    memset(S, 0, sizeof S);
    for (ll i = 1; i <= n; i++)
        for (ll j = 1; j <= n; j++) {
            S[i][j] = S[i - 1][j] + S[i][j - 1] - S[i - 1][j - 1] + (a[i - 1][j - 1] <= m ? 1 : 0);
            if (i >= k && j >= k && S[i][j] - S[i - k][j] - S[i][j - k] + S[i - k][j - k] >= (k * k + 1) / 2)
                return true;
        }
    return false;
}

int main() {
    scanf("%lld %lld", &n, &k);
    for (ll i = 0; i < n; i++)
        for (ll j = 0; j < n; j++)
            scanf("%lld", &a[i][j]);
    ll p = 0, q = MAXN;
    while (p < q) {
        ll mid = p + q >> 1;
        if (check(mid))
            q = mid;
        else
            p = mid + 1;
    }
    printf("%lld\n", p);
    return 0;
}
