#include <iostream>
#include <cstring>
#include <vector>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll t, n, a[N + 10][2];

int main() {
    scanf("%lld", &t);
    while (t--) {
        scanf("%lld", &n);
        vector<ll> cx(n + 1);
        vector<ll> cy(n + 1);
        for (ll i = 0; i < n; i++)
            scanf("%lld %lld", &a[i][0], &a[i][1]), cx[a[i][0]]++, cy[a[i][1]]++;
        ll ans = n * (n - 1) * (n - 2) / 6;
        for (ll i = 0; i < n; i++)
            ans -= (cx[a[i][0]] - 1) * (cy[a[i][1]] - 1);
        printf("%lld\n", ans);
    }
    return 0;
}
