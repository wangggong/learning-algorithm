#include <iostream>
#include <queue>
#include <climits>

using namespace std;
typedef long long ll;
const ll N = 2e5;
ll n, k, ops[N + 10][2], ans = LLONG_MIN, delta;

int main() {
    scanf("%lld %lld", &n, &k);
    ops[0][0] = 1;
    for (ll i = 1; i <= n; i++)
        scanf("%lld %lld", ops[i], ops[i] + 1);
    priority_queue<ll, vector<ll>, greater<ll>> Q;
    for (ll i = n; i >= 0; i--) {
        if (ops[i][0] != 2) {
            ans = max(ans, ops[i][1] + delta);
            if (!k--)
                break;
        } else if (ops[i][1] >= 0)
            delta += ops[i][1];
        else
            Q.push(-ops[i][1]);
        if (Q.size() > k)
            delta -= Q.top(), Q.pop();
    }
    printf("%lld\n", ans);
    return 0;
}
