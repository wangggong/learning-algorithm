#include <iostream>
#include <vector>
#include <cstring>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 2e5;
ll n, u, v, q, k, root = 1, mx, ans[N + 10], path[N + 10];
vector<ll> G[N + 10];
vector<pll> queries[N + 10];

void dfs(ll u, ll p, ll d) {
    if (mx < d)
        mx = d, root = u;
    path[d] = u;
    for (auto q : queries[u])
        if (q.first <= d)
            ans[q.second] = path[d - q.first];
    for (auto v : G[u])
        if (v != p)
            dfs(v, u, d + 1);
    return;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 1; i < n; i++)
        scanf("%lld %lld", &u, &v), G[u].push_back(v), G[v].push_back(u);
    scanf("%lld", &q);
    for (ll i = 0; i < q; i++)
        scanf("%lld %lld", &u, &k), queries[u].push_back({k, i});
    memset(ans, -1, sizeof ans);
    for (ll i = 0; i < 3; i++)
        mx = -1, dfs(root, -1, 0);
    for (ll i = 0; i < q; i++)
        printf("%lld\n", ans[i]);
    return 0;
}
