#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll n, m, v, a[N + 10];
vector<ll> ops[N + 10];

int main() {
    scanf("%lld %lld", &n, &m);
    for (ll i = 1; i <= n; i++) {
        scanf("%lld", &v);
        ll l = v >= 0 ? 1 : (-v + i - 1) / i;
        ll r = min(m, (n - 1 - v) / i);
        for (ll j = l; j <= r; j++)
            ops[j].push_back(v + j * i);
    }
    for (ll i = 1; i <= m; i++) {
        for (auto v : ops[i])
            a[v] = i;
        ll mex = 0;
        while (a[mex++] == i);
        printf("%lld\n", mex - 1);
    }
    return 0;
}
