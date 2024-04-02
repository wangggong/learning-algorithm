#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 1e5;
ll t, a, b;
char s[N + 10];

ll get() {
    ll l = 0, r = strlen(s), ans = 0;
    for (; l < r && s[l] == '0'; l++);
    for (; l < r && s[r - 1] == '0'; r--);
    if (l == r)
        return ans;
    for (ll p = l, q = l; q < r;) {
        for (p = q; q < r && s[q] == '1'; q++);
        if (q == r) {
            ans += a;
            break;
        }
        for (p = q; q < r && s[q] == '0'; q++);
        ans += min(a, (q - p) * b);
    }
    return ans;
}

int main() {
    scanf("%lld", &t); while (t--)
        scanf("%lld%lld%s", &a, &b, s), printf("%lld\n", get());
    return 0;
}
