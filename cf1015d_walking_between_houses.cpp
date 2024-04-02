#include <iostream>

using namespace std;

typedef long long ll;

int main() {
    ll n, k, s, cur = 1;
    cin >> n >> k >> s;
    if (s < k || (n - 1) * k < s) {
        cout << "NO" << endl;
        return 0;
    }
    cout << "YES" << endl;
    for (ll i = 0; i < k; i++) {
        cur += (1 - (i % 2) * 2) * ((s / k) + (i < s % k ? 1 : 0));
        cout << cur << ' ';
    }
    return 0;
}
