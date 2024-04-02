#include <iostream>
#include <algorithm>

using namespace std;
typedef long long ll;
const ll N = 50;
ll n, m, a[N + 10], b[N + 10];

int main() {
    scanf("%lld%lld", &n, &m);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    for (ll i = 0; i < m; i++)
        scanf("%lld", b + i);
    sort(a, a + n), sort(b, b + m);
    printf("%lld", min(
        max(max(a[1] * b[0], a[1] * b[m - 1]), max(a[n - 1] * b[0], a[n - 1] * b[m - 1])),
        max(max(a[0] * b[0], a[0] * b[m - 1]), max(a[n - 2] * b[0], a[n - 2] * b[m - 1]))));
    return 0;
}
