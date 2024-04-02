// x, y => x&y, x|y
// x = 1 << a, y = 1 << b => (1 << a)^2 + (1 << b)^2
// x&y = 0, x|y = 1 << b + 1 << a => 0 + (1 << a)^2 + (1 << b)^2 + 2*(1 << a)*(1 << b)
#include <iostream>

using namespace std;
typedef long long ll;

const ll D = 32;
ll cnt[D], n, k, ans;

int main() {
    scanf("%lld", &n);
    for (int i = 0; i < n; i++) {
        scanf("%lld", &k);
        for (int d = 0; d < D; d++)
            if (((k >> d) & 1) != 0)
                cnt[d]++;
    }
    while (n--) {
        ll v = 0;
        for (int d = 0; d < D; d++)
            if (cnt[d] > 0)
                v |= 1 << d, cnt[d]--;
        ans += v * v;
    }
    printf("%lld\n", ans);
    return 0;
}
