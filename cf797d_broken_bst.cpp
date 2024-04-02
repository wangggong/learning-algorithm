#include <iostream>
#include <unordered_map>

using namespace std;
typedef long long ll;
const ll N = 1e5, MAXN = 1e9 + 1;
ll n, a[N + 10][3], root, ans;
bool hasPa[N + 10];

void preorder(ll node, ll l, ll r, unordered_map<ll, ll> &m) {
    if (l <= a[node][0] && a[node][0] <= r)
        ans -= m[a[node][0]];
    if (a[node][1] != -1)
        preorder(a[node][1], l, min(a[node][0] - 1, r), m);
    if (a[node][2] != -1)
        preorder(a[node][2], max(l, a[node][0] + 1), r, m);
    return;
}

int main() {
    unordered_map<ll, ll> count;
    scanf("%lld", &n);
    for (ll i = 1; i <= n; i++)
        scanf("%lld %lld %lld", a[i], a[i] + 1, a[i] + 2), count[a[i][0]]++, hasPa[a[i][1]] = true, hasPa[a[i][2]] = true;
    for (root = 1; hasPa[root]; root++);
    ans = n;
    preorder(root, 0, MAXN, count);
    printf("%lld\n", ans);
    return 0;
}
