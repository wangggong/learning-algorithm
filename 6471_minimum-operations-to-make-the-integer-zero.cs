/*
 * @lc app=leetcode.cn id=minimum-operations-to-make-the-integer-zero lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6471] 得到整数零需要执行的最少操作数
 *
 * https://leetcode.cn/problems/minimum-operations-to-make-the-integer-zero/description/
 *
 * algorithms
 * Medium (22.59%)
 * Total Accepted:    1.2K
 * Total Submissions: 5.3K
 * Testcase Example:  '3\n-2'
 *
 * 给你两个整数：num1 和 num2 。
 * 
 * 在一步操作中，你需要从范围 [0, 60] 中选出一个整数 i ，并从 num1 减去 2^i + num2 。
 * 
 * 请你计算，要想使 num1 等于 0 需要执行的最少操作数，并以整数形式返回。
 * 
 * 如果无法使 num1 等于 0 ，返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：num1 = 3, num2 = -2
 * 输出：3
 * 解释：可以执行下述步骤使 3 等于 0 ：
 * - 选择 i = 2 ，并从 3 减去 2^2 + (-2) ，num1 = 3 - (4 + (-2)) = 1 。
 * - 选择 i = 2 ，并从 1 减去 2^2 + (-2) ，num1 = 1 - (4 + (-2)) = -1 。
 * - 选择 i = 0 ，并从 -1 减去 2^0 + (-2) ，num1 = (-1) - (1 + (-2)) = 0 。
 * 可以证明 3 是需要执行的最少操作数。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：num1 = 5, num2 = 7
 * 输出：-1
 * 解释：可以证明，执行操作无法使 5 等于 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= num1 <= 10^9
 * -10^9 <= num2 <= 10^9
 * 
 * 
 */
public class Solution
{
    private int CountOneDigits(long n)
    {
        var ans = 0;
        for (; n > 0; n &= n - 1)
        {
            ans++;
        }
        return ans;
    }

    public int MakeTheIntegerZero(int num1, int num2)
    {
        for (var (k, longNum1, longNum2) = (1, (long)num1, (long)num2);
            longNum1 > longNum2;
            k++)
        {
            longNum1 -= longNum2;
            if ((long)k <= longNum1 && CountOneDigits(longNum1) <= k)
            {
                return k;
            }
        }
        return -1;
    }
}
