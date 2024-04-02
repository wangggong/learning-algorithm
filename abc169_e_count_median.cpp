#include <iostream>
#include <algorithm>

using namespace std;
typedef long long ll;

const ll N = 2e5;
ll n, a[N + 10], b[N + 10];

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld %lld", a + i, b + i);
    sort(a, a + n);
    sort(b, b + n);
    ll mx = b[n / 2] + (n % 2 ? 0 : b[n / 2 - 1]), mn = a[n / 2] + (n % 2 ? 0 : a[n / 2 - 1]);
    printf("%lld\n", mx - mn + 1);
    return 0;
}
