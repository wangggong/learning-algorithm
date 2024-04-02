#include <iostream>
#include <vector>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)

using namespace std;
typedef long long ll;
string s, t;
ll n, ans;

int main() {
    __FAST__;
    cin >> n;
    cin >> s >> t;
    vector<ll> s0, t0;
    for (ll i = 0; i < n; i++)
        if (s[i] == '0')
            s0.push_back(i);
    for (ll i = 0; i < n; i++)
        if (t[i] == '0')
            t0.push_back(i);
    if (s0.size() != t0.size())
        goto FAIL;
    for (ll i = 0, m = s0.size(); i < m; i++)
        ans += s0[i] != t0[i];
    cout << ans << endl;
    return 0;
FAIL:
    cout << -1 << endl;
    return 0;
}
