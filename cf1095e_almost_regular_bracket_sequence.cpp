#include <iostream>

using namespace std;
typedef long long ll;
ll n, tot, ans, m = -1;
string s;
bool neg;

int main() {
    cin >> n >> s;
    for (ll i = 0; i < n; i++) {
        tot += s[i] == '(' ? 1 : -1;
        if (tot < -2)
            goto OUT0;
        if (tot < 0)
            neg = true;
        if (tot == -1 && m < 0)
            m = i;
    }
    if (tot == 2) {
        if (neg)
            goto OUT0;
        for (ll i = n - 1; tot >= 2; i--)
            if (s[i] == '(')
                ans++, tot--;
            else
                tot++;
    } else if (tot == -2) {
        for (ll i = 0; i <= m; i++)
            if (s[i] == ')')
                ans++;
    } else;
    cout << ans << endl;
    return 0;
OUT0:
    cout << 0 << endl;
}
