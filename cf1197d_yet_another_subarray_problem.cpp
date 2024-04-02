#include <iostream>

using namespace std;
typedef long long ll;

const ll M = 10, MAXN = 1e18;
ll n, m, k, x, tot, minS[M + 10], ans;

int main() {
    cin >> n >> m >> k;
    for (ll i = 1; i < m; i++)
        minS[i] = MAXN;
    for (ll i = 1; i <= n; i++) {
        cin >> x;
        tot += x * m - k;
        for (ll d = 0; d < m; d++)
            ans = max(ans, tot - minS[(i + d) % m] - k * d);
        minS[i % m] = min(minS[i % m], tot);
    }
    cout << (ans / m) << endl;
    return 0;
}
