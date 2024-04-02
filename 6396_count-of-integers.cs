/*
 * @lc app=leetcode.cn id=count-of-integers lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6396] 统计整数数目
 *
 * https://leetcode.cn/problems/count-of-integers/description/
 *
 * algorithms
 * Hard (34.67%)
 * Total Accepted:    795
 * Total Submissions: 2.3K
 * Testcase Example:  '"1"\n"12"\n1\n8'
 *
 * 给你两个数字字符串 num1 和 num2 ，以及两个整数 max_sum 和 min_sum 。如果一个整数 x
 * 满足以下条件，我们称它是一个好整数：
 * 
 * 
 * num1 <= x <= num2
 * min_sum <= digit_sum(x) <= max_sum.
 * 
 * 
 * 请你返回好整数的数目。答案可能很大，请返回答案对 10^9 + 7 取余后的结果。
 * 
 * 注意，digit_sum(x) 表示 x 各位数字之和。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：num1 = "1", num2 = "12", min_num = 1, max_num = 8
 * 输出：11
 * 解释：总共有 11 个整数的数位和在 1 到 8 之间，分别是 1,2,3,4,5,6,7,8,10,11 和 12 。所以我们返回 11 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：num1 = "1", num2 = "5", min_num = 1, max_num = 5
 * 输出：5
 * 解释：数位和在 1 到 5 之间的 5 个整数分别为 1,2,3,4 和 5 。所以我们返回 5 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= num1 <= num2 <= 10^22
 * 1 <= min_sum <= max_sum <= 400
 * 
 * 
 */
public class Solution
{
    private const long Mod = (long)1e9 + 7;
    public int Count(string num1, string num2, int minSum, int maxSum)
    {
        var memo = new Dictionary<(string, int, int, int, bool, bool), long>();
        long dfs(string s, int sum, int d, int total, bool isLimit, bool isNum)
        {
            var key = (s, sum, d, total, isLimit, isNum);
            if (memo.ContainsKey(key))
            {
                return memo[key];
            }
            memo[key] = 0;
            if (d == s.Length)
            {
                return (memo[key] = isNum ? 1 : 0);
            }
            if (!isNum)
            {
                memo[key] = (memo[key] + dfs(s, sum, d + 1, total, false, false)) % Mod;
            }
            var (l, r) = (isNum ? 0 : 1, isLimit ? (int)(s[d] - '0') : 9);
            for (var i = l; i <= r && total + i <= sum; i++)
            {
                memo[key] = (memo[key] + dfs(s, sum, d + 1, total + i, isLimit && i == r, true)) % Mod;
            }
            return memo[key];
        }
        var upper = (dfs(num2, maxSum, 0, 0, true, false) + Mod - dfs(num2, minSum - 1, 0, 0, true, false)) % Mod;
        var lower = (dfs(num1, maxSum, 0, 0, true, false) + Mod - dfs(num1, minSum - 1, 0, 0, true, false)) % Mod;
        var num1Sum = num1.Select(ch => (int)(ch - '0')).Sum();
        upper += minSum <= num1Sum && num1Sum <= maxSum ? 1 : 0;
        return (int)((upper + Mod - lower) % Mod);
    }
}
