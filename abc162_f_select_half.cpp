#include <iostream>

using namespace std;
typedef long long ll;

ll n, a, S, dp[3];

int main() {
    scanf("%lld %lld", &n, &S);
    for (ll i = 2; i <= n; i++) {
        scanf("%lld", &a);
        dp[i % 3] = max(dp[(i - 2) % 3] + a, i % 2 ? dp[(i - 1) % 3] : S), S += (i % 2) * a;
    }
    cout << dp[n % 3] << endl;
    return 0;
}
