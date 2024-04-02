#include <iostream>
#include <cstring>
#include <queue>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)
#define DEBUG true

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;

const ll N = 1e5, D = 17;
ll n, m, mi, a[N + 10], ans = N * D, cnt[N + 10], op[N + 10];
bool vis[N + 10];

int main() {
    __FAST__;
    cin >> n;
    for (ll i = 0; i < n; i++) {
        cin >> a[i];
        if (a[i] > a[mi])
            mi = i;
    }
    for (ll i = 0; i < n; i++) {
        memset(vis, false, sizeof vis);
        queue<pll> Q;
        Q.push({a[i], 0});
        vis[a[i]] = true;
        while (!Q.empty()) {
            auto q = Q.front(); Q.pop();
            cnt[q.first]++;
            op[q.first] += q.second;
            ll v = q.first << 1;
            if (!vis[v] && v <= N) {
                Q.push({v, q.second + 1});
                vis[v] = true;
            }
            v = q.first >> 1;
            if (!vis[v] && v <= N) {
                Q.push({v, q.second + 1});
                vis[v] = true;
            }
        }
    }
    for (ll i = 0; i <= N; i++)
        if (cnt[i] == n)
            ans = min(ans, op[i]);
    cout << ans << endl;
    return 0;
}
