#include <iostream>
#include <vector>
#define DEBUG false
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)

using namespace std;
typedef long long ll;

const ll N = 1e6;
ll n, fa[N + 10], w[N + 10], target, root, ans[2];
vector<ll> sons[N + 10];

ll dfs(int k) {
    ll tot = w[k];
    for (ll s : sons[k]) {
        ll cur = dfs(s);
        if (fa[s] == k)
            tot += cur;
    }
    if (tot == target) {
        if (ans[0] == 0)
            ans[0] = k;
        else if (ans[1] == 0)
            ans[1] = k;
        else;
        fa[k] = 0;
    }
    if (DEBUG)
        cout << '\t' << k << ' ' << tot << endl;
    return tot;
}

int main() {
    __FAST__;
    cin >> n;
    for (ll i = 1; i <= n; i++)
        cin >> fa[i] >> w[i], root = fa[i] == 0 ? i : root, sons[fa[i]].push_back(i), target += w[i];
    if (target % 3) {
        cout << -1 << endl;
        return 0;
    }
    target /= 3;
    dfs(root);
    if (ans[0] && ans[1] && ans[0] != root && ans[1] != root)
        cout << ans[0] << ' ' << ans[1] << endl;
    else
        cout << -1 << endl;
    return 0;
}
