#include <iostream>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(nullptr), cout.tie(nullptr)

using namespace std;
typedef long long ll;

const ll N = 2e5, MAXN = 2e14;
ll n, m, a[N + 10], fa[N + 10], c[N + 10], op, k, v;

ll query(ll k) {
    if (fa[k] != k)
        fa[k] = query(fa[k]);
    return fa[k];
}

int main() {
    __FAST__;
    cin >> n;
    for (ll i = 1; i <= n; i++)
        cin >> c[i], fa[i] = i;
    c[n + 1] = MAXN + 10, fa[n + 1] = n + 1;
    cin >> m;
    for (ll i = 0; i < m; i++) {
        cin >> op;
        switch (op) {
        case 1:
            cin >> k >> v;
            while (a[k] + v > c[k])
                v -= c[k] - a[k], a[k] = c[k], fa[query(k)] = query(k + 1), k = query(k);
            a[k] += v;
            break;
        case 2:
            cin >> k;
            cout << a[k] << endl;
            break;
        }
    }
    return 0;
}
