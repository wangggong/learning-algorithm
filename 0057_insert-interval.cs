/*
 * @lc app=leetcode.cn id=insert-interval lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [57] 插入区间
 *
 * https://leetcode.cn/problems/insert-interval/description/
 *
 * algorithms
 * Medium (42.31%)
 * Total Accepted:    159.9K
 * Total Submissions: 377.6K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * 给你一个 无重叠的 ，按照区间起始端点排序的区间列表。
 * 
 * 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
 * 输出：[[1,5],[6,9]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * 输出：[[1,2],[3,10],[12,16]]
 * 解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
 * 
 * 示例 3：
 * 
 * 
 * 输入：intervals = [], newInterval = [5,7]
 * 输出：[[5,7]]
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：intervals = [[1,5]], newInterval = [2,3]
 * 输出：[[1,5]]
 * 
 * 
 * 示例 5：
 * 
 * 
 * 输入：intervals = [[1,5]], newInterval = [2,7]
 * 输出：[[1,7]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * intervals[i].length == 2
 * 0 
 * intervals 根据 intervals[i][0] 按 升序 排列
 * newInterval.length == 2
 * 0 
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/insert-interval/solutions/472151/cha-ru-qu-jian-by-leetcode-solution/
//
// 官解的额外空间是 `O(1)` 的.
public class Solution
{
    public int[][] Insert(int[][] intervals, int[] newInterval)
    {
        var ans = new List<int[]>();
        var (left, right) = (newInterval[0], newInterval[1]);
        var inserted = false;
        foreach (var interval in intervals)
        {
            if (interval[1] < left)
            {
                ans.Add(interval);
                continue;
            }
            if (interval[0] > right)
            {
                if (!inserted)
                {
                    ans.Add(new int[] { left, right, });
                    inserted = true;
                }
                ans.Add(interval);
                continue;
            }
            left = Math.Min(left, interval[0]);
            right = Math.Max(right, interval[1]);
        }
        if (!inserted) { ans.Add(new int[] { left, right, }); }
        return ans.ToArray();
    }
}

/*
 * // 转换一下问题, 其实是把问题变容易了.
 * public class Solution
 * {
 *     public int[][] Insert(int[][] intervals, int[] newInterval)
 *     {
 *         var intervalList = new List<int[]>();
 *         var (k, n) = (0, intervals.Length);
 *         for (; k < n && intervals[k][0] < newInterval[0]; k++) { intervalList.Add(intervals[k]); }
 *         intervalList.Add(newInterval);
 *         for (; k < n; k++) { intervalList.Add(intervals[k]); }
 *         var ans = new List<int[]>();
 *         foreach (var interval in intervalList)
 *         {
 *             if (ans.Any() && ans.Last()[1] >= interval[0]) { ans.Last()[1] = Math.Max(ans.Last()[1], interval[1]); }
 *             else { ans.Add(interval); }
 *         }
 *         return ans.ToArray();
 *     }
 * }
 */
