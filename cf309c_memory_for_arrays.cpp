// 参考: https://codeforces.com/contest/309/submission/241246936
//
// 想到了对物品计数, 没想到对背包计数.
// 顺便吐槽一下这个是一个假背包, 本质上还是贪心.
#include <iostream>

using namespace std;
typedef long long ll;
const ll B = 29;
ll n, m, v, a[B + 10], b[B + 10], j, ans;

void setBits(ll k) {
    for (ll bit = 0; k; bit++, k >>= 1)
        if (k & 1)
            a[bit]++;
}

ll first(ll k) {
    for (ll i = k; i <= B; i++)
        if (a[i])
            return i;
    return -1;
}

int main() {
    scanf("%lld%lld", &n, &m);
    for (ll i = 0; i < n; i++)
        scanf("%lld", &v), setBits(v);
    for (ll i = 0; i < m; i++)
        scanf("%lld", &v), b[v]++;
    for (ll i = 0; i <= B; ) {
        if (!b[i] || (j = first(i)) < 0) {
            i++;
            continue;
        }
        ans++, b[i]--, a[j]--, setBits((1ll << j) - (1ll << i));
    }
    printf("%lld", ans);
    return 0;
}
