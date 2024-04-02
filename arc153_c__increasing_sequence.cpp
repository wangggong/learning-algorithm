#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 2e5, M = 1e12;
ll n, a[N + 10], b[N + 10], S, prefix, suffix;

bool check() {
    for (ll i = 0; i < n; i++)
        b[i] = i, S += a[i] * b[i];
    if (!S)
        return true;
    if ((S > 0) == (a[0] > 0)) {
        b[0] = -S / a[0];
        return true;
    }
    if ((S > 0) != (a[n - 1] > 0)) {
        b[n - 1] = -(S - a[n - 1] * (n - 1)) / a[n - 1];
        return true;
    }
    for (ll i = 0; i < n; i++) {
        prefix += a[i];
        if (prefix && ((prefix > 0) == (a[n - 1] > 0))) {
            S = 0;
            for (ll j = 0; j < n; j++)
                b[j] = j == i ? 0 : (j > i ? j : j - M), S += a[j] * b[j];
            b[i] = -S / a[i];
            return true;
        }
    }
    for (ll i = n - 1; i >= 0; i--) {
        suffix += a[i];
        if (suffix && ((suffix > 0) == (a[0] > 0))) {
            S = 0;
            for (ll j = 0; j < n; j++)
                b[j] = j == i ? 0 : (j < i ? j : M - n + j), S += a[j] * b[j];
            b[i] = -S / a[i];
            return true;
        }
    }
    return false;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    if (check()) {
        puts("Yes");
        for (ll i = 0; i < n; i++)
            printf("%lld ", b[i]);
    } else
        puts("No");
    return 0;
}
