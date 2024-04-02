#include <iostream>
#include <cstring>

using namespace std;

typedef long long ll;
const ll N = 18;
ll k;
char a[N + 10];

int main() {
    scanf("%s %lld", a, &k);
    ll n = strlen(a);
    for (ll i = 0; k > 0 && i < n; i++) {
        ll mx = i;
        for (ll j = 0; j <= k && i + j < n; j++)
            if (a[i + j] > a[mx])
                mx = i + j;
        for (; mx > i; mx--, k--)
            swap(a[mx], a[mx - 1]);
    }
    printf("%s\n", a);
    return 0;
}
