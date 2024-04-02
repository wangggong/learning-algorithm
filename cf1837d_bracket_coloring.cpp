#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll t, m;
bool a[N + 10];
char s[N + 10];

bool canBeBeautiful() {
    memset(a, false, sizeof a);
    ll d = 0;
    bool last = true;
    for (ll i = 0; i < m; i++) {
        d += s[i] == '(' ? 1 : -1;
        if (d == -1)
            last = false;
        if (d == 1)
            last = true;
        if (d >= 0 && last)
            a[i] = true;
    }
    return d == 0;
}

bool canSingle() {
    for (ll i = 1; i < m; i++)
        if (a[i] != a[0])
            return false;
    memset(a, true, sizeof a);
    return true;
}

int main() {
    for (scanf("%lld", &t); t; t--) {
        scanf("%lld%s", &m, s);
        if (!canBeBeautiful()) {
            puts("-1");
            continue;
        }
        puts(canSingle() ? "1": "2");
        for (ll i = 0; i < m; i++)
            printf("%lld ", a[i] ? 1ll : 2ll);
        puts("");
    }
    return 0;
}
