// 参考: https://codeforces.com/contest/1799/submission/228141047
//
// 偶数的时候前后均匀分布; 奇数的时候把单个字符放在最后 (除了后面只有一种的情况). 为啥?
#include <iostream>
#include <cstring>
#include <algorithm>

using namespace std;
typedef long long ll;
const ll N = 1e5;
ll t;
char s[N + 10], ans[N + 10];

int main() {
    for (scanf("%lld", &t); t; t--) {
        memset(ans, 0, sizeof ans);
        scanf("%s", s);
        ll n = strlen(s);
        sort(s, s + n);
        ll i = 0, p = 0, q = n - 1;
        for (; i + 1 < n; i += 2) {
            if (s[i] != s[i + 1]) {
                if (s[i + 1] == s[n - 1])
                    swap(s[i], s[(i + n) >> 1]);
                else
                    ans[q--] = s[i], i++;
                break;
            }
            ans[p++] = s[i], ans[q--] = s[i];
        }
        while (i < n)
            ans[p++] = s[i], i++;
        printf("%s\n", ans);
    }
    return 0;
}
