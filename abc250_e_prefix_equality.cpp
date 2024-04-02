#include <iostream>
#include <unordered_map>
#include <unordered_set>
#define YES "Yes"
#define NO "No"

using namespace std;
typedef long long ll;

const ll N = 2e5, MAXN = 1e9;
ll n, v, q, x, y, a[N + 10], b[N + 10];

int main() {
    unordered_map<ll, ll> M;
    scanf("%lld", &n);
    for (ll i = 1; i <= n; i++) {
        scanf("%lld", &v);
        if (!M.count(v))
            M[v] = M.size() + 1;
        a[i] = max(a[i - 1], M[v]);
    }
    unordered_set<ll> S;
    for (ll i = 1, mx = 0; i <= n; i++) {
        scanf("%lld", &v);
        v = M[v];
        mx = v ? max(mx, v) : MAXN;
        S.insert(v);
        if (mx == S.size())
            b[i] = mx;
    }
    scanf("%lld", &q);
    while (q--) {
        scanf("%lld %lld", &x, &y);
        puts(a[x] == b[y] ? YES : NO);
    }
    return 0;
}
