#include <iostream>

using namespace std;

typedef long long ll;
const ll mod = 1e9 + 7;

ll dp[3], cnt[3];

int main() {
    int n, l, r;
    cin >> n >> l >> r;
    l--;
    cnt[0] = r / 3 - l / 3, cnt[1] = (r + 2) / 3 - (l + 2) / 3, cnt[2] = (r + 1) / 3 - (l + 1) / 3;
    dp[0] = cnt[0], dp[1] = cnt[1], dp[2] = cnt[2];
    for (int i = 2; i <= n; i++) {
        ll last[3] = {dp[0], dp[1], dp[2]};
        dp[0] = (last[0] * cnt[0] + last[1] * cnt[2] + last[2] * cnt[1]) % mod;
        dp[1] = (last[0] * cnt[1] + last[1] * cnt[0] + last[2] * cnt[2]) % mod;
        dp[2] = (last[0] * cnt[2] + last[1] * cnt[1] + last[2] * cnt[0]) % mod;
    }
    cout << dp[0] << endl;
    return 0;
}
