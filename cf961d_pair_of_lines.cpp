#include <iostream>

using namespace std;
typedef long long ll;

const ll N = 1e5;
ll n, a[N + 10][2];

bool _check(ll u, ll v) {
    ll p = -1, q = -1;
    for (ll i = 0; i < n; i++) {
        if (i == u || i == v)
            continue;
        if ((a[i][0] - a[u][0]) * (a[i][1] - a[v][1]) != (a[i][0] - a[v][0]) * (a[i][1] - a[u][1])) {
            if (p < 0)
                p = i;
            else if (q < 0)
                q = i;
            else if ((a[i][0] - a[p][0]) * (a[i][1] - a[q][1]) != (a[i][0] - a[q][0]) * (a[i][1] - a[p][1]))
                return false;
        }
    }
    return true;
}

bool check() {
    if (n <= 4)
        return true;
    return _check(0, 1) || _check(1, 2) || _check(2, 0);
}

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld %lld", a[i], a[i] + 1);
    puts(check() ? "YES" : "NO");
    return 0;
}
