#include <iostream>
#include <cstring>
#define SWAP(x,y) (x)+=(y),(y)=(x)-(y),(x)-=(y)

using namespace std;
typedef long long ll;
const ll N = 1e5;
ll t, n, m, a, b, last[N + 10];

int main() {
    for (scanf("%lld", &t); t--; ) {
        memset(last, 0, sizeof last);
        for (scanf("%lld%lld", &n, &m); m--; ) {
            scanf("%lld%lld", &a, &b);
            if (a > b)
                SWAP(a, b);
            last[b] = max(last[b], a);
        }
        ll ans = 0;
        for (ll p = 1, q = 1; q <= n; q++)
            p = max(p, last[q] + 1), ans += q - p + 1;
        printf("%lld\n", ans);
    }
    return 0;
}
