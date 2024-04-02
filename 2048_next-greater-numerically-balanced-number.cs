/*
 * @lc app=leetcode.cn id=next-greater-numerically-balanced-number lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2048] 下一个更大的数值平衡数
 *
 * https://leetcode.cn/problems/next-greater-numerically-balanced-number/description/
 *
 * algorithms
 * Medium (46.28%)
 * Total Accepted:    12.5K
 * Total Submissions: 22.9K
 * Testcase Example:  '1'
 *
 * 如果整数  x 满足：对于每个数位 d ，这个数位 恰好 在 x 中出现 d 次。那么整数 x 就是一个 数值平衡数 。
 * 
 * 给你一个整数 n ，请你返回 严格大于 n 的 最小数值平衡数 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 1
 * 输出：22
 * 解释：
 * 22 是一个数值平衡数，因为：
 * - 数字 2 出现 2 次 
 * 这也是严格大于 1 的最小数值平衡数。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1000
 * 输出：1333
 * 解释：
 * 1333 是一个数值平衡数，因为：
 * - 数字 1 出现 1 次。
 * - 数字 3 出现 3 次。 
 * 这也是严格大于 1000 的最小数值平衡数。
 * 注意，1022 不能作为本输入的答案，因为数字 0 的出现次数超过了 0 。
 * 
 * 示例 3：
 * 
 * 
 * 输入：n = 3000
 * 输出：3133
 * 解释：
 * 3133 是一个数值平衡数，因为：
 * - 数字 1 出现 1 次。
 * - 数字 3 出现 3 次。 
 * 这也是严格大于 3000 的最小数值平衡数。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 <= n <= 10^6
 * 
 * 
 */
public class Solution
{
    private const int N = (int)1e6;
    private static List<int> numbers = new();

    static Solution()
    {
        bool isBalanced(int k) => k
            .ToString()
            .GroupBy(x => x)
            .All(g => g.Count() == (int)(g.Key - '0'));
        for (var i = 1; numbers.LastOrDefault(0) < N; i++)
        {
            if (isBalanced(i)) { numbers.Add(i); }
        }
    }

    public int NextBeautifulNumber(int n)
    {
        var (p, q) = (0, numbers.Count());
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (numbers[mid] > n) { q = mid; }
            else {p = mid + 1; }
        }
        return numbers[p];
    }
}
