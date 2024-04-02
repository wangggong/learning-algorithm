#include <iostream>
#include <algorithm>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll n, m, p, A[N + 10], B[N + 10], S[N + 10], ans;

ll search(ll k) {
    ll p = 0, q = m;
    while (p < q) {
        ll mid = (p + q) >> 1;
        if (B[mid] >= k)
            q = mid;
        else
            p = mid + 1;
    }
    return p;
}

int main() {
    scanf("%lld %lld %lld", &n, &m, &p);
    for (ll i = 0; i < n; i++)
        scanf("%lld", A + i);
    for (ll i = 0; i < m; i++)
        scanf("%lld", B + i);
    sort(A, A + n);
    sort(B, B + m);
    for (ll i = 0; i < m; i++)
        S[i + 1] = S[i] + B[i];
    for (ll i = 0; i < n; i++) {
        ll l = search(p - A[i]);
        ans += p * (m - l) + A[i] * l + S[l];
    }
    printf("%lld", ans);
    return 0;
}
