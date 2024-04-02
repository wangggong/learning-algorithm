#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;
const int N = 2e5;
ll n, q, v, l, r, x;
vector<ll> a[N + 10];

ll bisearch(ll x, ll k) {
    ll p = 0, q = a[x].size();
    while (p < q) {
        ll mid = (p + q) >> 1;
        if (a[x][mid] > k)
            q = mid;
        else
            p = mid + 1;
    }
    return p - 1;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 1; i <= n; i++) {
        scanf("%lld", &v);
        a[v].push_back(i);
    }
    scanf("%lld", &q);
    while (q--) {
        scanf("%lld %lld %lld", &l, &r, &x);
        if (a[x].empty()) {
            puts("0");
            continue;
        }
        printf("%lld\n", bisearch(x, r) - bisearch(x, l - 1));
    }
    return 0;
}
