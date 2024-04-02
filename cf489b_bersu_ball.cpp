#include <iostream>
#include <algorithm>

using namespace std;
typedef long long ll;
const ll N = 100;
ll n, m, a[N + 10], b[N + 10], ans;

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    scanf("%lld", &m);
    for (ll i = 0; i < m; i++)
        scanf("%lld", b + i);
    sort(a, a + n), sort(b, b + m);
    for (ll i = 0, j = 0; i < n; i++) {
        for (; j < m && b[j] < a[i] - 1; j++);
        if (j < m && b[j] <= a[i] + 1)
            ans++, j++;
    }
    printf("%lld\n", ans);
    return 0;
}
