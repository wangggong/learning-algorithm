#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 500, M = 1000;
ll n, m, a[N + 10], b[M + 10], c[N + 10], ans;
bool vis[N + 10];

int main() {
    scanf("%lld%lld", &n, &m);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    for (ll i = 0; i < m; i++) {
        scanf("%lld", b + i), b[i]--;
        memset(vis, false, sizeof vis);
        for (ll j = i - 1; j >= 0 && b[j] != b[i]; j--)
            vis[b[j]] = true;
        for (ll j = 0; j < n; j++)
            if (vis[j])
                c[j]++;
    }
    for (ll i = 0; i < n; i++)
        ans += a[i] * c[i];
    printf("%lld", ans);
    return 0;
}
