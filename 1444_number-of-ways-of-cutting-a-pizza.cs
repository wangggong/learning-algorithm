/*
 * @lc app=leetcode.cn id=number-of-ways-of-cutting-a-pizza lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1444] 切披萨的方案数
 *
 * https://leetcode.cn/problems/number-of-ways-of-cutting-a-pizza/description/
 *
 * algorithms
 * Hard (54.42%)
 * Total Accepted:    9.4K
 * Total Submissions: 15.7K
 * Testcase Example:  '["A..","AAA","..."]\n3'
 *
 * 给你一个 rows x cols 大小的矩形披萨和一个整数 k ，矩形包含两种字符： 'A' （表示苹果）和 '.' （表示空白格子）。你需要切披萨
 * k-1 次，得到 k 块披萨并送给别人。
 * 
 * 
 * 切披萨的每一刀，先要选择是向垂直还是水平方向切，再在矩形的边界上选一个切的位置，将披萨一分为二。如果垂直地切披萨，那么需要把左边的部分送给一个人，如果水平地切，那么需要把上面的部分送给一个人。在切完最后一刀后，需要把剩下来的一块送给最后一个人。
 * 
 * 请你返回确保每一块披萨包含 至少 一个苹果的切披萨方案数。由于答案可能是个很大的数字，请你返回它对 10^9 + 7 取余的结果。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 输入：pizza = ["A..","AAA","..."], k = 3
 * 输出：3 
 * 解释：上图展示了三种切披萨的方案。注意每一块披萨都至少包含一个苹果。
 * 
 * 
 * 示例 2：
 * 
 * 输入：pizza = ["A..","AA.","..."], k = 3
 * 输出：1
 * 
 * 
 * 示例 3：
 * 
 * 输入：pizza = ["A..","A..","..."], k = 1
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= rows, cols <= 50
 * rows == pizza.length
 * cols == pizza[i].length
 * 1 <= k <= 10
 * pizza 只包含字符 'A' 和 '.' 。
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/number-of-ways-of-cutting-a-pizza/solutions/2392051/ji-bai-100cong-di-gui-dao-di-tui-dao-you-dxz5/
//
// 记忆化搜索. 每次枚举最左侧 / 最上侧一刀即可.
public class Solution
{
    public int Ways(string[] pizza, int k)
    {
        const long Mod = (long)1e9 + 7;
        var (n, m) = (pizza.Length, pizza[0].Length);
        var S = new int[n + 1][];
        for (var i = 0; i <= n; i++) { S[i] = new int[m + 1]; }
        for (var i = 0; i < n; i++)
        {
            for (var j = 0; j < m; j++)
            {
                S[i + 1][j + 1] = pizza[i][j] is 'A' ? 1 : 0;
                S[i + 1][j + 1] += S[i + 1][j] + S[i][j + 1] - S[i][j];
            }
        }
        int getTotal(int x, int y) => S[n][m] - S[x][m] - S[n][y] + S[x][y];
        var memo = new Dictionary<(int, int, int), long>();
        long dfs(int x0, int y0, int k)
        {
            var key = (x0, y0, k);
            if (memo.ContainsKey(key)) { return memo[key]; }
            var total = getTotal(x0, y0);
            if (k is 1) { return memo[key] = total > 0 ? 1 : 0; }
            memo[key] = 0;
            for (var x = x0 + 1; x < n; x++)
            {
                if (total == getTotal(x, y0)) { continue; }
                memo[key] = (memo[key] + dfs(x, y0, k - 1)) % Mod;
            }
            for (var y = y0 + 1; y < m; y++)
            {
                if (total == getTotal(x0, y)) { continue; }
                memo[key] = (memo[key] + dfs(x0, y, k - 1)) % Mod;
            }
            return memo[key];
        }
        return (int)dfs(0, 0, k);
    }
}
