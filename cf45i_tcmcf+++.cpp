#include <iostream>
#include <climits>

using namespace std;
typedef long long ll;
const ll N = 100;
ll n, a[N + 10], pos, neg, z, m = LLONG_MIN;

int main() {
    scanf("%lld", &n); for (ll i = 0; i < n; i++)
        scanf("%lld", a + i), pos += a[i] > 0, neg += a[i] < 0, z += a[i] == 0, a[i] < 0 && (m = max(m, a[i]));
    if (!(pos || neg > 1))
        printf("%lld ", z ? 0 : m);
    else for (ll i = 0; i < n; i++)
        if (a[i]) {
            if (neg % 2 && a[i] == m)
                neg--;
            else
                printf("%lld ", a[i]);
        }
    return 0;
}
