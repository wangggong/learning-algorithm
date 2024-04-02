#include <iostream>
#define __FAST_READ__ std::ios::sync_with_stdio(false), std::cin.tie(nullptr), std::cout.tie(nullptr);

using namespace std;
typedef long long ll;

ll h, T, n, k, mn, mx;

int main() {
    __FAST_READ__
    cin >> T;
    while (T--) {
        bool valid = true;
        cin >> n >> k;
        for (int i = 0; i < n; i++) {
            cin >> h;
            if (i == 0) {
                mn = h, mx = h;
                continue;
            }
            if (h + (i == n - 1 ? 1 : 2) * (k - 1) < mn || h > mx + k - 1)
                valid = false;
            mn = max(mn - k + 1, h), mx = min(mx + k - 1, h + k - 1);
        }
        cout << (valid ? "YES" : "NO") << endl;
    }
    return 0;
}
