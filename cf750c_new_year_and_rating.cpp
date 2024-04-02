// 参考: https://www.luogu.com.cn/blog/endlesscheng/cf750c-new-year-and-rating-ti-xie

#include <iostream>

using namespace std;

typedef long long ll;
const ll MAXN = 1e9;
const ll N = 1900;
ll n, c, d, s, mx = MAXN, mn = -MAXN;

int main() {
    for (scanf("%lld", &n); n; n--) {
        scanf("%lld %lld", &c, &d);
        if (d == 1)
            mn = max(mn, N - s);
        else
            mx = min(mx, N - 1 - s);
        s += c;
    }
    if (mn > mx)
        puts("Impossible");
    else if (mx == MAXN)
        puts("Infinity");
    else
        printf("%lld\n", mx + s);
    return 0;
}
