/*
 * @lc app=leetcode.cn id=smallest-integer-divisible-by-k lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1015] 可被 K 整除的最小整数
 *
 * https://leetcode.cn/problems/smallest-integer-divisible-by-k/description/
 *
 * algorithms
 * Medium (37.32%)
 * Total Accepted:    15.5K
 * Total Submissions: 34.8K
 * Testcase Example:  '1'
 *
 * 给定正整数 k ，你需要找出可以被 k 整除的、仅包含数字 1 的最 小 正整数 n 的长度。
 * 
 * 返回 n 的长度。如果不存在这样的 n ，就返回-1。
 * 
 * 注意： n 不符合 64 位带符号整数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：k = 1
 * 输出：1
 * 解释：最小的答案是 n = 1，其长度为 1。
 * 
 * 示例 2：
 * 
 * 
 * 输入：k = 2
 * 输出：-1
 * 解释：不存在可被 2 整除的正整数 n 。
 * 
 * 示例 3：
 * 
 * 
 * 输入：k = 3
 * 输出：3
 * 解释：最小的答案是 n = 111，其长度为 3。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= k <= 10^5
 * 
 * 
 */


/*
 * public class Solution
 * {
 *     public int SmallestRepunitDivByK(int k)
 *     {
 *         var S = new HashSet<int>();
 *         for (var (n, curr) = (1, 1 % k); true; n++, curr = (curr * 10 + 1) % k)
 *         {
 *             if (curr == 0) { return n; }
 *             if (S.Contains(curr)) { return -1; }
 *             S.Add(curr);
 *         }
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/smallest-integer-divisible-by-k/solution/san-chong-suan-fa-you-hua-pythonjavacgo-tk4cj/
// 
// 可以证明除了 2 和 5 的倍数都有解. 反证见上面链接. 关键词: 抽屉原理, 11...11 = (10 ^ n - 1) / 9
public class Solution
{
    public int SmallestRepunitDivByK(int k)
    {
        if (k % 2 == 0 || k % 5 == 0)
        {
            return -1;
        }
        for (var (n, curr) = (1, 1 % k); true; n++, curr = (curr * 10 + 1) % k)
        {
            if (curr == 0)
            {
                return n;
            }
        }
    }
}
