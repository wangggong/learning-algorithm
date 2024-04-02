#include <iostream>

using namespace std;
typedef long long ll;

const ll N = 1e5;
ll a[N + 10], n, b[N + 10];

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i);
    a[n] = a[n - 1];
    for (ll i = 0; i < n; ) {
        if (a[i] == a[i + 1]) {
            if (b[i] == 0)
                b[i] = i > 0 && b[i - 1] == 2 ? 3 : 2;
            i++;
            continue;
        }
        ll st = i;
        for (i += 2; i < n && a[i] != a[i - 1] && (a[i] > a[i - 1]) == (a[i - 1] > a[i - 2]); i++);
        if (a[st] > a[st + 1]) {
            b[st] = st > 0 && b[st - 1] == 5 ? 4 : 5;
            for (st++; st < i; st++)
                b[st] = b[st - 1] - 1;
        } else {
            b[st] = st > 0 && b[st - 1] == 1 ? 2 : 1;
            for (st++; st < i; st++)
                b[st] = b[st - 1] + 1;
        }
        i--;
        if (b[i] < 1 || b[i] > 5) {
            puts("-1");
            return 0;
        }

    }
    for (ll i = 0; i < n; i++)
        printf("%lld ", b[i]);
    return 0;
}
