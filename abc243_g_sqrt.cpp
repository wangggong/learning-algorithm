// 参考: https://atcoder.jp/contests/abc243/submissions/45083883
//
// `O(n^1/4)` 的做法, 避免在 `9e18` 的数据量下超时.
#include <iostream>
#include <cmath>

using namespace std;
typedef unsigned long long llu;
const llu M = 54772;    // sqrt(sqrt(9e18));
llu t, x, dp[M + 10];

llu get(llu n) {
    llu ans = 0, k = sqrt((long double)n);
    if (k * k > n)
        k--;
    for (llu i = 1; i * i <= k; i++)
        ans += dp[i] * (k - i * i + 1);
    return ans;
}

int main() {
    dp[1] = 1;
    for (llu i = 2, rt = 2; i <= M; i++) {
        dp[i] = dp[i - 1];
        if (rt * rt == i)
            dp[i] += dp[rt], rt++;
    }
    cin >> t;
    while (t--) {
        cin >> x;
        cout << get(x) << endl;
    }
    return 0;
}
