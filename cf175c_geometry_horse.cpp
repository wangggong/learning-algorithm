#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;

int main() {
    ll n, s, c, t, p, tot = 0;
    ll ans = 0;
    vector<pll> arr;
    vector<ll> count;
    cin >> n;
    while (n--) {
        cin >> c >> s;
        arr.push_back({s, c});
    }
    sort(arr.begin(), arr.end());
    cin >> t;
    while (t--) {
        cin >> p;
        count.push_back(p);
    }
    count.push_back(LLONG_MAX);
    for (ll i = 0, j = 0; i < arr.size(); i++) {
        for (ll s = arr[i].first, r = arr[i].second; r; ) {
            ll k = min(r, count[j] - tot);
            ans += (j + 1) * s * k;
            tot += k, r -= k;
            if (tot >= count[j])
                j++;
        }
    }
    cout << ans << endl;
    return 0;
}
