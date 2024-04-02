/*
 * @lc app=leetcode.cn id=minimum-number-of-operations-to-make-x-and-y-equal lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100159] 使 X 和 Y 相等的最少操作次数
 *
 * https://leetcode.cn/problems/minimum-number-of-operations-to-make-x-and-y-equal/description/
 *
 * algorithms
 * Medium (38.92%)
 * Total Accepted:    1.6K
 * Total Submissions: 4.1K
 * Testcase Example:  '26\n1'
 *
 * 给你两个正整数 x 和 y 。
 * 
 * 一次操作中，你可以执行以下四种操作之一：
 * 
 * 
 * 如果 x 是 11 的倍数，将 x 除以 11 。
 * 如果 x 是 5 的倍数，将 x 除以 5 。
 * 将 x 减 1 。
 * 将 x 加 1 。
 * 
 * 
 * 请你返回让 x 和 y 相等的 最少 操作次数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：x = 26, y = 1
 * 输出：3
 * 解释：我们可以通过以下操作将 26 变为 1 ：
 * 1. 将 x 减 1
 * 2. 将 x 除以 5
 * 3. 将 x 除以 5
 * 将 26 变为 1 最少需要 3 次操作。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：x = 54, y = 2
 * 输出：4
 * 解释：我们可以通过以下操作将 54 变为 2 ：
 * 1. 将 x 加 1
 * 2. 将 x 除以 11
 * 3. 将 x 除以 5
 * 4. 将 x 加 1
 * 将 54 变为 2 最少需要 4 次操作。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：x = 25, y = 30
 * 输出：5
 * 解释：我们可以通过以下操作将 25 变为 30 ：
 * 1. 将 x 加 1
 * 2. 将 x 加 1
 * 3. 将 x 加 1
 * 4. 将 x 加 1
 * 5. 将 x 加 1
 * 将 25 变为 30 最少需要 5 次操作。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= x, y <= 10^4
 * 
 * 
 */
public class Solution
{
    public int MinimumOperationsToMakeEqual(int x, int y)
    {
        const int M = 11;
        const int N = 5;
        var memo = new int[x + M + 1];
        Array.Fill(memo, -1);
        int dfs(int x) => memo[x] = (memo[x] is not -1 ? memo[x] : (
            x <= y ? y - x : new int[]
            {
                x - y,
                x % M + 1 + dfs(x / M),
                x % N + 1 + dfs(x / N),
                // (x / M + 1) * M - x + 1 + dfs(x / M + 1),
                // (x / N + 1) * N - x + 1 + dfs(x / N + 1),
                M - x % M + 1 + dfs(x / M + 1),
                N - x % N + 1 + dfs(x / N + 1),
            }.Min()));
        return dfs(x);
    }
}
