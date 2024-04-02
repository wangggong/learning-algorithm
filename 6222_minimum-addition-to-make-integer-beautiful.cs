/*
 * @lc app=leetcode.cn id=minimum-addition-to-make-integer-beautiful lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6222] 美丽整数的最小增量
 *
 * https://leetcode.cn/problems/minimum-addition-to-make-integer-beautiful/description/
 *
 * algorithms
 * Medium (33.16%)
 * Total Accepted:    3.6K
 * Total Submissions: 10.4K
 * Testcase Example:  '16\n6'
 *
 * 给你两个正整数 n 和 target 。
 * 
 * 如果某个整数每一位上的数字相加小于或等于 target ，则认为这个整数是一个 美丽整数 。
 * 
 * 找出并返回满足 n + x 是 美丽整数 的最小非负整数 x 。生成的输入保证总可以使 n 变成一个美丽整数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 16, target = 6
 * 输出：4
 * 解释：最初，n 是 16 ，且其每一位数字的和是 1 + 6 = 7 。在加 4 之后，n 变为 20 且每一位数字的和变成 2 + 0 = 2
 * 。可以证明无法加上一个小于 4 的非负整数使 n 变成一个美丽整数。
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 467, target = 6
 * 输出：33
 * 解释：最初，n 是 467 ，且其每一位数字的和是 4 + 6 + 7 = 17 。在加 33 之后，n 变为 500 且每一位数字的和变成 5 + 0
 * + 0 = 5 。可以证明无法加上一个小于 33 的非负整数使 n 变成一个美丽整数。
 * 
 * 示例 3：
 * 
 * 输入：n = 1, target = 1
 * 输出：0
 * 解释：最初，n 是 1 ，且其每一位数字的和是 1 ，已经小于等于 target 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^12
 * 1 <= target <= 150
 * 生成的输入保证总可以使 n 变成一个美丽整数。
 * 
 * 
 */
public class Solution
{
    public long MakeIntegerBeautiful(long n, int target)
    {
        long ans = n, mask = 1;
        while (getDigitCnt(ans) > target)
        {
            for (; ans % (mask * 10) == 0; mask *= 10);
            ans += (10 - (ans / mask % 10)) * mask;
        }
        return ans - n;
    }

    private int getDigitCnt(long n)
    {
        int ans = 0;
        for (; n > 0; n /= 10)
        {
            ans += (int) (n % 10);
        }
        return ans;
    }
}
