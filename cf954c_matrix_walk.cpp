#include <iostream>
#define __FAST_READ__ std::ios::sync_with_stdio(false), std::cin.tie(nullptr), std::cout.tie(nullptr);

using namespace std;
typedef long long ll;

const ll N = 2e5, MAXN = 1e9;
ll a[N + 10], n, c = -1;

int main() {
    __FAST_READ__
    cin >> n;
    for (ll i = 0; i < n; i++)
        cin >> a[i];
    for (ll i = 1; i < n; i++)
        if (abs(a[i] - a[i - 1]) != 1) {
            c = abs(a[i] - a[i - 1]);
            break;
        }
    if (c == -1) {
        c = 1;
        goto SUCC;
    }
    if (c == 0)
        goto FAIL;
    for (ll i = 1; i < n; i++) {
        ll u = a[i], v = a[i - 1];
        if (u < v)
            swap(u, v);
        if (u == v + 1) {   // right
            if (v % c == 0)
                goto FAIL;
            continue;
        }
        if (u - v != c)     // down
            goto FAIL;
    }
SUCC:
    cout << "YES" << endl << (ll) MAXN << " " << c << endl;
    return 0;
FAIL:
    cout << "NO" << endl;
    return 0;
}
