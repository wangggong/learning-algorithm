#include <iostream>
using namespace std;
typedef long long ll;
ll n, m, k, t;
int main() {
    cin >> t;
    while (t--) {
        cin >> n >> m >> k;
        if (n == 1) {
            cout << (k == m / 2 ? "YES" : "NO") << endl;
            continue;
        }
        if (m == 1) {
            cout << (k == 0 ? "YES" : "NO") << endl;
            continue;
        }
        if (m % 2 != 0) {
            cout << (k % 2 == 0 && k <= m / 2 * n ? "YES" : "NO") << endl;
            continue;
        }
        if (n % 2 != 0) {
            cout << (k >= m / 2 && (k - m / 2) % 2 == 0 ? "YES" : "NO") << endl;
            continue;
        }
        cout << (k % 2 == 0 ? "YES" : "NO") << endl;
    }
    return 0;
}
