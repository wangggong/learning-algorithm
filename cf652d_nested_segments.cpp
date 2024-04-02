#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
typedef pair<pll, ll> pplll;
const ll N = 2e5;
ll n, v, cnt[N + 10], tr[2 * N + 10], kth[2 * N + 10];
vector<pll> values;
vector<pplll> vec;

ll lowBit(ll x) { return x & (-x); }
ll query(ll x) {
    ll ans = 0;
    for (; x; x -= lowBit(x))
        ans += tr[x];
    return ans;
}
void add(ll x) {
    for (; x <= 2 * n; x += lowBit(x))
        tr[x]++;
    return;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < 2 * n; i++)
        scanf("%lld", &v), values.push_back({v, i});
    sort(values.begin(), values.end());
    for (ll i = 0; i < 2 * n; i++)
        kth[values[i].second] = i + 1;
    for (ll i = 0; i < n; i++)
        vec.push_back({{kth[i * 2], kth[i * 2 + 1]}, i});
    sort(vec.begin(), vec.end(), [](pplll &p, pplll &q) { return p.first.second < q.first.second; });
    for (ll i = 0; i < n; i++) {
        ll l = vec[i].first.first, k = vec[i].second;
        cnt[k] = i - query(l);
        add(l);
    }
    for (ll i = 0; i < n; i++)
        printf("%lld\n", cnt[i]);
    return 0;
}
