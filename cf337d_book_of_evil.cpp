#include <iostream>
#include <vector>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 1e5;
ll n, m, d, u, v, ans, up[N + 10], fi[N + 10], se[N + 10];
bool M[N + 10];
vector<ll> G[N + 10];

void dfs1(ll u, ll p) {
    if (M[u])
        up[u] = fi[u] = se[u] = 0;
    for (ll v : G[u])
        if (v != p) {
            dfs1(v, u);
            if (fi[v] + 1 > fi[u])
                se[u] = fi[u], fi[u] = fi[v] + 1;
            else if (fi[v] + 1 > se[u])
                se[u] = fi[v] + 1;
        }
}

void dfs2(ll u, ll p) {
    for (ll v : G[u])
        if (v != p) {
            up[v] = max(up[u], fi[v] + 1 == fi[u] ? se[u] : fi[u]) + 1;
            dfs2(v, u);
        }
}

int main() {
    scanf("%lld %lld %lld", &n, &m, &d);
    // 这里需要先把最大距离设置为极小的负值, 否则会爆.
    memset(up, 0xc0, sizeof up);
    memset(fi, 0xc0, sizeof fi);
    memset(se, 0xc0, sizeof se);
    for (ll i = 0; i < m; i++)
        scanf("%lld", &u), M[u] = true;
    for (ll i = 1; i < n; i++)
        scanf("%lld %lld", &u, &v), G[u].push_back(v), G[v].push_back(u);
    dfs1(1, -1);
    dfs2(1, -1);
    for (ll i = 1; i <= n; i++)
        if (up[i] <= d && fi[i] <= d)
            ans++;
    printf("%lld\n", ans);
    return 0;
}
