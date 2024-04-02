#include <iostream>
#include <cstring>
#include <unordered_map>

using namespace std;
typedef long long ll;
const ll N = 2e5, S = 5e6, A = 26;
ll n, c[N + 10], e[N + 10], ans;
char s[S + 10];

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++) {
        scanf("%s", s);
        for (ll j = strlen(s) - 1; j >= 0; j--)
            e[i] |= 1ll << (s[j] - 'a'), c[i] ^= 1ll << (s[j] - 'a');
    }
    for (ll a = 0; a < A; a++) {
        ll target = ((1ll << A) - 1ll) ^ (1ll << a);
        unordered_map<ll, ll> m;
        for (ll i = 0; i < n; i++) {
            if (e[i] & (1ll << a))
                continue;
            m[c[i]]++;
            ans += m[target ^ c[i]];
        }
    }
    printf("%lld\n", ans);
    return 0;
}
