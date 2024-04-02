#include <iostream>
#include <vector>
#include <algorithm>
#include <queue>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
typedef pair <pll, ll> pplll;
ll n, m, pl, pr, l, r, k;

// 是我不配的带模拟...
int main() {
    cin >> n >> m >> pl >> pr;
    vector<pplll> a;
    for (ll i = 0; i < n - 1; i++)
        cin >> l >> r, a.push_back({{l - pr, r - pl}, i}), pl = l, pr = r;
    vector<pll> b;
    for (ll i = 1; i <= m; i++)
        cin >> k, b.push_back({k, i});
    sort(a.begin(), a.end());
    sort(b.begin(), b.end());
    priority_queue<pll, vector<pll>, greater<pll>> Q;
    ll k = 0;
    vector<ll> ans(n - 1, 0);
    for (auto v : b) {
        for (; k < n - 1 && a[k].first.first <= v.first; k++)
            Q.push({a[k].first.second, a[k].second});
        if (!Q.empty()) {
            auto q = Q.top();
            if (q.first < v.first)
                break;
            Q.pop();
            ans[q.second] = v.second;
        }
    }
    if (!Q.empty() || k < n - 1)
        cout << "No" << endl;
    else {
        cout << "Yes" << endl;
        for (auto a : ans)
            cout << a << " ";
    }
    return 0;
}
