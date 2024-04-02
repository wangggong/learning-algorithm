#include <iostream>

using namespace std;
typedef long long ll;
ll n, a, b, i;

int main() {
    scanf("%lld %lld %lld", &n, &a, &b);
    if (a * b < n || a + b > n + 1) {
        puts("-1");
        return 0;
    }
    while (a--) {
        ll sz = min(b, n - a);
        for (ll j = i + sz; j > i; j--)
            printf("%lld ", j);
        i += sz;
        n -= sz;
    }
    return 0;
}
