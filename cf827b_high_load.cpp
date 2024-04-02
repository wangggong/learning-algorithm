#include <iostream>

using namespace std;
typedef long long ll;

ll n, k;

int main() {
    cin >> n >> k;
    ll m = (n - 1) / k;
    ll len = 2 * m + ((n - 1) % k > 1 ? 2 : (n - 1) % k);
    cout << len << endl;
    ll root = m + 1;
    for (ll i = 1; i < root; i++)
        cout << i << ' ' << i + 1 << endl;
    for (ll i = 1; i < k; i++) {
        cout << root << ' ' << root + (i - 1) * m + 1 << endl;
        for (ll j = 1; j < m; j++)
            cout << root + (i - 1) * m + j << ' ' << root + (i - 1) * m + j + 1 << endl;
    }
    ll cur = k * m + 1;
    for (ll i = 1; i <= (n - 1) % k; i++)
        cout << root + i * m << ' ' << cur + i << endl;
    return 0;
}
