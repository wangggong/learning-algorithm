#include <iostream>
#include <cstring>

using namespace std;
typedef long long ll;
const ll N = 1e5, A = 26;
char s[N + 10];
ll n, counts[A];
bool visits[A];

bool check(ll k) {
    memset(counts, 0, sizeof counts);
    memset(visits, 0, sizeof visits);
    for (ll i = 0; i < k; i++)
        counts[s[i] - 'a']++;
    for (ll i = 0; i < A; i++)
        if (counts[i])
            visits[i] = true;
    for (ll i = k; i < n; i++) {
        counts[s[i] - 'a']++, counts[s[i - k] - 'a']--;
        if (!counts[s[i - k] - 'a'])
            visits[s[i - k] - 'a'] = false;
    }
    for (ll i = 0; i < A; i++)
        if (visits[i])
            return true;
    return false;
}

int main() {
    scanf("%s", s);
    n = strlen(s);
    ll p = 0, q = n;
    while (p < q)
    {
        ll mid = p + q >> 1;
        if (check(mid))
            q = mid;
        else
            p = mid + 1;
    }
    printf("%lld\n", p);
    return 0;
}
