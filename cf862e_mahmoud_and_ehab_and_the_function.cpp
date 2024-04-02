#include <iostream>
#include <algorithm>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)
#define __ODD__(x, a, b) ((x) % 2 ? (a) : (b))
#define DEBUG false

using namespace std;
typedef long long ll;

const ll N = 1e5;
ll n, m, q, l, r, v, a, b[N + 10], Sa, Sb[N + 10];

ll query() {
    auto it = lower_bound(Sb + 1, Sb + m - n + 2, Sa);
    if (it == Sb + 1)
        return abs(Sa - *it);
    if (it == Sb + m - n + 2)
        return abs(Sa - *(Sb + m - n + 1));
    return min(abs(Sa - *it), abs(Sa - *(it - 1)));
}

int main() {
    __FAST__;
    cin >> n >> m >> q;
    for (ll i = 1; i <= n; i++)
        cin >> a, Sa += a * __ODD__(i, 1, -1);
    for (ll i = 1; i <= m; i++)
        cin >> b[i];
    for (ll i = 1; i <= n; i++)
        Sb[1] += b[i] * __ODD__(i, 1, -1);
    for (ll i = n + 1; i <= m; i++)
        Sb[i - n + 1] = -(Sb[i - n] - b[i - n]) + b[i] * __ODD__(n, 1, -1);
    if (DEBUG) {
        cout << '\t';
        for (ll i = 1; i <= m - n + 1; i++)
            cout << Sb[i] << ' '; cout << endl;
    }
    sort(Sb + 1, Sb + m - n + 2);
    if (DEBUG) {
        cout << '\t';
        for (ll i = 1; i <= m - n + 1; i++)
            cout << Sb[i] << ' '; cout << endl;
    }
    cout << query() << endl;
    for (ll i = 0; i < q; i++) {
        cin >> l >> r >> v;
        Sa += __ODD__(l, __ODD__(r, v, 0), __ODD__(r, 0, -v));
        cout << query() << endl;
    }
    return 0;
}
