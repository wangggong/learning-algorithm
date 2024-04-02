#include <iostream>

using namespace std;
typedef long long ll;

const ll N = 2e5;
ll a[N + 10], n, s, t;
 
int main() {
    cin >> t;
    while (t--) {
        cin >> n >> s;
        for (ll i = 1; i <= n; i++)
            cin >> a[i];
        ll l = -1, r = -1, cur = 0;
        for (ll p = 1, q = 1; q <= n; ) {
            for (; q <= n && cur >= -s; cur += a[q++]);
            if (cur >= -s && q - p > r - l)
                l = p, r = q;
            if (q - p - 1 > r - l)
                l = p, r = q - 1;
            for (; p <= n && cur < -s; cur -= a[p++]);
        }
        if (r <= l)
            cout << -1 << endl;
        else
            cout << l << " " << (r - 1) << endl;
    }
    return 0;
}
