#include <iostream>

using namespace std;
typedef long long ll;

string s;
ll n, ans;

int main() {
    cin >> s >> n;
    for (char ch : s)
        ans = (ans << 1) | (ch == '1' ? 1 : 0);
    for (ll i = 0, m = s.size(); i < m; i++)
        if (s[i] == '?' && (ans | (1ll << (m - 1 - i))) <= n)
            ans |= 1ll << (m - 1 - i);
    cout << (ans <= n ? ans : -1) << endl;
    return 0;
}
