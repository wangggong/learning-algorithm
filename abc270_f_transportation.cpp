#include <iostream>
#include <climits>
#include <vector>
#include <algorithm>

using namespace std;
typedef long long ll;
typedef pair<ll, pair<ll, ll>> plpll;
const ll N = 2e5 + 10;
ll n, m, u, v, w, ans = LLONG_MAX, pa[N + 10];
vector<plpll> edges[3], es;

void init() { for (ll i = 0; i <= N; i++) pa[i] = i; }
ll query(ll k) { return k != pa[k] ? (pa[k] = query(pa[k])) : pa[k]; }
void merge(ll p, ll q) { pa[query(p)] = query(q); }

void build(ll k) {
    es.clear();
    for (auto e : edges[2])
        es.push_back(e);
    for (ll i = 0; i <= 1; i++)
        if (k & (1 << i))
            for (auto e : edges[i])
                es.push_back(e);
    return;
}

void kruskral(ll n) {
    init();
    sort(es.begin(), es.end());
    ll tot = 0;
    for (auto e : es) {
        u = e.second.first, v = e.second.second, w = e.first;
        if (query(u) != query(v))
            merge(u, v), tot += w, n--;
    }
    if (es.size() > m || n <= 1)
        ans = min(ans, tot);
    return;
}

int main() {
    scanf("%lld %lld", &n, &m);
    for (ll k = 1; k <= 2; k++)
        for (ll i = 1; i <= n; i++)
            scanf("%lld", &w), edges[k - 1].push_back({w, {i, n + k}});
    for (ll i = 0; i < m; i++)
        scanf("%lld %lld %lld", &u, &v, &w), edges[2].push_back({w, {u, v}});
    for (ll k = 0; k < 4; k++)
        build(k), kruskral(n);
    printf("%lld\n", ans);
    return 0;
}
