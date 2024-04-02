/*
 * @lc app=leetcode.cn id=number-of-beautiful-integers-in-the-range lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [8013] 范围中美丽整数的数目
 *
 * https://leetcode.cn/problems/number-of-beautiful-integers-in-the-range/description/
 *
 * algorithms
 * Hard (23.86%)
 * Total Accepted:    1K
 * Total Submissions: 4.4K
 * Testcase Example:  '10\n20\n3'
 *
 * 给你正整数 low ，high 和 k 。
 * 
 * 如果一个数满足以下两个条件，那么它是 美丽的 ：
 * 
 * 
 * 偶数数位的数目与奇数数位的数目相同。
 * 这个整数可以被 k 整除。
 * 
 * 
 * 请你返回范围 [low, high] 中美丽整数的数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：low = 10, high = 20, k = 3
 * 输出：2
 * 解释：给定范围中有 2 个美丽数字：[12,18]
 * - 12 是美丽整数，因为它有 1 个奇数数位和 1 个偶数数位，而且可以被 k = 3 整除。
 * - 18 是美丽整数，因为它有 1 个奇数数位和 1 个偶数数位，而且可以被 k = 3 整除。
 * 以下是一些不是美丽整数的例子：
 * - 16 不是美丽整数，因为它不能被 k = 3 整除。
 * - 15 不是美丽整数，因为它的奇数数位和偶数数位的数目不相等。
 * 给定范围内总共有 2 个美丽整数。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：low = 1, high = 10, k = 1
 * 输出：1
 * 解释：给定范围中有 1 个美丽数字：[10]
 * - 10 是美丽整数，因为它有 1 个奇数数位和 1 个偶数数位，而且可以被 k = 1 整除。
 * 给定范围内总共有 1 个美丽整数。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：low = 5, high = 5, k = 2
 * 输出：0
 * 解释：给定范围中有 0 个美丽数字。
 * - 5 不是美丽整数，因为它的奇数数位和偶数数位的数目不相等。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 < low <= high <= 10^9
 * 0 < k <= 20
 * 
 * 
 */
public class Solution
{
    public int NumberOfBeautifulIntegers(int low, int high, int k)
    {
        var memo = new Dictionary<(string, int, bool, bool, int, int), int>();
        int dfs(string s, int d, bool isLimit, bool isNum, int diff, int rem)
        {
            var key = (s, d, isLimit, isNum, diff, rem);
            if (memo.ContainsKey(key)) { return memo[key]; }
            if (d == s.Length)
            {
                return memo[key] = (isNum, diff, rem) is (true, 0, 0) ? 1 : 0;
            }
            memo[key] = 0;
            if (!isNum) { memo[key] += dfs(s, d + 1, false, isNum, diff, rem); }
            var digit = (int)(s[d] - '0');
            var (l, r) = (isNum ? 0 : 1, isLimit ? digit : 9);
            for (var i = l; i <= r; i++)
            {
                var nd = diff + (i % 2 == 0 ? 1 : -1);
                var nr = (rem * 10 + i) % k;
                memo[key] += dfs(s, d + 1, isLimit && i == r, true, nd, nr);
            }
            return memo[key];
        }
        int count(int n) => dfs(n.ToString(), 0, true, false, 0, 0);
        return count(high) - (low is 0 ? 0 : count(low - 1));
    }
}
