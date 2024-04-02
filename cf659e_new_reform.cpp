#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;
const ll N = 1e5;
ll n, m, u, v, ans;
vector<ll> G[N + 10];
bool is_tree, vis[N + 10];

void dfs(ll u, ll p) {
    if (vis[u]) {
        is_tree = false;
        return;
    }
    vis[u] = true;
    for (ll v : G[u])
        if (v != p)
            dfs(v, u);
}

int main() {
    scanf("%lld%lld", &n, &m);
    for (ll i = 0; i < m; i++)
        scanf("%lld%lld", &u, &v), G[u].push_back(v), G[v].push_back(u);
    for (ll i = 1; i <= n; i++)
        if (!vis[i]) {
            is_tree = true;
            dfs(i, -1);
            ans += is_tree;
        }
    printf("%lld", ans);
    return 0;
}
