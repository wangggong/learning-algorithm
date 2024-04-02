#include <iostream>
#include <cstring>
#include <vector>

using namespace std;
typedef long long ll;
const ll M = 4e5;
ll n, u, v, idx, mp[M + 10], pa[M + 10], sz[M + 10], ans;
vector<ll> G[M + 10];
bool visit[M + 10];

ll query(ll k) {
    return pa[k] == k ? pa[k] : (pa[k] = query(pa[k]));
}

void merge(ll p, ll q) {
    sz[query(q)] += sz[query(p)];
    pa[query(p)] = query(q);
}

bool dfs(ll u, ll p) {
    visit[u] = true;
    bool valid = true;
    for (ll v : G[u])
        if (v != p)
            if (visit[v] || !dfs(v, u)) {
                valid = false;
            }
    return valid;
}

int main() {
    memset(mp, -1, sizeof mp);
    for (ll i = 0; i <= M; i++)
        pa[i] = i, sz[i] = 1;
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++) {
        scanf("%lld %lld", &u, &v);
        if (mp[u] < 0)
            mp[u] = idx++;
        if (mp[v] < 0)
            mp[v] = idx++;
        G[mp[u]].push_back(mp[v]);
        G[mp[v]].push_back(mp[u]);
        if (query(mp[u]) != query(mp[v]))
            merge(mp[u], mp[v]);
    }
    for (ll i = 0; i < idx; i++)
        if (!visit[i]) {
            ans += dfs(i, -1) ? sz[query(i)] - 1 : sz[query(i)];
        }
    printf("%lld\n", ans);
    return 0;
}
