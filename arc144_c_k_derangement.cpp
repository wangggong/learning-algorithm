#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 3e5;
ll n, k, ans[N + 10];
bool st[N + 10];

int main() {
    scanf("%lld %lld", &n, &k);
    if (2 * k > n) {
        puts("-1");
        return 0;
    }
    ll m = 1;
    for (ll i = 1; i <= n; i++) {
        ll j = i + k;
        if (j > n || st[j] || j + k <= n) {
            j = m;
            if (abs(i - j) < k)
                j = i + k;
        }
        ans[i] = j;
        st[j] = true;
        for (; m <= n && st[m]; m++);
    }
    for (ll i = 1; i <= n; i++)
        printf("%lld ", ans[i]);
    return 0;
}
