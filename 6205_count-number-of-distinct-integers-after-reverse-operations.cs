/*
 * @lc app=leetcode.cn id=count-number-of-distinct-integers-after-reverse-operations lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6205] 反转之后不同整数的数目
 *
 * https://leetcode.cn/problems/count-number-of-distinct-integers-after-reverse-operations/description/
 *
 * algorithms
 * Medium (74.77%)
 * Total Accepted:    5.9K
 * Total Submissions: 7.9K
 * Testcase Example:  '[1,13,10,12,31]'
 *
 * 给你一个由 正 整数组成的数组 nums 。
 * 
 * 你必须取出数组中的每个整数，反转其中每个数位，并将反转后得到的数字添加到数组的末尾。这一操作只针对 nums 中原有的整数执行。
 * 
 * 返回结果数组中 不同 整数的数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,13,10,12,31]
 * 输出：6
 * 解释：反转每个数字后，结果数组是 [1,13,10,12,31,1,31,1,21,13] 。
 * 反转后得到的数字添加到数组的末尾并按斜体加粗表示。注意对于整数 10 ，反转之后会变成 01 ，即 1 。
 * 数组中不同整数的数目为 6（数字 1、10、12、13、21 和 31）。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,2,2]
 * 输出：1
 * 解释：反转每个数字后，结果数组是 [2,2,2,2,2,2] 。
 * 数组中不同整数的数目为 1（数字 2）。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^6
 * 
 * 
 */
public class Solution
{
    public int CountDistinctIntegers(int[] nums)
    {
        var S = new HashSet<int>();
        foreach (var n in nums)
        {
            S.Add(n);
            S.Add(Reverse(n));
        }
        return S.Count;
    }

    private int Reverse(int n)
    {
        int ans = 0;
        for (; n != 0; n /= 10)
        {
            ans = ans * 10 + (n % 10);
        }
        return ans;
    }
}
