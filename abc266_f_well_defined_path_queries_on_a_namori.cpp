#include <iostream>
#include <vector>
#include <queue>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll n, q, u, v, deg[N + 10], ID[N + 10];
vector<ll> G[N + 10];

void dfs(ll k, ll v, ll u) {
    ID[v] = k;
    for (auto w : G[v])
        if (w != u && deg[w] < 2)
            dfs(k, w, v);
    return;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 1; i <= n; i++)
        scanf("%lld %lld", &u, &v), G[u].push_back(v), G[v].push_back(u), deg[u]++, deg[v]++;
    queue<ll> Q;
    for (ll i = 1; i <= n; i++)
        if (deg[i] == 1)
            Q.push(i);
    while (!Q.empty()) {
        auto q = Q.front(); Q.pop();
        for (auto v : G[q])
            if (--deg[v] == 1)
                Q.push(v);
    }
    for (ll i = 1, k = 1; i <= n; i++)
        if (deg[i] > 1)
            dfs(k++, i, -1);
    scanf("%lld", &q);
    while (q--)
        scanf("%lld %lld", &u, &v), puts(ID[u] == ID[v] ? "Yes" : "No");
    return 0;
}
