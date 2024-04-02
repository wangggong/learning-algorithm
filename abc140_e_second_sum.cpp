#include <iostream>
#include <vector>
#include <stack>

using namespace std;
typedef long long ll;

ll n, v, ans;

int main() {
    scanf("%lld", &n);
    vector<ll> L(n + 2), R(n + 1), pos(n + 1);
    for (ll i = 1; i <= n; i++)
        scanf("%lld", &v), pos[v] = i, L[i] = i - 1, R[i] = i + 1;
    for (ll v = 1; v <= n; v++) {
        ll i = pos[v];
        ll l = L[i], r = R[i];
        if (l > 0)
            ans += v * (l - L[l]) * (r - i);
        if (r <= n)
            ans += v * (R[r] - r) * (i - l);
        L[r] = l, R[l] = r;
    }
    printf("%lld\n", ans);
    return 0;
}
