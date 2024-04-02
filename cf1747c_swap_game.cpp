#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 1e9 + 7;
ll t, n, a, b, c;

int main() {
    for (scanf("%lld", &t); t--; ) {
        c = N;
        scanf("%lld", &n);
        for (ll i = 0; i < n; i++) {
            scanf("%lld", &a);
            if (!i)
                b = a;
            else
                c = min(c, a);
        }
        puts(b > c ? "Alice" : "Bob");
    }
    return 0;
}
