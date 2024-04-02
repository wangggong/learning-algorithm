#include <iostream>
#include <queue>
#include <unordered_map>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)
#define hash(x, y) ((x) * (N + 10) + (y))

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;

const ll N = 2e5;
ll n, x, y, dirs[5] = {0, 1, 0, -1, 0};
pll a[N + 10], ans[N + 10];

int main() {
    __FAST__;
    cin >> n;
    unordered_map<ll, ll> m;
    for (ll i = 1; i <= n; i++)
        cin >> x >> y, a[i] = {x, y}, m[hash(x, y)] = i;
    queue<ll> Q, Qi;
    for (ll i = 1; i <= n; i++) {
        ll x = a[i].first, y = a[i].second;
        for (ll d = 0; d < 4; d++) {
            ll nx = x + dirs[d], ny = y + dirs[d + 1], h = hash(nx, ny);
            if (!m[h]) {
                ans[i] = {nx, ny}, Q.push(i), Qi.push(i);
                break;
            }
        }
    }
    while (!Qi.empty()) {
        ll q = Qi.front(); Qi.pop();
        ll x = a[q].first, y = a[q].second, h = hash(x, y);
        m[h] = 0;
    }
    while (!Q.empty()) {
        ll q = Q.front(); Q.pop();
        ll x = a[q].first, y = a[q].second;
        for (ll d = 0; d < 4; d++) {
            ll nx = x + dirs[d], ny = y + dirs[d + 1], h = hash(nx, ny);
            if (m.count(h) && m[h]) {
                ans[m[h]] = ans[q], Q.push(m[h]), m[h] = 0;
            }
        }
    }
    for (ll i = 1; i <= n; i++)
        cout << ans[i].first << ' ' << ans[i].second << endl;
    return 0;
}
