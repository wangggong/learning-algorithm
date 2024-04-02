/*
 * @lc app=leetcode.cn id=numbers-with-repeated-digits lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1012] 至少有 1 位重复的数字
 *
 * https://leetcode.cn/problems/numbers-with-repeated-digits/description/
 *
 * algorithms
 * Hard (43.55%)
 * Total Accepted:    12.7K
 * Total Submissions: 25.4K
 * Testcase Example:  '20'
 *
 * 给定正整数 n，返回在 [1, n] 范围内具有 至少 1 位 重复数字的正整数的个数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 20
 * 输出：1
 * 解释：具有至少 1 位重复数字的正数（<= 20）只有 11 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 100
 * 输出：10
 * 解释：具有至少 1 位重复数字的正数（<= 100）有 11，22，33，44，55，66，77，88，99 和 100 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：n = 1000
 * 输出：262
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^9
 * 
 * 
 */
public class Solution
{
    private Dictionary<(int, int, bool, bool), int> Memo = new();

    int Dfs(string s, int d, int mask, bool isLimit, bool isNum)
    {
        var key = (d, mask, isLimit, isNum);
        if (Memo.ContainsKey(key))
        {
            return Memo[key];
        }
        Memo[key] = 0;
        if (d == s.Length)
        {
            return (Memo[key] = isNum ? 1 : 0);
        }
        if (!isNum)
        {
            Memo[key] += Dfs(s, d + 1, mask, false, false);
        }
        var (l, r) = (isNum ? 0 : 1, isLimit ? (int)(s[d] - '0') : 9);
        for (var i = l; i <= r; i++)
        {
            if ((mask & (1 << i)) == 0)
            {
                Memo[key] += Dfs(s, d + 1, mask | (1 << i), isLimit && i == r, true);
            }
        }
        return Memo[key];
    }

    public int NumDupDigitsAtMostN(int n) => n - Dfs(n.ToString(), 0, 0, true, false);
}
