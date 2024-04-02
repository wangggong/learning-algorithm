/*
 * @lc app=leetcode.cn id=count-the-digits-that-divide-a-number lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6278] 统计能整除数字的位数
 *
 * https://leetcode.cn/problems/count-the-digits-that-divide-a-number/description/
 *
 * algorithms
 * Easy (89.96%)
 * Total Accepted:    3.2K
 * Total Submissions: 3.6K
 * Testcase Example:  '7'
 *
 * 给你一个整数 num ，返回 num 中能整除 num 的数位的数目。
 * 
 * 如果满足 nums % val == 0 ，则认为整数 val 可以整除 nums 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：num = 7
 * 输出：1
 * 解释：7 被自己整除，因此答案是 1 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：num = 121
 * 输出：2
 * 解释：121 可以被 1 整除，但无法被 2 整除。由于 1 出现两次，所以返回 2 。
 * 
 * 
 * 示例 3：
 * 
 * 输入：num = 1248
 * 输出：4
 * 解释：1248 可以被它每一位上的数字整除，因此答案是 4 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= num <= 10^9
 * num 的数位中不含 0
 * 
 * 
 */
public class Solution
{
    public int CountDigits(int num)
    {
        var ans = 0;
        for (int k = num; k > 0; k /= 10)
        {
            if (num % (k % 10) == 0)
            {
                ans++;
            }
        }
        return ans;
    }
}
