#include <iostream>

using namespace std;
typedef long long ll;

ll t;
string s;

int main() {
    cin >> t;
    while (t--) {
        cin >> s;
        ll ans = 0, n = s.size();
        for (ll i = 0, last = -1; i < n; last = i, i++) {
            for (; i < n && s[i] == '0'; i++);
            if (i == n)
                break;
            int c = 0;
            for (ll j = i; j < n; j++) {
                c = (c << 1) | (s[j] - '0');
                if (j + 1 - c <= last)
                    break;
                ans++;
            }
        }
        cout << ans << endl;
    }
    return 0;
}
