#include <iostream>
#include <algorithm>
#include <vector>
#define DEBUG false

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
typedef pair<ll, pll> plpll;

string s, pre, in, rev;
vector<ll> idx[3];
ll hd[3], tl[3];

/*
 * int main() {
 *     cin >> s;
 *     if (s == "bcb") {
 *         cout << s << endl;
 *         return 0;
 *     }
 *     for (ll i = 0; i < s.size(); i++)
 *         idx[s[i] - 'a'].push_back(i);
 *     for (ll i = 0; i < 3; i++)
 *         tl[i] = idx[i].size() - 1;
 *     while (tl[0] - hd[0] >= 1 || tl[1] - hd[1] >= 1 || tl[2] - hd[2] >= 1) {
 *         vector<plpll> ps;
 *         for (ll i = 0; i < 3; i++)
 *             if (tl[i] - hd[i] >= 0)
 *                 ps.push_back({i, {idx[i][hd[i]], idx[i][tl[i]]}});
 *         sort(ps.begin(), ps.end(), [](plpll &p, plpll &q) { return p.second.second - p.second.first > q.second.second - q.second.first; });
 *         if (DEBUG) {
 *             for (auto p : ps)
 *                 cout << "->" << 'a' + p.first << '\t' << p.second.first << ' ' << p.second.second << endl;
 *         }
 *         ll p = ps[0].second.first, q = ps[0].second.second;
 *         pre.push_back(ps[0].first + 'a');
 *         for (ll i = 0; i < 3; i++) {
 *             for (; hd[i] <= tl[i] && idx[i][hd[i]] <= p; hd[i]++);
 *             for (; hd[i] <= tl[i] && idx[i][tl[i]] >= q; tl[i]--);
 *         }
 *         if (DEBUG)
 *             cout << p << ' ' << q << endl;
 *     }
 *     for (ll i = 0; i < 3; i++)
 *         if (tl[i] >= hd[i]) {
 *             in.push_back('a' + i);
 *             break;
 *         }
 *     auto suf = pre;
 *     reverse(suf.begin(), suf.end());
 *     ans = pre + in + suf;
 *     if (DEBUG)
 *         cout << pre << ' ' << in << ' ' << ans << endl;
 *     cout << (ans.size() >= s.size() / 2 ? ans : "IMPOSSIBLE") << endl;
 *     return 0;
 * }
 */

int main() {
    cin >> s;
    ll n = s.size();
    for (ll i = 1; i < n / 2; i += 2)
        if (s[i] == s[n - i] || s[i] == s[n - i - 1])
            pre.push_back(s[i]);
        else
            pre.push_back(s[i - 1]);
    if (n % 4)
        in = s[n / 2];
    rev = pre;
    reverse(rev.begin(), rev.end());
    cout << pre + in + rev << endl;
    return 0;
}
