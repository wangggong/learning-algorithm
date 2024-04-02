#include <iostream>
#include <map>

using namespace std;
typedef long long ll;
typedef pair<ll, ll> pll;
ll n, x, y, ans;
map<ll, ll> mx, my;
map<pll, ll> m;

int main() {
    for (scanf("%lld", &n); n--; )
        scanf("%lld%lld", &x, &y), ans += mx[x] + my[y] - m[{x, y}], mx[x]++, my[y]++, m[{x, y}]++;
    printf("%lld", ans);
    return 0;
}
