/*
 * @lc app=leetcode.cn id=count-stepping-numbers-in-range lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6957] 统计范围内的步进数字数目
 *
 * https://leetcode.cn/problems/count-stepping-numbers-in-range/description/
 *
 * algorithms
 * Hard (32.89%)
 * Total Accepted:    1K
 * Total Submissions: 3.1K
 * Testcase Example:  '"1"\n"11"'
 *
 * 给你两个正整数 low 和 high ，都用字符串表示，请你统计闭区间 [low, high] 内的 步进数字 数目。
 * 
 * 如果一个整数相邻数位之间差的绝对值都 恰好 是 1 ，那么这个数字被称为 步进数字 。
 * 
 * 请你返回一个整数，表示闭区间 [low, high] 之间步进数字的数目。
 * 
 * 由于答案可能很大，请你将它对 10^9 + 7 取余 后返回。
 * 
 * 注意：步进数字不能有前导 0 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：low = "1", high = "11"
 * 输出：10
 * 解释：区间 [1,11] 内的步进数字为 1 ，2 ，3 ，4 ，5 ，6 ，7 ，8 ，9 和 10 。总共有 10 个步进数字。所以输出为 10
 * 。
 * 
 * 示例 2：
 * 
 * 输入：low = "90", high = "101"
 * 输出：2
 * 解释：区间 [90,101] 内的步进数字为 98 和 101 。总共有 2 个步进数字。所以输出为 2 。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= int(low) <= int(high) < 10^100
 * 1 <= low.length, high.length <= 100
 * low 和 high 只包含数字。
 * low 和 high 都不含前导 0 。
 * 
 * 
 */

// 数位 DP. 参考 902.
public class Solution
{
    private const long Mod = (long)1e9 + 7;
    private const int All = (1 << 10) - 1;

    public int CountSteppingNumbers(string low, string high)
    {
        bool isStepping(string s) => Enumerable.Range(1, s.Length - 1)
            .All(i => Math.Abs((int)s[i] - (int)s[i - 1]) is 1);
        var memo = new Dictionary<(string, int, int, bool, bool), long>();
        long dfs(string s, int d, int mask, bool isLimit, bool isNum)
        {
            var key = (s, d, mask, isLimit, isNum);
            if (memo.ContainsKey(key)) { return memo[key]; }
            if (d - s.Length is 0) { return memo[key] = isNum ? 1 : 0; }
            memo[key] = 0;
            if (!isNum) { memo[key] += dfs(s, d + 1, mask, false, isNum); }
            var (l, r) = (isNum ? 0 : 1, isLimit ? (int)(s[d] - '0') : 9);
            for (var i = l; i <= r; i++)
            {
                if ((mask & 1 << i) is 0) { continue; }
                var next = 1 << i + 1 | (i is 0 ? 0 : 1 << i - 1);
                var limit = isLimit && i == r;
                memo[key] = (memo[key] + dfs(s, d + 1, next, limit, true)) % Mod;
            }
            return memo[key];
        }
        long count(string s) => dfs(s, 0, All, true, false);
        return (int)(
            (count(high) - count(low) + (isStepping(low) ? 1 : 0) + Mod) % Mod);
    }
}
