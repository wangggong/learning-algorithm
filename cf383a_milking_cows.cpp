// 参考: https://codeforces.com/problemset/submission/383/252029619
#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 200000;
ll n, a[N + 10], c0, ans;

int main() {
    scanf("%lld", &n); for (ll i = 0; i < n; i++)
        scanf("%lld", a + i), c0 += a[i] == 0;
    for (ll i = 0; i < n; i++)
        if (a[i] == 0)
            c0--;
        else
            ans += c0;
    printf("%lld", ans);
    return 0;
}
