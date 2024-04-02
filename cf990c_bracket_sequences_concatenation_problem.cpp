#include <iostream>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(nullptr), cout.tie(nullptr)

using namespace std;
typedef long long ll;

const ll N = 3e5;
ll n, ans, cp[N + 10], cn[N + 10];
string s;

bool valid(ll m) {
    ll d = 0;
    if (m >= 0)
        for (ll i = 0; i < s.size(); i++) {
            d += s[i] == '(' ? 1 : -1;
            if (d < 0)
                return false;
        }
    else
        for (ll i = s.size() - 1; i >= 0; i--) {
            d += s[i] == ')' ? 1 : -1;
            if (d < 0)
                return false;
        }
    return true;
}

int main() {
    __FAST__;
    cin >> n;
    for (ll i = 0; i < n; i++) {
        cin >> s;
        ll d = 0;
        for (ll j = 0; j < s.size(); j++)
            d += s[j] == '(' ? 1 : -1;
        if (!valid(d))
            continue;
        if (d > 0)
            cp[d]++;
        else if (d < 0)
            cn[-d]++;
        else
            cp[d]++, cn[d]++;
    }
    for (ll i = 0; i < N; i++)
        ans += cp[i] * cn[i];
    cout << ans << endl;
    return 0;
}
