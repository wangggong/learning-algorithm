#include <iostream>
#include <algorithm>
#include <queue>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
ll T, n, l, r;

bool get(vector<pll> &vec) {
    sort(vec.begin(), vec.end());
    priority_queue<ll, vector<ll>, greater<ll>> pq;
    for (ll i = 0, left = 1; i < n || !pq.empty(); ) {
        for (; i < n && vec[i].first == left; i++)
            pq.push(vec[i].second);
        if (pq.empty()) {
            left = vec[i].first;
            continue;
        }
        auto q = pq.top(); pq.pop();
        if (q < left)
            return false;
        left++;
    }
    return true;
}

int main() {
    scanf("%lld", &T);
    while (T--) {
        scanf("%lld", &n);
        vector<pll> vec;
        for (auto i = 0; i < n; i++)
            scanf("%lld %lld", &l, &r), vec.push_back({l, r});
        puts(get(vec) ? "Yes" : "No");
    }
    return 0;
}
