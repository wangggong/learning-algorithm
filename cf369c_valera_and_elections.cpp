#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;

const ll N = 1e5;
ll n, x, y, t, ans;
vector<pll> G[N + 10];
bool vis[N + 10];

void dfs1(ll u, ll p) {
    for (auto [v, w] : G[u]) {
        if (v == p)
            continue;
        dfs1(v, u);
        if (w == 2)
            vis[v] = true;
    }
}

bool dfs2(ll u, ll p) {
    ll ans = vis[u];
    for (auto [v, w] : G[u]) {
        if (v == p)
            continue;
        ll cur = dfs2(v, u);
        ans |= cur;
        if (cur)
            vis[u] = false;
    }
    return ans;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 1; i < n; i++) {
        scanf("%lld %lld %lld", &x, &y, &t);
        G[x].push_back({y, t});
        G[y].push_back({x, t});
    }
    dfs1(1, -1);
    dfs2(1, -1);
    for (ll i = 1; i <= n; i++)
        if (vis[i])
            ans++;
    printf("%lld\n", ans);
    for (ll i = 1; i <= n; i++)
        if (vis[i])
            printf("%lld ", i);
    return 0;
}
