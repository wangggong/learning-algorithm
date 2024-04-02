/*
 * @lc app=leetcode.cn id=number-of-visible-people-in-a-queue lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1944] 队列中可以看到的人数
 *
 * https://leetcode.cn/problems/number-of-visible-people-in-a-queue/description/
 *
 * algorithms
 * Hard (63.92%)
 * Total Accepted:    13.5K
 * Total Submissions: 19.9K
 * Testcase Example:  '[10,6,8,5,11,9]'
 *
 * 有 n 个人排成一个队列，从左到右 编号为 0 到 n - 1 。给你以一个整数数组 heights ，每个整数 互不相同，heights[i] 表示第
 * i 个人的高度。
 * 
 * 一个人能 看到 他右边另一个人的条件是这两人之间的所有人都比他们两人 矮 。更正式的，第 i 个人能看到第 j 个人的条件是 i < j 且
 * min(heights[i], heights[j]) > max(heights[i+1], heights[i+2], ...,
 * heights[j-1]) 。
 * 
 * 请你返回一个长度为 n 的数组 answer ，其中 answer[i] 是第 i 个人在他右侧队列中能 看到 的 人数 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 
 * 
 * 输入：heights = [10,6,8,5,11,9]
 * 输出：[3,1,2,1,1,0]
 * 解释：
 * 第 0 个人能看到编号为 1 ，2 和 4 的人。
 * 第 1 个人能看到编号为 2 的人。
 * 第 2 个人能看到编号为 3 和 4 的人。
 * 第 3 个人能看到编号为 4 的人。
 * 第 4 个人能看到编号为 5 的人。
 * 第 5 个人谁也看不到因为他右边没人。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：heights = [5,1,2,3,10]
 * 输出：[4,1,1,1,0]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == heights.length
 * 1 <= n <= 10^5
 * 1 <= heights[i] <= 10^5
 * heights 中所有数 互不相同 。
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     public int[] CanSeePersonsCount(int[] heights)
 *     {
 *         const int Maxn = 0x3f3f3f3f;
 *         var n = heights.Length;
 *         var rights = new int[n];
 *         int getHeight(int k) => k >= n
 *             ? Maxn
 *             : heights[k];
 *         var S = new Stack<int>();
 *         S.Push(n);
 *         for (var i = n - 1; i >= 0; i--)
 *         {
 *             for (; getHeight(S.Peek()) < getHeight(i); S.Pop()) { }
 *             rights[i] = S.Peek();
 *             S.Push(i);
 *         }
 *         var memo = new int[n];
 *         Array.Fill(memo, -1);
 *         int dfs(int k) => k >= n
 *             ? 0
 *             : memo[k] = (memo[k] >= 0
 *                 ? memo[k]
 *                 : dfs(rights[k]) + 1);
 *         return Enumerable.Range(0, n)
 *             .Select(i => dfs(i + 1) - dfs(rights[i]) + (rights[i] != n ? 1 : 0))
 *             .ToArray();
 *     }
 * }
 */

// 参考: https://leetcode.cn/problems/number-of-visible-people-in-a-queue/solutions/2591558/dan-diao-zhan-de-ben-zhi-ji-shi-qu-diao-8tp3s/
//
// 相比于维护树的深度差而言简单直接.
public class Solution
{
    public int[] CanSeePersonsCount(int[] heights)
    {
        var n = heights.Length;
        var ans = new int[n];
        var S = new Stack<int>();
        for (var i = n - 1; i >= 0; i--)
        {
            for (; S.Count > 0 && heights[S.Peek()] < heights[i]; S.Pop())
            { ans[i]++; }
            if (S.Count > 0) { ans[i]++; }
            S.Push(i);
        }
        return ans;
    }
}
