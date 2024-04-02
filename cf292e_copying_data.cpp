#include <iostream>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(nullptr), cout.tie(nullptr)
#define L 0
#define R 1
#define D 2

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;

const ll N = 1e5, MAXN = 1e9;
ll n, m, op, x, y, k, a[N + 10], b[N + 10], seg[(N + 10) << 2][3];

void build(ll l, ll r, ll o) {
    seg[o][L] = l, seg[o][R] = r, seg[o][D] = MAXN;
    if (l == r)
        return;
    ll mid = l + r >> 1;
    build(l, mid, o << 1);
    build(mid + 1, r, o << 1 | 1);
}

void update(ll l, ll r, ll o, ll d) {
    if (l <= seg[o][L] && seg[o][R] <= r) {
        seg[o][D] = d;
        return;
    }
    ll dd = seg[o][D];
    if (dd != MAXN)
        seg[o << 1][D] = dd, seg[o << 1 | 1][D] = dd, seg[o][D] = MAXN;
    ll mid = seg[o][L] + seg[o][R] >> 1;
    if (l <= mid)
        update(l, r, o << 1, d);
    if (mid < r)
        update(l, r, o << 1 | 1, d);
    return;
}

ll query(ll o, ll i) {
    if (seg[o][D] != MAXN || seg[o][L] == seg[o][R])
        return seg[o][D];
    ll mid = seg[o][L] + seg[o][R] >> 1;
    return query(i <= mid ? o << 1 : o << 1 | 1, i);
}

int main() {
    __FAST__;
    cin >> n >> m;
    for (ll i = 1; i <= n; i++)
        cin >> a[i];
    for (ll i = 1; i <= n; i++)
        cin >> b[i];
    build(1, n, 1);
    while (m--) {
        cin >> op >> x;
        switch (op) {
        case 1:
            cin >> y >> k;
            update(y, y + k - 1, 1, x - y);
            break;
        case 2:
            ll d = query(1, x);
            cout << (d == MAXN ? b[x] : a[x + d]) << endl;
        }
    }
    return 0;
}
