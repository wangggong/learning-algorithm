#include <iostream>
#include <algorithm>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 2e5;
ll n, s, t, e1 = -1, e2 = -1;
pll a[N + 10];

bool check() {
    for (ll i = 0; i < n; i++) {
        s = a[i].first, t = a[i].second;
        if (e1 < s)
            e1 = t;
        else if (e2 < s)
            e2 = t;
        else
            return false;
    }
    return true;
}

int main() {
    scanf("%lld", &n); for (ll i = 0; i < n; i++)
        scanf("%lld%lld", &s, &t), a[i] = make_pair(s, t);
    sort(a, a + n);
    puts(check() ? "YES" : "NO");
    return 0;
}
