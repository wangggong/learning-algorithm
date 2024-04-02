// Problem: D. Book of Evil
// Contest: Codeforces - Codeforces Round 196 (Div. 2)
// URL: https://codeforces.com/problemset/problem/337/D
// Memory Limit: 256 MB
// Time Limit: 2000 ms
//
// Powered by CP Editor (https://cpeditor.org)

#include <iostream>
#include <vector>
#include <array>
#include <set>

using namespace std;

using ll = long long;
using pii = pair<int, int>;
using vl = vector<ll>;
using vi = vector<int>;
using a3i = array<int, 3>;

int n, m, dist;
vi g[100003];
set<int> gs;
vector<a3i> stat;
int ans = 0;

a3i dfs(int u, int fa)
{
    a3i ret{INT_MIN / 2, 0, INT_MIN / 2};
    if (gs.count(u))
        ret[0] = 0, ret[1] = u;
    for (auto v : g[u]) {
        if (v == fa)
            continue;
        a3i sub = dfs(v, u);
        if (sub[0] + 1 > ret[0]) {
            ret[2] = ret[0];
            ret[0] = sub[0] + 1;
            ret[1] = v;
        } else if (sub[0] > ret[2])
            ret[2] = sub[0] + 1;
    }
    return stat[u] = ret;
}
void reroot(int u, int fa, int d)
{
    if (d > dist)
        return;
    const a3i &s = stat[u];
    if (s[0] <= dist && d <= dist)
        ans++;
    for (auto v : g[u]) {
        if (v == fa)
            continue;
        if (v == s[1])
            reroot(v, u, max(d, s[2]) + 1);
        else
            reroot(v, u, max(d, s[0]) + 1);
    }
}

int main(int argc, char **argv)
{
    cin >> n >> m >> dist;
    int a, b;
    for (int i = 0; i < m; ++i)
        cin >> a, gs.insert(a);
    for (int i = 1; i < n; ++i)
        cin >> a >> b, g[a].push_back(b), g[b].push_back(a);
    stat = vector<a3i>(100003);
    dfs(1, -1);
    reroot(1, -1, INT_MIN / 2);
    for (ll i = 1; i <= n; i++)
        cout << stat[i][0] << ' ' << stat[i][2] << endl;
    cout << ans << endl;
    return 0;
};

