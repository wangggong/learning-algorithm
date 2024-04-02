#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 1e5;
const ll MOD = 1e9 + 7;
ll n, t[N + 10], a[N + 10];

ll getCount() {
    ll ans = 1;
    for (ll i = 1; i <= n; i++) {
        if (t[i] > t[i - 1]) {
            if (t[i] > a[i])
                return 0;
        }
        if (a[i] > a[i + 1]) {
            if (a[i] > t[i])
                return 0;
        }
        if (t[i] == t[i - 1] && a[i] == a[i + 1])
            ans = ans * min(t[i], a[i]) % MOD;
    }
    return ans;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 1; i <= n; i++)
        scanf("%lld", t + i);
    for (ll i = 1; i <= n; i++)
        scanf("%lld", a + i);
    printf("%lld", getCount());
    return 0;
}
