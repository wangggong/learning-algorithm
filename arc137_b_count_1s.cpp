// 别说了, 我是傻逼
#include <iostream>

using namespace std;
typedef long long ll;
ll n, v, _inc, _dec, maxInc, maxDec;

int main() {
    scanf("%lld", &n);
    while (n--) {
        scanf("%lld", &v);
        _inc = max(_inc, 0ll) + (v == 0 ? 1 : -1), maxInc = max(maxInc, _inc);
        _dec = max(_dec, 0ll) + (v == 0 ? -1 : 1), maxDec = max(maxDec, _dec);
    }
    printf("%lld", maxInc + maxDec + 1);
    return 0;
}
