/*
 * https://codeforces.com/problemset/submission/573/164666280
 * 
 * 提示 1-1：把关注点放在每列积木被清除的时刻上。答案为所有积木列被清除时刻的最大值。
 * 
 * 提示 1-2：递推。
 * 
 * 提示 2-1：如果在某个时刻第 i-1 列或第 i+1 列积木被清除了，那么下一时刻第 i 列积木就被清除了。另外，一列积木被清除的时间不会超过这列积木的高度。
 * 
 * 提示 2-2：从左到右扫一遍，从右到左扫一遍。
 */
#include <iostream>

using namespace std;
const int N = 1e5;

int h[N + 10], dp[N + 10], n, ans;

int main() {
    scanf("%d", &n);
    for (int i = 0; i < n; i++)
        scanf("%d", h + i), dp[i] = h[i];
    dp[0] = dp[n - 1] = 1;
    for (int i = 1; i < n; i++)
        dp[i] = min(dp[i], dp[i - 1] + 1);
    for (int i = n - 2; i >= 0; i--)
        dp[i] = min(dp[i], dp[i + 1] + 1);
    for (int i = 0; i < n; i++)
        ans = max(ans, dp[i]);
    printf("%d\n", ans);
    return 0;
}
