#include <iostream>
#include <vector>
#include <algorithm>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)

using namespace std;
typedef long long ll;
ll n;
string s;
vector<string> vec;

ll get(string s) {
    ll tot = 0, cur = 0;
    for (auto ch : s)
        if (ch == 'X')
            cur++;
        else
            tot += (ch - '0') * cur;
    return tot;
}

int main() {
    __FAST__;
    cin >> n;
    for (ll i = 0; i < n; i++)
        cin >> s, vec.push_back(s);
    sort(vec.begin(), vec.end(), [](string &p, string &q) { return get(p + q) > get(q + p); });
    ll tot = 0, cur = 0;
    for (auto s : vec)
        for (auto ch : s)
            if (ch == 'X')
                cur++;
            else
                tot += (ch - '0') * cur;
    cout << tot << endl;
    return 0;
}
