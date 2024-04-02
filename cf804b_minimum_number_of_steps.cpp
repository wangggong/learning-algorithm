#include <iostream>

using namespace std;
typedef long long ll;
const ll MOD = 1e9 + 7;
string s;
ll ans, cnt;

int main() {
    cin >> s;
    for (ll i = s.size() - 1; i >= 0; i--)
        if (s[i] == 'b')
            cnt++;
        else {
            ans = (ans + cnt) % MOD;
            cnt = cnt * 2 % MOD;
        }
    cout << ans << endl;
    return 0;
}
