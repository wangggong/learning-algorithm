/*
 * 如果 n 是奇数，输出 -1，因为正中间的那个字母等于它自己，不满足题目要求。
 * 
 * 如果 s 有个字母 x，其出现次数超过 n/2，那么输出 -1。即使左半边都是 x，也仍然有多出的 x 在右半边，不满足题目要求。
 * 
 * 否则可以把 s 变成反回文串。
 * 定义 diff[k] 为满足 s[i]=s[n-1-i]=k 的出现次数，其中 i 从 0 到 n/2-1。
 * 设 m = max(diff), tot = sum(diff)。
 * 我们应该优先「内部消化」，例如一对 a 和一对 b 都在对称位置上（对称位置指 i 和 n-1-i），那么交换其中一个 a 和其中一个 b 就可以让这两对字母都不一样。
 * 如果 tot 是奇数，我们可以找一个 s[i]!=s[n-1-i] 的字母交换，这样的字母一定存在，因为所有字母的出现次数都不超过 n/2。
 * 看上去答案是 ceil(tot/2)。
 * 
 * 但是 (万恶的但是)，如果有一对字母的出现次数特别多，即 m*2 > tot，我们必须把其它所有不同于 m 的字母都拿来交换，所以交换次数是 m。这样的操作一定可以完成，因为所有字母的出现次数都不超过 n/2。
 * 综上所述，答案为 max(ceil(tot/2), m)。
 * 
 */
#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 2e5, A = 26;
ll t, n, c[A + 10], d[A + 10];
char s[N + 10];

ll getCount() {
    if (n % 2)
        return -1;
    memset(c, 0, sizeof c), memset(d, 0, sizeof d);
    ll cnt = 0, mx = 0;
    for (ll i = 0; i < n; i++) {
        c[s[i] - 'a']++;
        if (i < n - 1 - i && s[i] == s[n - 1 - i])
            d[s[i] - 'a']++, cnt++;
    }
    for (ll i = 0; i < A; i++) {
        if (c[i] * 2 > n)
            return -1;
        mx = max(mx, d[i]);
    }
    return max((cnt + 1) / 2, mx);
}

int main() {
    for (scanf("%lld", &t); t--; )
        scanf("%lld%s", &n, s), printf("%lld\n", getCount());
    return 0;
}
