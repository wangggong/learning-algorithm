#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 1e3;
ll t, n, m, a[N + 10][N + 10], S[N + 10][N + 10];

int main() {
    scanf("%lld", &t);
    while (t--) {
        scanf("%lld %lld", &n, &m);
        for (ll i = 0; i < n; i++)
            for (ll j = 0; j < m; j++)
                scanf("%lld", &a[i][j]);
        if ((m + n) % 2 == 0) {
            puts("NO");
            continue;
        }
        for (ll i = 0; i < n; i++)
            for (ll j = 0; j < m; j++)
                S[i][j] = (i == 0 && j == 0
                    ? 0
                    : (i == 0 || j == 0
                        ? (i == 0 ? S[i][j - 1] : S[i - 1][j])
                        : min(S[i - 1][j], S[i][j - 1]))) + a[i][j];
        ll minVal = S[n - 1][m - 1];
        for (ll i = 0; i < n; i++)
            for (ll j = 0; j < m; j++)
                S[i][j] = (i == 0 && j == 0
                    ? 0
                    : (i == 0 || j == 0
                        ? (i == 0 ? S[i][j - 1] : S[i - 1][j])
                        : max(S[i - 1][j], S[i][j - 1]))) + a[i][j];
        ll maxVal = S[n - 1][m - 1];
        puts(minVal <= 0 && maxVal >= 0 ? "YES" : "NO");
    }
    return 0;
}
