/*
 * @lc app=leetcode.cn id=maximum-number-of-integers-to-choose-from-a-range-i lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6304] 从一个范围内选择最多整数 I
 *
 * https://leetcode.cn/problems/maximum-number-of-integers-to-choose-from-a-range-i/description/
 *
 * algorithms
 * Medium (56.60%)
 * Total Accepted:    2.1K
 * Total Submissions: 3.7K
 * Testcase Example:  '[1,6,5]\n5\n6'
 *
 * 给你一个整数数组 banned 和两个整数 n 和 maxSum 。你需要按照以下规则选择一些整数：
 * 
 * 
 * 被选择整数的范围是 [1, n] 。
 * 每个整数 至多 选择 一次 。
 * 被选择整数不能在数组 banned 中。
 * 被选择整数的和不超过 maxSum 。
 * 
 * 
 * 请你返回按照上述规则 最多 可以选择的整数数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：banned = [1,6,5], n = 5, maxSum = 6
 * 输出：2
 * 解释：你可以选择整数 2 和 4 。
 * 2 和 4 在范围 [1, 5] 内，且它们都不在 banned 中，它们的和是 6 ，没有超过 maxSum 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：banned = [1,2,3,4,5,6,7], n = 8, maxSum = 1
 * 输出：0
 * 解释：按照上述规则无法选择任何整数。
 * 
 * 
 * 示例 3：
 * 
 * 输入：banned = [11], n = 7, maxSum = 50
 * 输出：7
 * 解释：你可以选择整数 1, 2, 3, 4, 5, 6 和 7 。
 * 它们都在范围 [1, 7] 中，且都没出现在 banned 中，它们的和是 28 ，没有超过 maxSum 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= banned.length <= 10^4
 * 1 <= banned[i], n <= 10^4
 * 1 <= maxSum <= 10^9
 * 
 * 
 */
public class Solution
{
    public int MaxCount(int[] banned, int n, int maxSum)
    {
        var S = banned.ToHashSet();
        var ans = 0;
        var tot = 0;
        for (var i = 1; i <= n; i++)
        {
            if (!S.Contains(i) && tot + i <= maxSum)
            {
                ans++;
                tot += i;
            }
        }
        return ans;
    }
}
