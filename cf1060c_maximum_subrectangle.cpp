#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;

const ll N = 2000;
ll sa[N + 10], sb[N + 10], ma[N + 10], mb[N + 10], n, m, k, target, ans;

// 纯傻逼, 不需要往二维前缀和上推的, 非要推到二维上...
int main() {
    cin >> n >> m;
    for (ll i = 0; i < n; i++)
        cin >> k, sa[i + 1] = sa[i] + k;
    for (ll i = 0; i < m; i++)
        cin >> k, sb[i + 1] = sb[i] + k;
    cin >> target;
    memset(ma, 0x3f, sizeof ma); memset(mb, 0x3f, sizeof mb);
    for (ll i = 1; i <= n; i++)
        for (ll j = i; j <= n; j++)
            ma[i] = min(ma[i], sa[j] - sa[j - i]);
    for (ll i = 1; i <= m; i++)
        for (ll j = i; j <= m; j++)
            mb[i] = min(mb[i], sb[j] - sb[j - i]);
    for (ll i = 1; i <= n; i++)
        for (ll j = 1; j <= m; j++)
            if (ma[i] * mb[j] <= target)
                ans = max(ans, i * j);
    cout << ans << endl;
    return 0;
}
