#include <iostream>

using namespace std;
typedef long long ll;

ll l, r, ans;

ll bitcnt(ll n) {
    ll ans = 0;
    for (; n; n >>= 1, ans++);
    return ans;
}

int main() {
    cin >> l >> r;
    cout << ((1LL << bitcnt(l ^ r)) - 1LL) << endl;
    return 0;
}
