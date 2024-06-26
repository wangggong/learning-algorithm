/*
 * @lc app=leetcode.cn id=minimum-operations-to-reduce-an-integer-to-0 lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6365] 将整数减少到零需要的最少操作数
 *
 * https://leetcode.cn/problems/minimum-operations-to-reduce-an-integer-to-0/description/
 *
 * algorithms
 * Easy (42.14%)
 * Total Accepted:    2.9K
 * Total Submissions: 6.8K
 * Testcase Example:  '39'
 *
 * 给你一个正整数 n ，你可以执行下述操作 任意 次：
 * 
 * 
 * n 加上或减去 2 的某个 幂
 * 
 * 
 * 返回使 n 等于 0 需要执行的 最少 操作数。
 * 
 * 如果 x == 2^i 且其中 i >= 0 ，则数字 x 是 2 的幂。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 39
 * 输出：3
 * 解释：我们可以执行下述操作：
 * - n 加上 2^0 = 1 ，得到 n = 40 。
 * - n 减去 2^3 = 8 ，得到 n = 32 。
 * - n 减去 2^5 = 32 ，得到 n = 0 。
 * 可以证明使 n 等于 0 需要执行的最少操作数是 3 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 54
 * 输出：3
 * 解释：我们可以执行下述操作：
 * - n 加上 2^1 = 2 ，得到 n = 56 。
 * - n 加上 2^3 = 8 ，得到 n = 64 。
 * - n 减去 2^6 = 64 ，得到 n = 0 。
 * 使 n 等于 0 需要执行的最少操作数是 3 。 
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^5
 * 
 * 
 */
public class Solution
{
    public int MinOperations(int n)
    {
        for (var i = 0; true; i++)
        {
            if (n == 0)
            {
                return i;
            }
            for (; (n & 1) == 0; n >>= 1) { }
            n = ((n >> 1) & 1) == 1 ? n + 1 : n - 1;
        }
    }
}
