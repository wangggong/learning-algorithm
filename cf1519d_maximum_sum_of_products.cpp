#include <iostream>

using namespace std;
typedef long long ll;

const int N = 5e3;
ll a[N + 10], b[N + 10], n;

int main() {
    cin >> n;
    for (int i = 0; i < n; i++)
        cin >> a[i];
    for (int i = 0; i < n; i++)
        cin >> b[i];
    ll tot = 0, diff = 0;
    for (int i = 0; i < n; i++)
        tot += a[i] * b[i];
    for (int i = 0; i < n; i++) {
        ll cur = 0;
        for (int j = 1; i - j >= 0 && i + j < n; j++)
            cur += b[i - j] * (a[i + j] - a[i - j]) + b[i + j] * (a[i - j] - a[i + j]), diff = max(diff, cur); // cout << i - j << " " << i + j << " " << cur << endl;
    }
    for (int i = 0; i + 1 < n; i++) {
        ll cur = 0;
        for (int j = 0; i - j >= 0 && i + j + 1 < n; j++)
            cur += b[i + j + 1] * (a[i - j] - a[i + j + 1]) + b[i - j] * (a[i + j + 1] - a[i - j]), diff = max(diff, cur); // cout << i - j << " " << i + j + 1 << " " << cur << endl;
    }
    // cout << tot << endl;
    cout << tot + diff << endl;
    return 0;
}
