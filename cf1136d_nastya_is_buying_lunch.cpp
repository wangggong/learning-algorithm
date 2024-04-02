#include <iostream>
#include <vector>
#include <unordered_set>

using namespace std;
typedef pair<int, int> pii;

int main() {
    int n, m, x, u, v;
    cin >> n >> m;
    vector<int> p;
    for (int i = 0; i < n; i++) {
        cin >> x;
        p.push_back(x);
    }
    vector<vector<int>> vec(n + 1, vector<int>());
    for (int i = 0; i < m; i++) {
        cin >> u >> v;
        vec[u].push_back(v);
    }
    unordered_set<int> S;
    S.insert(p[n - 1]);
    for (int i = n - 2; i >= 0; i--) {
        int cnt = 0;
        for (int v : vec[p[i]])
            if (S.count(v))
                cnt++;
        if (cnt != S.size())
            S.insert(p[i]);
    }
    cout << n - S.size() << endl;
    return 0;
}
