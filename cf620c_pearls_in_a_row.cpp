#include <iostream>
#include <vector>
#include <set>

using namespace std;

typedef long long ll;
typedef pair<ll, ll> pll;
const ll N = 3e5;
ll A[N + 10], n, _right;

int main() {
    scanf("%lld", &n);
    for (ll i = 1; i <= n; i++)
        scanf("%lld", A + i);
    vector<pll> vec;
    set<ll> S;
    for (_right = n; _right > 0 && !S.count(A[_right]); _right--)
        S.insert(A[_right]);
    if (_right == 0) {
        puts("-1");
        return 0;
    }
    for (ll i = 1; i < _right; i++) {
        S.clear();
        ll start = i;
        for (; i < _right && !S.count(A[i]); i++)
            S.insert(A[i]);
        vec.push_back({start, i == _right ? n : i});
    }
    if (vec.size() == 0 || vec.rbegin()->second <= _right)
        vec.push_back({_right, n});
    printf("%lld\n", (ll)vec.size());
    for (auto [l, r] : vec)
        printf("%lld %lld\n", l, r);
    return 0;
}
