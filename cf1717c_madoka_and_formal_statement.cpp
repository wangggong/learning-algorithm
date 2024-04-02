// 参考: https://codeforces.com/contest/1717/submission/243855296
#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll t, n, a[N + 10], b[N + 10];

bool check() {
    for (ll i = 0; i < n; i++)
        if (a[i] > b[i] || (a[i] < b[i] && b[i] > b[(i + 1) % n] + 1))
            return false;
    return true;
}

int main() {
    for (scanf("%lld", &t); t--; ) {
        scanf("%lld", &n);
        for (ll i = 0; i < n; i++)
            scanf("%lld", a + i);
        for (ll i = 0; i < n; i++)
            scanf("%lld", b + i);
        puts(check() ? "YES" : "NO");
    }
    return 0;
}
