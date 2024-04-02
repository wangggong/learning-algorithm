#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll n, a[N + 10], b[N + 10], tr[N + 10], ans;

inline ll lowBit(ll x) { return x & -x; }

ll query(ll k) {
    ll ans = 0;
    for (; k > 0; k -= lowBit(k))
        ans += tr[k];
    return ans;
}

void add(ll k) {
    for (k++; k <= n; k += lowBit(k))
        tr[k]++;
    return;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 0; i < n; i++)
        scanf("%lld", a + i), a[i] += i;
    for (ll i = 0; i < n; i++)
        scanf("%lld", b + i), b[i] += i;
    unordered_map<ll, vector<ll>> M;
    for (ll i = n - 1; i >= 0; i--)
        M[a[i]].push_back(i);
    for (ll i = 0; i < n; i++) {
        if (M[b[i]].size() == 0)
            goto FAIL;
        ll j = M[b[i]].back();
        M[b[i]].pop_back();
        ans += i - query(j);
        add(j);
    }
    printf("%lld\n", ans);
    return 0;
FAIL:
    puts("-1");
    return 0;
}
