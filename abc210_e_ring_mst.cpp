#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
ll n, m, a, c, ans;

ll gcd(ll x, ll y) { return y == 0 ? x : gcd(y, x % y); }

int main() {
    scanf("%lld %lld", &n, &m);
    vector<pll> vec;
    // 按照花费排序. 对于每个操作, 执行一次会减少一个联通块; 也就是说, 执行次数是联通块数目的差.
    for (ll i = 0; i < m; i++)
        scanf("%lld %lld", &a, &c), vec.push_back({c, a});
    sort(vec.begin(), vec.end());
    for (auto v : vec) {
        if (n == 1)
            break;
        ll g = gcd(n, v.second);
        ans += (n - g) * v.first;
        n = g;
    }
    printf("%lld\n", n == 1 ? ans : -1);
    return 0;
}
