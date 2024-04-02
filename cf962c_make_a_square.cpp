#include <iostream>
#include <queue>
#include <cmath>
#include <algorithm>

using namespace std;
typedef long long ll;
string s;

bool isSquare(string s) {
    ll v = stol(s);
    ll sq = sqrt(v);
    return sq > 0 && sq * sq == v;
}

ll bfs() {
    queue<string> Q;
    Q.push(s);
    for (ll i = 0; !Q.empty(); i++)
        for (ll c = Q.size(); c > 0; c--) {
            string q = Q.front(); Q.pop();
            if (q.size() > 1 && q[0] == '0')
                continue;
            if (isSquare(q))
                return i;
            ll m = q.size();
            if (m == 1)
                continue;
            for (ll j = 0; j < m; j++)
                Q.push(q.substr(0, j) + q.substr(j + 1, m - j - 1));
        }
    return -1;
}

int main() {
    cin >> s;
    cout << bfs() << endl;
    return 0;
}
