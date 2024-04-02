#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;

const ll N = 2e5;
vector<ll> G[N + 10];
ll n, q, u, v, p, x, a[N + 10];

void dfs(ll u, ll p) {
    a[u] += p < 0 ? 0 : a[p];
    for (auto v : G[u]) {
        if (v == p)
            continue;
        dfs(v, u);
    }
}

int main() {
    scanf("%lld %lld", &n, &q);
    for (ll i = 0; i + 1 < n; i++) {
        scanf("%lld %lld", &u, &v);
        G[u].push_back(v);
        G[v].push_back(u);
    }
    for (ll i = 0; i < q; i++) {
        scanf("%lld %lld", &p, &x);
        a[p] += x;
    }
    dfs(1, -1);
    for (ll i = 1; i <= n; i++)
        printf("%lld ", a[i]);
    return 0;
}
