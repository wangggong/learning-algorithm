// 参考: https://atcoder.jp/contests/abc279/submissions/46217952
//
// 一眼并查集, 但具体实现上有点懵.
// 利用正排倒排两个索引建立箱子编号和箱子号码的双射, 实现 "合并箱子后老号码的复用".
#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 3e5;
ll n, mx, q, k, x, y, pa[N * 2 + 10], inv[N * 2 + 10], idx[N * 2 + 10];
ll query(ll q) {
    return pa[q] = q == pa[q] ? pa[q] : query(pa[q]);
}

int main() {
    scanf("%lld %lld", &n, &q);
    for (ll i = 1; i <= n; i++)
        pa[i] = i, inv[i] = i, idx[i] = i;
    mx = n + q;
    while (q--) {
        scanf("%lld %lld", &k, &x);
        switch (k) {
            case 1:
                scanf("%lld", &y);
                pa[inv[y]] = inv[x], pa[mx] = mx, inv[y] = mx, idx[mx--] = y;
                break;
            case 2:
                pa[++n] = inv[x];
                break;
            case 3:
                printf("%lld\n", idx[query(x)]);
                break;
        }
    }
    return 0;
}
