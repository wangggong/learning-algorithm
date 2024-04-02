#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll n, a[N + 10], b[N + 10], cnt[N + 10];

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i), cnt[a[i]]++;
    for (ll i = 0; i < n; i++)
        scanf("%lld", b + (n - 1 - i)), cnt[b[n - 1 - i]]++;
    for (ll i = 1; i <= n; i++)
        if (cnt[i] > n) {
            puts("No");
            return 0;
        }
    puts("Yes");
    ll j = 0;
    for (ll i = 0; i < n; i++)
        if (a[i] == b[i]) {
            for (; j < n && (a[i] == b[j] || a[j] == b[i]); j++);
            swap(b[i], b[j]);
        }
    for (ll i = 0; i < n; i++)
        printf("%lld ", b[i]);
    return 0;
}
