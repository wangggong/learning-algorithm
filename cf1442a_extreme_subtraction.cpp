#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 1e6;
ll t, n, a[N + 10];

int main() {
    scanf("%lld", &t);
    while (t--) {
        scanf("%lld", &n);
        for (ll i = 0; i < n; i++)
            scanf("%lld", a + i);
        ll tot = 0;
        for (ll i = 1; i < n; i++)
            if (a[i - 1] > a[i])
                tot += a[i - 1] - a[i];
        puts(tot <= a[0] ? "YES" : "NO");
    }
    return 0;
}
