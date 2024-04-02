#include <iostream>

using namespace std;
typedef long long ll;

ll l, r, t;

ll bitsCnt1(ll n) {
    ll ans = 0;
    for (; n; n &= n - 1)
        ans++;
    return ans;
}

ll getBits(ll l, ll r) {
    if (bitsCnt1(r + 1) == 1)
        return r;
    ll d = 127;
    for (d = 127; d >= 0; d--)
        if ((1ll << d) & r)
            break;
    if (l <= (1ll << d) - 1)
        return (1ll << d) - 1;
    return (1ll << d) + getBits(l - (1ll << d), r - (1ll << d));
}

int main() {
    cin >> t;
    while (t--) {
        cin >> l >> r;
        cout << getBits(l, r) << endl;
    }
    return 0;
}

