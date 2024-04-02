#include <iostream>
#include <unordered_set>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll n, m, q, v, pa[N + 10], ch[N + 10];

int main() {
    scanf("%lld %lld", &n, &q);
    for (ll i = 2; i <= n; i++)
        scanf("%lld", pa + i), ch[pa[i]]++;
    while (q--) {
        scanf("%lld", &m);
        unordered_set<ll> S;
        while (m--)
            scanf("%lld", &v), S.insert(v);
        ll ans = 0;
        for (auto v : S)
            ans += ch[v] + (S.count(pa[v]) ? -1 : 1);
        printf("%lld\n", ans);
    }
    return 0;
}
