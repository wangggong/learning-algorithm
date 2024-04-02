#include <iostream>
#include <vector>
#define __FAST_READ__ std::ios::sync_with_stdio(false), std::cin.tie(nullptr), std::cout.tie(nullptr);

using namespace std;
typedef long long ll;

const ll N = 3e5;
ll a[N + 10], inv[N + 10], n, m, ans, s, t, x;
vector<ll> P[N + 10];
bool vis[N + 10];

/*
 * int main() {
 *     __FAST_READ__
 *     cin >> n >> m;
 *     for (ll i = 0; i < n; i++)
 *         cin >> a[i], inv[a[i]] = i;
 *     for (ll i = 0; i < m; i++)
 *         cin >> s >> t, P[s].push_back(t), P[t].push_back(s);
 *     for (ll p = 0, q = 0; p < n; ) {
 *         ll index = -1;
 *         for (; q < n; q++) {
 *             vis[a[q]] = true;
 *             for (ll v : P[a[q]])
 *                 if (vis[v]) {
 *                     index = max(index, inv[v]);
 *                 }
 *             if (index >= 0)
 *                 break;
 *         }
 *         if (index >= p)
 *             for (; p <= index; p++)
 *                 ans += q - p, vis[a[p]] = false;
 *         else
 *             ans += q - p, vis[a[p]] = false, p++;
 *         // cout << p << " " << q << " " << index << " " << ans << endl;
 *     }
 *     cout << ans << endl;
 *     return 0;
 * }
 */

int main() {
    __FAST_READ__;
    cin >> n >> m;
    for (ll i = 1; i <= n; i++)
        cin >> x, inv[x] = i;
    for (ll i = 0; i < m; i++) {
        cin >> s >> t;
        ll p = inv[s], q = inv[t];
        if (p > q)
            swap(p, q);
        a[q] = max(a[q], p);
    }
    for (ll i = 1, mx = 0; i <= n; i++)
        mx = max(mx, a[i]), ans += i - mx;
    cout << ans << endl;
    return 0;
}
