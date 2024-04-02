#include <iostream>

using namespace std;
typedef long long ll;
const ll N = 5000;
ll n, a[N + 10], b[N + 10], c[N + 10], va, vb;
bool duplicate;

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i), c[a[i]]++, duplicate = c[a[i]] > 1 ? true : duplicate;
    for (ll i = 0; i < n; i++)
        scanf("%lld", b + i), c[b[i]]--;
    for (ll i = 0; i < n; i++)
        for (ll j = i + 1; j < n; j++)
            va += a[i] > a[j], vb += b[i] > b[j];
    for (ll i = 0; i <= N; i++)
        if (c[i])
            goto FAIL;
    if (duplicate)
        goto SUCC;
    if ((va - vb) % 2)
        goto FAIL;
SUCC:
    puts("Yes");
    return 0;
FAIL:
    puts("No");
    return 0;
}
