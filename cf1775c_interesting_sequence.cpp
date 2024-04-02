// 参考: https://codeforces.com/problemset/submission/1775/249584540
#include <iostream>

using namespace std;
typedef long long ll;
ll t, n, x;

ll get(ll n, ll x) {
    if (n == x)
        return n;
    while (n > x) {
        ll lowBit = n & -n;
        n -= lowBit;
        if (n == x && !(n & (lowBit << 1)))
            return n | (lowBit << 1);
    }
    return -1;
}

int main() {
    scanf("%lld", &t); while (t--)
        scanf("%lld%lld", &n, &x), printf("%lld\n", get(n, x));
    return 0;
}
