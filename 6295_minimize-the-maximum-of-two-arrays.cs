/*
 * @lc app=leetcode.cn id=minimize-the-maximum-of-two-arrays lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6295] 最小化两个数组中的最大值
 *
 * https://leetcode.cn/problems/minimize-the-maximum-of-two-arrays/description/
 *
 * algorithms
 * Medium (26.52%)
 * Total Accepted:    708
 * Total Submissions: 2.6K
 * Testcase Example:  '2\n7\n1\n3'
 *
 * 给你两个数组 arr1 和 arr2 ，它们一开始都是空的。你需要往它们中添加正整数，使它们满足以下条件：
 * 
 * 
 * arr1 包含 uniqueCnt1 个 互不相同 的正整数，每个整数都 不能 被 divisor1 整除 。
 * arr2 包含 uniqueCnt2 个 互不相同 的正整数，每个整数都 不能 被 divisor2 整除 。
 * arr1 和 arr2 中的元素 互不相同 。
 * 
 * 
 * 给你 divisor1 ，divisor2 ，uniqueCnt1 和 uniqueCnt2 ，请你返回两个数组中 最大元素 的 最小值 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：divisor1 = 2, divisor2 = 7, uniqueCnt1 = 1, uniqueCnt2 = 3
 * 输出：4
 * 解释：
 * 我们可以把前 4 个自然数划分到 arr1 和 arr2 中。
 * arr1 = [1] 和 arr2 = [2,3,4] 。
 * 可以看出两个数组都满足条件。
 * 最大值是 4 ，所以返回 4 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：divisor1 = 3, divisor2 = 5, uniqueCnt1 = 2, uniqueCnt2 = 1
 * 输出：3
 * 解释：
 * arr1 = [1,2] 和 arr2 = [3] 满足所有条件。
 * 最大值是 3 ，所以返回 3 。
 * 
 * 示例 3：
 * 
 * 
 * 输入：divisor1 = 2, divisor2 = 4, uniqueCnt1 = 8, uniqueCnt2 = 2
 * 输出：15
 * 解释：
 * 最终数组为 arr1 = [1,3,5,7,9,11,13,15] 和 arr2 = [2,6] 。
 * 上述方案是满足所有条件的最优解。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= divisor1, divisor2 <= 10^5
 * 1 <= uniqueCnt1, uniqueCnt2 < 10^9
 * 2 <= uniqueCnt1 + uniqueCnt2 <= 10^9
 * 
 * 
 */
public class Solution
{
    public int MinimizeSet(int divisor1, int divisor2, int uniqueCnt1, int uniqueCnt2)
    {
        var d1 = (long) divisor1;
        var d2 = (long) divisor2;
        long gcd(long x, long y) => y == 0 ? x : gcd(y, x % y);
        var lcm = d1 * d2 / gcd(d1, d2);
        bool check(long k)
        {
            var a1 = k / d2 - k / lcm;
            var a2 = k / d1 - k / lcm;
            var rest = k - a1 - a2 - k / lcm;
            return rest >= Math.Max(uniqueCnt1 - a1, 0) + Math.Max(uniqueCnt2 - a2, 0);
        }
        var p = (long) 1;
        var q = (long) 2e9 + 10;
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (check(mid))
            {
                q = mid;
            }
            else
            {
                p = mid + 1;
            }
        }
        return (int) p;
    }
}
