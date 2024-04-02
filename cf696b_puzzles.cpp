// 参考: https://codeforces.com/problemset/submission/696/247605500
//
// 考虑子树 `i`, 有若干个兄弟, 每个兄弟 `j` 有 1/2 的概率出现在其前面 / 后面. 导致期望的贡献为 `size[T_j] / 2.0, i != j`. 求和就是 `(n - size[T_i]) / 2.0`.
#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 1e5;
ll n, p[N + 10], s[N + 10], d[N + 10];

int main() {
    scanf("%lld", &n); for (ll i = 1; i < n; i++)
        scanf("%lld", p + i), p[i]--;
    for (ll i = n - 1; i > 0; i--)
        s[i]++, s[p[i]] += s[i]; s[0]++;
    for (ll i = 0; i < n; i++)
        d[i] = d[p[i]] + 1, printf("%.1f ", (n - s[i] + d[i] + 1) / 2.0);
    return 0;
}
