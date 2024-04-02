/*
 * @lc app=leetcode.cn id=count-the-number-of-powerful-integers lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100163] 统计强大整数的数目
 *
 * https://leetcode.cn/problems/count-the-number-of-powerful-integers/description/
 *
 * algorithms
 * Hard (28.78%)
 * Total Accepted:    804
 * Total Submissions: 2.8K
 * Testcase Example:  '1\n6000\n4\n"124"'
 *
 * 给你三个整数 start ，finish 和 limit 。同时给你一个下标从 0 开始的字符串 s ，表示一个 正 整数。
 * 
 * 如果一个 正 整数 x 末尾部分是 s （换句话说，s 是 x 的 后缀），且 x 中的每个数位至多是 limit ，那么我们称 x 是 强大的 。
 * 
 * 请你返回区间 [start..finish] 内强大整数的 总数目 。
 * 
 * 如果一个字符串 x 是 y 中某个下标开始（包括 0 ），到下标为 y.length - 1 结束的子字符串，那么我们称 x 是 y
 * 的一个后缀。比方说，25 是 5125 的一个后缀，但不是 512 的后缀。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：start = 1, finish = 6000, limit = 4, s = "124"
 * 输出：5
 * 解释：区间 [1..6000] 内的强大数字为 124 ，1124 ，2124 ，3124 和 4124 。这些整数的各个数位都 <= 4 且
 * "124" 是它们的后缀。注意 5124 不是强大整数，因为第一个数位 5 大于 4 。
 * 这个区间内总共只有这 5 个强大整数。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：start = 15, finish = 215, limit = 6, s = "10"
 * 输出：2
 * 解释：区间 [15..215] 内的强大整数为 110 和 210 。这些整数的各个数位都 <= 6 且 "10" 是它们的后缀。
 * 这个区间总共只有这 2 个强大整数。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：start = 1000, finish = 2000, limit = 4, s = "3000"
 * 输出：0
 * 解释：区间 [1000..2000] 内的整数都小于 3000 ，所以 "3000" 不可能是这个区间内任何整数的后缀。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= start <= finish <= 10^15
 * 1 <= limit <= 9
 * 1 <= s.length <= floor(log10(finish)) + 1
 * s 数位中每个数字都小于等于 limit 。
 * s 不包含任何前导 0 。
 * 
 * 
 */

// 理解还是不到位啊.
public class Solution
{
    public long NumberOfPowerfulInt(long start, long finish, int limit, string suffix)
    {
        var m = suffix.Length;
        var (s, t) = ((start - 1).ToString(), finish.ToString());
        if (s.Length < t.Length)
        { s = new string('0', t.Length - s.Length) + s; }
        var memo = new Dictionary<(string, int, bool, bool), long>();
        long dfs(string s, int d, bool isLimit, bool isNum)
        {
            var key = (s, d, isLimit, isNum);
            if (memo.ContainsKey(key)) { return memo[key]; }
            var n = s.Length;
            if (n - d is 0) { return memo[key] = isNum ? 1 : 0; }
            memo[key] = 0;
            // if (!isNum) { memo[key] += dfs(s, d + 1, false, isNum); }
            var (l, r) = (0 /* isNum ? 0 : 1 */, isLimit ? (int)(s[d] - '0') : limit);
            for (var i = l; i <= r; i++)
            {
                if (n - d <= m && i != (int)(suffix[m - (n - d)] - '0')) { continue; }
                if (i > limit) { continue; }
                memo[key] += dfs(s, d + 1, isLimit && i == r, true);
            }
            return memo[key];
        }
        return dfs(t, 0, true, false) - dfs(s, 0, true, false);
    }
}
