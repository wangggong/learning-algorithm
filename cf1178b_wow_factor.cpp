#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 1e6;
ll n;
char s[N + 10];

int main() {
    scanf("%s", s);
    n = strlen(s);
    ll ans = 0, cur = 0, total = 0;
    for (ll p = 0, q = 0; p < n; p = q) {
        for (; p < n && s[p] != 'v'; p++);
        if (p == n)
            break;
        for (q = p; q < n && s[q] == 'v'; q++);
        total += q - p - 1;
    }
    for (ll p = 0, q = 0; p < n; p = q) {
        for (; p < n && s[p] != 'v'; p++)
            ans += cur * (total - cur);
        if (p == n)
            break;
        for (q = p; q < n && s[q] == 'v'; q++);
        cur += q - p - 1;
    }
    printf("%lld", ans);
    return 0;
}
