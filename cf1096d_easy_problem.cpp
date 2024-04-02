#include <iostream>

using namespace std;
typedef long long ll;

const int N = 1e5;
int n, a[N + 10];
string s, sh = "hard";
ll dp[N + 10][5];

/*
 * int main() {
 *     cin >> n >> s;
 *     for (int i = 0; i < n; i++)
 *         cin >> a[i];
 *     for (int i = 0; i <= n; i++)
 *         dp[i][0] = LLONG_MAX;
 *     for (int i = 1; i <= n; i++)
 *         for (int j = 1; j <= 5; j++)
 *             // `dp[i][j]`: `s[:i]` 中删除字母使得不存在 `sh[:j]` 的最小代价
 *             // 那么当 `s[i - 1] == sh[j - 1]` 时:
 *             // - 删: 允许有 `sh[:j]` 的前缀, 代价为 `dp[i - 1][j] + a[i - 1]`
 *             // - 不删: 允许有 `sh[:j-1]` 的前缀, 代价为 `dp[i - 1][j - 1]`
 *             // 初始状态: 设为系统最大.
 *             dp[i][j] = s[i - 1] == sh[j - 1] ? min(dp[i - 1][j - 1], dp[i - 1][j] + a[i - 1]) : dp[i - 1][j];
 *     cout << dp[n][4] << endl;
 *     return 0;
 * }
 */

int main() {
    cin >> n >> s;
    for (int i = 0; i < n; i++)
        cin >> a[i];
    for (int i = 1; i <= n; i++)
        dp[i][0] = s[i - 1] == sh[0] ? dp[i - 1][0] + a[i - 1] : dp[i - 1][0];
    for (int i = 1; i <= n; i++)
        for (int j = 1; j < 4; j++)
            dp[i][j] = s[i - 1] == sh[j] ? min(dp[i - 1][j] + a[i - 1], dp[i - 1][j - 1]) : dp[i - 1][j];
    cout << dp[n][3] << endl;
    return 0;
}
