#include <iostream>
#include <vector>
#define __FAST__ ios::sync_with_stdio(false); cin.tie(NULL); cout.tie(NULL);

using namespace std;
typedef long long ll;
ll n, tot;
string s;
vector<string> vec;

int main() {
    __FAST__;
    cin >> n;
    for (ll i = 0; i < n; i++)
        cin >> s, vec.push_back(s);
    for (ll i = n - 2; i >= 0; i--) {
        ll j = 0, m = min(vec[i].size(), vec[i + 1].size());
        for (; j < m && vec[i][j] == vec[i + 1][j]; j++);
        if (j < m && vec[i][j] < vec[i + 1][j])
            continue;
        tot += m - j;
        vec[i] = vec[i].substr(0, j);
    }
    for (auto s : vec)
        cout << s << endl;
    return 0;
}
