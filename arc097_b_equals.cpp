#include <unordered_map>
#include <set>

using namespace std;
typedef long long ll;
const ll N = 1e5;
ll n, m, a[N + 10], u, v, pa[N + 10], ans;
set<int> S[N + 10];

ll query(ll k) {
    if (pa[k] != k) {
        pa[k] = query(pa[k]);
    }
    return pa[k];
}

void merge(ll p, ll q) { pa[query(p)] = query(q); }

int main() {
    scanf("%lld %lld", &n, &m);
    for (ll i = 1; i <= n; i++)
        scanf("%lld", a + i);
    for (ll i = 1; i <= n; i++)
        pa[i] = i;
    for (ll i = 0; i < m; i++) {
        scanf("%lld %lld", &u, &v);
        merge(u, v);
    }
    for (ll i = 1; i <= n; i++)
        S[query(i)].insert(i);
    for (ll i = 1; i <= n; i++)
        if (S[query(i)].count(a[i]))
            ans++;
    printf("%lld\n", ans);
    return 0;
}
