#include <iostream>
#include <algorithm>
#include <climits>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 2e5;
ll t, n, a[N + 10], b, S[N + 10];
pll A[N + 10];

int main() {
    for (scanf("%lld", &t); t; t--) {
        scanf("%lld", &n);
        for (ll i = 0; i < n; i++)
            scanf("%lld", a + i);
        for (ll i = 0; i < n; i++)
            scanf("%lld", &b), A[i] = {a[i], b};
        sort(A, A + n, [](const pll &p, const pll &q) {
            return p.first < q.first || (p.first == q.first && p.second > q.second);
        });
        for (ll i = 0; i < n; i++)
            S[i + 1] = S[i] + A[i].second;
        ll ans = S[n];
        for (ll i = 0; i < n; i++)
            ans = min(ans, max(A[i].first, S[n] - S[i + 1]));
        printf("%lld\n", ans);
    }
    return 0;
}
