#include <iostream>
#include <algorithm>
#include <queue>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 1e5;
ll n, x, ans[N + 10];
pll a[N + 10];

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", &x), a[i] = {x, i};
    sort(a, a + n);
    priority_queue<ll> Q;
    for (ll i = 0; i < n; i++) {
        ans[a[i].second] = max(Q.empty() ? 0 : Q.top(), a[i].second) - a[i].second - 1;
        Q.push(a[i].second);
    }
    for (ll i = 0; i < n; i++)
        printf("%lld ", ans[i]);
    return 0;
}
