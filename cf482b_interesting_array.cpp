#include <iostream>
#include <cstring>
#define __FAST_IO__ ios::sync_with_stdio(false), cin.tie(nullptr), cout.tie(nullptr);

using namespace std;
typedef long long ll;
typedef pair<ll, pair<ll, ll>> plpll;

const ll N = 1e5;
ll n, m, l, r, q, diff[N + 10], ans[N + 10], sum[N + 10];
plpll a[N + 10];

int main() {
    cin >> n >> m;
    for (ll i = 0; i < m; i++)
        cin >> l >> r >> q, a[i] = {q, {l - 1, r}};
    for (ll d = 0; d < 30; d++) {
        memset(diff, 0, sizeof diff);
        for (ll i = 0; i < m; i++)
            if (a[i].first & (1 << d))
                diff[a[i].second.first]--, diff[a[i].second.second]++;
        for (ll i = 0; i < n; i++) {
            diff[i + 1] += diff[i], sum[i + 1] = sum[i];
            if (diff[i])
                ans[i] |= 1 << d;
            else
                sum[i + 1]++;
        }
        for (ll i = 0; i < m; i++)
            if ((a[i].first & (1 << d)) == 0 && sum[a[i].second.first] == sum[a[i].second.second]) {
                cout << "NO" << endl;
                return 0;
            }
    }
    cout << "YES" << endl;
    for (ll i = 0; i < n; i++)
        cout << ans[i] << " ";
    return 0;
}
