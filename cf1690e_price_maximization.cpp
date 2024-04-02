// 参考: https://codeforces.com/contest/1690/submission/227301239
//
// 相向双指针
#include <iostream>
#include <algorithm>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll t, n, k, a[N + 10];

int main() {
    for (scanf("%lld", &t); t; t--) {
        scanf("%lld %lld", &n, &k);
        ll ans = 0;
        for (ll i = 0; i < n; i++)
            scanf("%lld", a + i), ans += a[i] / k, a[i] %= k;
        sort(a, a + n);
        for (ll i = 0, j = n - 1; i < j; i++)
            if (a[i] + a[j] >= k)
                ans++, j--;
        printf("%lld\n", ans);
    }
    return 0;
}
