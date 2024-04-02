/*
 * @lc app=leetcode.cn id=split-array-with-same-average lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [805] 数组的均值分割
 *
 * https://leetcode.cn/problems/split-array-with-same-average/description/
 *
 * algorithms
 * Hard (30.27%)
 * Total Accepted:    4.8K
 * Total Submissions: 15.3K
 * Testcase Example:  '[1,2,3,4,5,6,7,8]'
 *
 * 给定你一个整数数组 nums
 * 
 * 我们要将 nums 数组中的每个元素移动到 A 数组 或者 B 数组中，使得 A 数组和 B 数组不为空，并且 average(A) ==
 * average(B) 。
 * 
 * 如果可以完成则返回true ， 否则返回 false  。
 * 
 * 注意：对于数组 arr ,  average(arr) 是 arr 的所有元素除以 arr 长度的和。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: nums = [1,2,3,4,5,6,7,8]
 * 输出: true
 * 解释: 我们可以将数组分割为 [1,4,5,8] 和 [2,3,6,7], 他们的平均值都是4.5。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: nums = [3,1]
 * 输出: false
 * 
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= nums.length <= 30
 * 0 <= nums[i] <= 10^4
 * 
 * 
 */
public class Solution
{
    public bool SplitArraySameAverage(int[] nums)
    {
        var n = nums.Length;
        if (n == 1) { return false; }
        var m = n / 2;
        var tot = (from num in nums select num).Sum();
        var sumList = new List<HashSet<int>>();
        for (int i = 0; i <= m; i++) { sumList.Add(new HashSet<int>()); }
        void PreDfs(int k, int cnt, int cur)
        {
            if (k == m)
            {
                sumList[cnt].Add(cur);
                return;
            }
            PreDfs(k + 1, cnt, cur);
            PreDfs(k + 1, cnt + 1, cur + nums[k]);
            return;
        }
        PreDfs(0, 0, 0);
        bool SufDfs(int k, int cnt, int cur)
        {
            if (k == n)
            {
                for (int i = 0; i <= m; i++)
                {
                    var totCnt = i + cnt;
                    if (totCnt == 0 || totCnt == n || tot * totCnt % n != 0) { continue; }
                    var target = tot * totCnt / n - cur;
                    if (sumList[i].Contains(target)) { return true; }
                }
                return false;
            }
            return SufDfs(k + 1, cnt, cur) || SufDfs(k + 1, cnt + 1, cur + nums[k]);
        }
        return SufDfs(m, 0, 0);
    }
}
