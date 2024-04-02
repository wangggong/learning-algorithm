/*
 * @lc app=leetcode.cn id=max-value-of-equation lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1499] 满足不等式的最大值
 *
 * https://leetcode.cn/problems/max-value-of-equation/description/
 *
 * algorithms
 * Hard (41.05%)
 * Total Accepted:    5.2K
 * Total Submissions: 12.3K
 * Testcase Example:  '[[1,3],[2,0],[5,10],[6,-10]]\n1'
 *
 * 给你一个数组 points 和一个整数 k 。数组中每个元素都表示二维平面上的点的坐标，并按照横坐标 x 的值从小到大排序。也就是说 points[i]
 * = [xi, yi] ，并且在 1 <= i < j <= points.length 的前提下， xi < xj 总成立。
 * 
 * 请你找出 yi + yj + |xi - xj| 的 最大值，其中 |xi - xj| <= k 且 1 <= i < j <=
 * points.length。
 * 
 * 题目测试数据保证至少存在一对能够满足 |xi - xj| <= k 的点。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：points = [[1,3],[2,0],[5,10],[6,-10]], k = 1
 * 输出：4
 * 解释：前两个点满足 |xi - xj| <= 1 ，代入方程计算，则得到值 3 + 0 + |1 - 2| = 4 。第三个和第四个点也满足条件，得到值
 * 10 + -10 + |5 - 6| = 1 。
 * 没有其他满足条件的点，所以返回 4 和 1 中最大的那个。
 * 
 * 示例 2：
 * 
 * 输入：points = [[0,0],[3,0],[9,2]], k = 3
 * 输出：3
 * 解释：只有前两个点满足 |xi - xj| <= 3 ，代入方程后得到值 0 + 0 + |0 - 3| = 3 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= points.length <= 10^5
 * points[i].length == 2
 * -10^8 <= points[i][0], points[i][1] <= 10^8
 * 0 <= k <= 2 * 10^8
 * 对于所有的1 <= i < j <= points.length ，points[i][0] < points[j][0] 都成立。也就是说，xi
 * 是严格递增的。
 * 
 * 
 */
public class Solution
{
    public int FindMaxValueOfEquation(int[][] points, int k)
    {
        var n = points.Length;
        var ans = int.MinValue;
        var Q = new PriorityQueue<int[], int>();
        for (var (p, q) = (0, 0); p < n; p++)
        {
            var point = points[p];
            for (; q < n && points[q][0] - point[0] <= k; q++)
            {
                Q.Enqueue(points[q], -(points[q][0] + points[q][1]));
            }
            for (; Q.Count > 0; Q.Dequeue())
            {
                var d = Q.Peek()[0] - point[0];
                if (0 < d && d <= k) { break; }
            }
            if (Q.Count > 0)
            {
                Q.TryPeek(out var _, out var val);
                ans = Math.Max(ans, point[1] - point[0] - val);
            }
        }
        // 这里其实是没用的: 一次就够了.
        // for (var (p, q) = (n - 1, n - 1); p >= 0; p--)
        // {
        //     var point = points[p];
        //     for (; q >= 0 && point[0] - points[q][0] <= k; q--)
        //     {
        //         Q.Enqueue(points[q], -(points[q][1] - points[q][0]));
        //     }
        //     for (; Q.Count > 0; Q.Dequeue())
        //     {
        //         var d = point[0] - Q.Peek()[0];
        //         if (0 < d && d <= k) { break; }
        //     }
        //     if (Q.Count > 0)
        //     {
        //         Q.TryPeek(out var _, out var val);
        //         ans = Math.Max(ans, point[1] + point[0] - val);
        //     }
        // }
        return ans;
    }

    // 官解给了一个单调队列.
    /*
     * public int FindMaxValueOfEquation(int[][] points, int k)
     * {
     *     var ans = int.MinValue;
     *     var Q = new int[points.Length * 2][];
     *     var (head, tail) = (0, 0);
     *     foreach (var point in points)
     *     {
     *         while (head < tail && point[0] - Q[head][0] > k) { head++; }
     *         if (head < tail)
     *         {
     *             ans = Math.Max(ans, point[1] + Q[head][1] + point[0] - Q[head][0]);
     *         }
     *         while (head < tail && Q[tail - 1][1] - Q[tail - 1][0] < point[1] - point[0])
     *         {
     *             tail--;
     *         }
     *         Q[tail] = point;
     *         tail++;
     *     }
     *     return ans;
     * }
     */
}
