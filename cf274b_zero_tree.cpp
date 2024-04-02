#include <iostream>
#include <vector>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(nullptr), cout.tie(nullptr)
#define __DEBUG__ true

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;

const ll N = 1e5;
vector<ll> edges[N + 10];
ll n, u, v, W[N + 10];

pll dfs(ll k, ll fa) {
    ll pos = 0, neg = 0;
    for (auto e : edges[k])
        if (e != fa) {
            auto ans = dfs(e, k);
            pos = max(pos, ans.first), neg = max(neg, ans.second);
        }
    W[k] += pos - neg;
    if (W[k] >= 0)
        neg += W[k];
    else
        pos -= W[k];
    return {pos, neg};
}

int main() {
    __FAST__;
    cin >> n;
    for (ll i = 0; i < n - 1; i++)
        cin >> u >> v, edges[u].push_back(v), edges[v].push_back(u);
    for (ll i = 1; i <= n; i++)
        cin >> W[i];
    auto ans = dfs(1, -1);
    cout << ans.first + ans.second << endl;
    return 0;
}
