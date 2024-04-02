#include <iostream>

using namespace std;
typedef long long ll;
const ll MOD = 1e9 + 7;
ll n, m, v, x, y;

int main() {
    scanf("%lld %lld", &n, &m);
    for (ll i = 0; i < n; i++) {
        scanf("%lld", &v);
        x = (x + (i - (n - 1 - i)) * v % MOD) % MOD;
    }
    for (ll i = 0; i < m; i++) {
        scanf("%lld", &v);
        y = (y + (i - (m - 1 - i)) * v % MOD) % MOD;
    }
    printf("%lld\n", x * y % MOD);
    return 0;
}
