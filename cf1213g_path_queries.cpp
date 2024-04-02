#include <iostream>
#include <vector>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(nullptr), cout.tie(nullptr)

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;

const ll N = 2e5;
ll n, m, u, v, w, x, fa[N + 10], sz[N + 10], ans[N + 10];
vector<pll> edges[N + 10];

void init() {
    for (ll i = 0; i <= N; i++)
        fa[i] = i, sz[i] = 1;
}

ll query(ll k) {
    if (fa[k] != k)
        fa[k] = query(fa[k]);
    return fa[k];
}

void merge(ll p, ll q) {
    ll fp = query(p), fq = query(q);
    if (sz[fp] > sz[fq])
        swap(fp, fq);
    fa[fp] = fq, sz[fq] += sz[fp];
    return;
}

int main() {
    __FAST__;
    init();
    cin >> n >> m;
    for (ll i = 0; i < n - 1; i++) {
        cin >> u >> v >> w;
        edges[w].push_back({u, v});
    }
    for (ll i = 1; i <= N; i++) {
        ans[i] = ans[i - 1];
        for (auto e : edges[i])
            ans[i] += sz[query(e.first)] * sz[query(e.second)],
            merge(e.first, e.second);
    }
    for (ll i = 0; i < m; i++)
        cin >> x, cout << ans[x] << " ";
    return 0;
}
