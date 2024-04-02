#include <iostream>
#include <vector>

using namespace std;
typedef long long ll;

const ll N = 2e5;
ll n, w[N + 10], u, v, ans[N + 10];
vector<ll> e[N + 10];

void dfs(int root, int pa, vector<int> &vec) {
    ll p = 0, q = vec.size();
    while (p < q) {
        ll mid = p + q >> 1;
        if (w[vec[mid]] >= w[root])
            q = mid;
        else
            p = mid + 1;
    }
    // ans[root] = p + 1;
    if (p < vec.size()) {
        ll last = vec[p];
        vec[p] = root;
        ans[root] = vec.size();
        for (int ch : e[root])
            if (ch != pa)
                dfs(ch, root, vec);
        vec[p] = last;
    } else {
        vec.push_back(root);
        ans[root] = vec.size();
        for (int ch : e[root])
            if (ch != pa)
                dfs(ch, root, vec);
        vec.pop_back();
    }
    return;
}

int main() {
    scanf("%lld", &n);
    for (ll i = 1; i <= n; i++)
        scanf("%lld", w + i);
    for (ll i = 1; i < n; i++) {
        scanf("%lld %lld", &u, &v);
        e[u].push_back(v);
        e[v].push_back(u);
    }
    vector<int> dis;
    dfs(1, -1, dis);
    for (ll i = 1; i <= n; i++)
        printf("%lld\n", ans[i]);
    return 0;
}
