#include <iostream>
#include <algorithm>
#include <set>

using namespace std;
typedef long long ll;

const ll N = 2e5;
ll t, n, m, A[N + 10], B[N + 10], a[N + 10], b[N + 10], lena, lenb;
set<ll> S;

ll solve(ll lena, ll lenb) {
    ll cnt = 0;
    S.clear();
    for (ll i = 0; i < lena; i++)
        S.insert(a[i]);
    for (ll i = 0; i < lenb; i++)
        if (S.count(b[i]))
            cnt++;
    ll ans = cnt;
    for (ll i = 0; i < lenb; i++) {
        if (S.count(b[i])) {
            cnt--;
            continue;
        }
        ll pp = upper_bound(a, a + lena, b[i]) - a;
        ll qq = upper_bound(b, b + lenb, b[i] - pp) - b;
        ans = max(ans, cnt + i - qq + 1);
    }
    return ans;
}

int main() {
    for (cin >> t; t; t--) {
        cin >> n >> m;
        for (ll i = 0; i < n; i++)
            cin >> A[i];
        for (ll i = 0; i < m; i++)
            cin >> B[i];
        lena = 0, lenb = 0;
        for (ll i = 0; i < n; i++)
            if (A[i] > 0)
                a[lena++] = A[i];
        for (ll i = 0; i < m; i++)
            if (B[i] > 0)
                b[lenb++] = B[i];
        ll ans = solve(lena, lenb);
        lena = 0, lenb = 0;
        for (ll i = n - 1; i >= 0; i--)
            if (A[i] < 0)
                a[lena++] = -A[i];
        for (ll i = m - 1; i >= 0; i--)
            if (B[i] < 0)
                b[lenb++] = -B[i];
        ans += solve(lena, lenb);
        cout << ans << endl;
    }
    return 0;
}
