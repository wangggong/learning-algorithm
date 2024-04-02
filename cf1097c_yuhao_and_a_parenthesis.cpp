#include <iostream>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(nullptr), cout.tie(nullptr)

using namespace std;
typedef long long ll;

const ll N = 5e5;
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
        if (d >= 0)
            cp[d]++;
        else
            cn[-d]++;
    }
    ans += cp[0] >> 1;
    for (ll i = 1; i < N; i++)
        ans += min(cp[i], cn[i]);
    cout << ans << endl;
    return 0;
}
