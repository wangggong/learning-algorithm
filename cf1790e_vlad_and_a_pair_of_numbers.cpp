#include <iostream>
#define __DEBUG__ false

using namespace std;
typedef long long ll;
const ll B = 32;
ll t, n, p, q;

bool get() {
    p = 0, q = 0;
    if (n & 1)
        return false;
    for (ll b = 0; b < B; b++) {
        ll pre = n & (1 << (b + 1)), cur = n & (1 << b);
        if (pre && cur)
            return false;
        if (cur)
            p |= 1 << b;
        else if (pre)
            p |= 1 << b, q |= 1 << b;
    }
    return true;
}

int main() {
    for (scanf("%lld", &t); t; t--) {
        scanf("%lld", &n);
        if (!get())
            puts("-1");
        else {
            printf("%lld %lld\n", p, q);
            if (__DEBUG__)
                puts(p + q == n << 1 && (p ^ q) == n ? "\tTrue" : "\tFalse");
        }
    }
    return 0;
}
