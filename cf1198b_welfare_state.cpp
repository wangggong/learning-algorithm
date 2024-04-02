// 参考: https://codeforces.com/problemset/submission/1198/247412775
//
// 后缀数组应用.
#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 2e6;
ll n, a[N + 10], q, op, p, x, t[N + 10], mx[N + 10];

int main() {
    scanf("%lld", &n);
    for (ll i = 1; i <= n; i++)
        scanf("%lld", a + i), t[i] = 1;
    scanf("%lld", &q);
    for (ll i = 1; i <= q; i++) {
        scanf("%lld", &op);
        if (op == 1)
            scanf("%lld%lld", &p, &x), a[p] = x, t[p] = i;
        else
            scanf("%lld", &x), mx[i] = x;
    }
    for (ll i = q - 1; i > 0; i--)
        mx[i] = max(mx[i], mx[i + 1]);
    for (ll i = 1; i <= n; i++)
        printf("%lld ", max(a[i], mx[t[i]]));
    return 0;
}
