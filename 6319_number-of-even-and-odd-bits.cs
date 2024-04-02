/*
 * @lc app=leetcode.cn id=number-of-even-and-odd-bits lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6319] 奇偶位数
 *
 * https://leetcode.cn/problems/number-of-even-and-odd-bits/description/
 *
 * algorithms
 * Easy (73.67%)
 * Total Accepted:    4.9K
 * Total Submissions: 6.6K
 * Testcase Example:  '17'
 *
 * 给你一个 正 整数 n 。
 * 
 * 用 even 表示在 n 的二进制形式（下标从 0 开始）中值为 1 的偶数下标的个数。
 * 
 * 用 odd 表示在 n 的二进制形式（下标从 0 开始）中值为 1 的奇数下标的个数。
 * 
 * 返回整数数组 answer ，其中 answer = [even, odd] 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 17
 * 输出：[2,0]
 * 解释：17 的二进制形式是 10001 。 
 * 下标 0 和 下标 4 对应的值为 1 。 
 * 共有 2 个偶数下标，0 个奇数下标。
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 2
 * 输出：[0,1]
 * 解释：2 的二进制形式是 10 。 
 * 下标 1 对应的值为 1 。 
 * 共有 0 个偶数下标，1 个奇数下标。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 1000
 * 
 * 
 */
public class Solution
{
    private const int D = 32;

    public int[] EvenOddBit(int n)
    {
        var ans = new int[2];
        for (var i = 0; i < D; i++, n >>= 1)
        {
            if ((n & 1) != 0)
            {
                ans[i % 2]++;
            }
        }
        return ans;
    }
}
