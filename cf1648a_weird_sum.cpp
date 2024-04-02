#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;
typedef long long ll;
const ll N = 1e5;
ll n, m, v, ans;
vector<ll> xs[N + 10], ys[N + 10];

void calc(vector<ll> &vec) {
    if (vec.size() == 0) return;
    // for (auto k : vec) cout << k << ' '; cout << endl;
    sort(vec.begin(), vec.end());
    ll n = vec.size(), S = 0;
    for (ll k : vec)
        S += k;
    for (ll k : vec)
        ans += S - n * k, S -= k, n--;
}

int main() {
    scanf("%lld %lld", &n, &m);
    for (ll i = 0; i < n; i++)
        for (ll j = 0; j < m; j++) {
            scanf("%lld", &v);
            xs[v].push_back(i), ys[v].push_back(j);
        }
    for (auto vec : xs)
        calc(vec);
    for (auto vec : ys)
        calc(vec);
    printf("%lld\n", ans);
    return 0;
}
