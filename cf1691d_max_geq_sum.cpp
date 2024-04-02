#include <iostream>
#include <algorithm>
#include <stack>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 2e5;
ll t, n, a[N + 10];

bool _check() {
    stack<pll> S;
    ll tot = 0;
    for (ll i = 0; i < n; i++) {
        for (; !S.empty() && S.top().first <= a[i]; S.pop())
            if (S.top().second < tot)
                return false;
        S.push({a[i], tot});
        tot += a[i];
    }
    return true;
}

bool check() {
    if (!_check())
        return false;
    reverse(a, a + n);
    return _check();
}

int main() {
    scanf("%lld", &t);
    while (t--) {
        scanf("%lld", &n);
        for (ll i = 0; i < n; i++)
            scanf("%lld", a + i);
        puts(check() ? "YES" : "NO");
    }
    return 0;
}
