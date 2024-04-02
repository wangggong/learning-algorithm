#include <iostream>

using namespace std;

typedef long long ll;
const ll N = 2e5;
ll n, a[N + 10], ans;

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    for (ll p = 0, q = 0, mask = 0; p < n; p++) {
        for (; q < n && (mask & a[q]) == 0; q++)
            mask |= a[q];
        ans += q - p;
        mask ^= a[p];
    }
    printf("%lld\n", ans);
    return 0;
}
