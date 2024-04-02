#include <iostream>
#include <vector>
#include <algorithm>
#include <unordered_map>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 1500, MAXN = 1e9;
ll n, a[N + 10], S[N + 10], ans, mx;

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    for (ll i = 0; i < n; i++)
        S[i + 1] = S[i] + a[i];
    unordered_map<ll, vector<pll>> M;
    for (ll i = 0; i < n; i++)
        for (ll j = i; j < n; j++)
            M[S[j + 1] - S[i]].push_back({i, j});
    for (auto kv : M) {
        auto k = kv.first;
        auto v = kv.second;
        sort(v.begin(), v.end(), [](pll &p, pll &q) { return p.second < q.second; });
        ll cur = 0, last = -MAXN;
        for (auto pr : v) {
            if (pr.first > last)
                cur++, last = pr.second;
        }
        if (cur > ans)
            ans = cur, mx = k;
    }
    printf("%lld\n", ans);
    sort(M[mx].begin(), M[mx].end(), [](pll &p, pll &q) { return p.second < q.second; });
    ll last = -MAXN;
    for (auto pr : M[mx])
        if (pr.first > last)
            printf("%lld %lld\n", pr.first + 1, pr.second + 1), last = pr.second;
    return 0;
}
