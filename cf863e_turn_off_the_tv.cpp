#include <iostream>
#include <algorithm>
#include <vector>
#define __FAST__ ios::sync_with_stdio(false), cin.tie(NULL), cout.tie(NULL)

using namespace std;
typedef long long ll;

struct I {
    ll s, e, i;
};
ll n, s, e;

int main() {
    __FAST__;
    vector<I> vec;
    cin >> n;
    for (ll i = 1; i <= n; i++)
        cin >> s >> e, vec.push_back({s, e, i});
    sort(vec.begin(), vec.end(), [](I &p, I &q) { return p.s < q.s || (p.s == q.s && p.e > q.e); });
    for (ll i = 1; i < n; i++)
        if (vec[i].e <= vec[i - 1].e) {
            cout << vec[i].i << endl;
            return 0;
        }
    for (ll i = 1; i + 1 < n; i++)
        if (vec[i + 1].s <= vec[i - 1].e + 1) {
            cout << vec[i].i << endl;
            return 0;
        }
    cout << -1 << endl;
    return 0;
}
