#include <iostream>

using namespace std;
typedef long long ll;

const ll N = 3e5;
ll n, a[N + 10], inv[N + 10], S[N + 10];
char s[N + 10];

bool check() {
    for (ll i = 1; i <= n; i++) {
        ll j = inv[i];
        if (abs(S[j] - S[i]) < abs(j - i))
            return false;
    }
    return true;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 1; i <= n; i++)
        scanf("%lld", a + i), inv[a[i]] = i;
    scanf("%s", s);
    for (ll i = 1; i <= n; i++)
        S[i + 1] = S[i] + (s[i - 1] - '0');
    puts(check() ? "YES" : "NO");
    return 0;
}
