/*
 * @lc app=leetcode.cn id=find-score-of-an-array-after-marking-all-elements lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6351] 标记所有元素后数组的分数
 *
 * https://leetcode.cn/problems/find-score-of-an-array-after-marking-all-elements/description/
 *
 * algorithms
 * Medium (45.35%)
 * Total Accepted:    1.8K
 * Total Submissions: 4K
 * Testcase Example:  '[2,1,3,4,5,2]'
 *
 * 给你一个数组 nums ，它包含若干正整数。
 * 
 * 一开始分数 score = 0 ，请你按照下面算法求出最后分数：
 * 
 * 
 * 从数组中选择最小且没有被标记的整数。如果有相等元素，选择下标最小的一个。
 * 将选中的整数加到 score 中。
 * 标记 被选中元素，如果有相邻元素，则同时标记 与它相邻的两个元素 。
 * 重复此过程直到数组中所有元素都被标记。
 * 
 * 
 * 请你返回执行上述算法后最后的分数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums = [2,1,3,4,5,2]
 * 输出：7
 * 解释：我们按照如下步骤标记元素：
 * - 1 是最小未标记元素，所以标记它和相邻两个元素：[2,1,3,4,5,2] 。
 * - 2 是最小未标记元素，所以标记它和左边相邻元素：[2,1,3,4,5,2] 。
 * - 4 是仅剩唯一未标记的元素，所以我们标记它：[2,1,3,4,5,2] 。
 * 总得分为 1 + 2 + 4 = 7 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：nums = [2,3,5,1,3,2]
 * 输出：5
 * 解释：我们按照如下步骤标记元素：
 * - 1 是最小未标记元素，所以标记它和相邻两个元素：[2,3,5,1,3,2] 。
 * - 2 是最小未标记元素，由于有两个 2 ，我们选择最左边的一个 2 ，也就是下标为 0 处的 2 ，以及它右边相邻的元素：[2,3,5,1,3,2]
 * 。
 * - 2 是仅剩唯一未标记的元素，所以我们标记它：[2,3,5,1,3,2] 。
 * 总得分为 1 + 2 + 2 = 5 。
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
    public long FindScore(int[] nums)
    {
        var n = nums.Length;
        var count = new Dictionary<int, int>();
        var index = new Dictionary<int, Queue<int>>();
        var values = new List<int>();
        for (var i = 0; i < n; i++)
        {
            var num = nums[i];
            if (!count.ContainsKey(num))
            {
                count[num] = 0;
                index[num] = new();
                values.Add(num);
            }
            count[num]++;
            index[num].Enqueue(i);
        }
        var m = values.Count();
        values.Sort();
        var visit = new bool[n];
        long score = 0;
        for (var i = 0; i < m; )
        {
            for (; i < m && count[values[i]] == 0; i++) { }
            if (i == m)
            {
                break;
            }
            var v = values[i];
            var currentIndex = index[v];
            for (; visit[currentIndex.Peek()]; currentIndex.Dequeue()) { }
            var k = currentIndex.Peek();
            score += (long)v;
            for (var j = Math.Max(k - 1, 0); j <= Math.Min(k + 1, n - 1); j++)
            {
                if (!visit[j])
                {
                    visit[j] = true;
                    count[nums[j]]--;
                }
            }
        }
        return score;
    }
}
