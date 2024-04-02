#include <iostream>
#include <cstring>
#include <queue>

using namespace std;
typedef long long ll;
typedef pair<pair<ll, ll>, ll> pplll;
const ll N = 2e5;
ll a[N + 10][2], inv[N + 10][2], c[N + 10], idx[N + 10], t, n;

bool check() {
    memset(c, 0, sizeof c);
    memset(idx, 0, sizeof idx);
    scanf("%lld", &n);
    bool valid = true;
    for (ll i = 1; i <= n; i++) {
        scanf("%lld%lld", a[i], a[i] + 1);
        ll x = a[i][0], y = a[i][1];
        if (c[x] >= 2 || c[y] >= 2)
            valid = false;
        inv[x][c[x]++] = i, inv[y][c[y]++] = i;
    }
    if (!valid)
        return false;
    for (ll i = 1; i <= n; i++)
        if (c[i] != 2)
            return false;
    queue<pplll> Q;
    for (ll i = 1; i <= n; i++) {
        if (idx[i])
            continue;
        idx[i] = 1;
        Q.push({{i, a[i][0]}, idx[i]}), Q.push({{i, a[i][1]}, idx[i]});
        while (!Q.empty()) {
            auto q = Q.front(); Q.pop();
            ll pos = q.first.first, x = q.first.second, k = q.second;
            ll j = inv[x][0] + inv[x][1] - pos;
            if (idx[j]) {
                if (idx[j] == 3 - k)
                    continue;
                else
                    return false;
            }
            idx[j] = 3 - k;
            Q.push({{j, a[j][0]}, idx[j]}), Q.push({{j, a[j][1]}, idx[j]});
        }
    }
    return true;
}

int main() {
    scanf("%lld", &t); while (t--)
        puts(check() ? "YES" : "NO");
    return 0;
}
