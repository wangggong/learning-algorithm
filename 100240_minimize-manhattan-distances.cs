/*
 * @lc app=leetcode.cn id=minimize-manhattan-distances lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100240] 最小化曼哈顿距离
 *
 * https://leetcode.cn/problems/minimize-manhattan-distances/description/
 *
 * algorithms
 * Hard (24.97%)
 * Total Accepted:    1.1K
 * Total Submissions: 4.2K
 * Testcase Example:  '[[3,10],[5,15],[10,2],[4,4]]'
 *
 * 给你一个下标从 0 开始的数组 points ，它表示二维平面上一些点的整数坐标，其中 points[i] = [xi, yi] 。
 * 
 * 两点之间的距离定义为它们的曼哈顿距离。
 * 
 * 请你恰好移除一个点，返回移除后任意两点之间的 最大 距离可能的 最小 值。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：points = [[3,10],[5,15],[10,2],[4,4]]
 * 输出：12
 * 解释：移除每个点后的最大距离如下所示：
 * - 移除第 0 个点后，最大距离在点 (5, 15) 和 (10, 2) 之间，为 |5 - 10| + |15 - 2| = 18 。
 * - 移除第 1 个点后，最大距离在点 (3, 10) 和 (10, 2) 之间，为 |3 - 10| + |10 - 2| = 15 。
 * - 移除第 2 个点后，最大距离在点 (5, 15) 和 (4, 4) 之间，为 |5 - 4| + |15 - 4| = 12 。
 * - 移除第 3 个点后，最大距离在点 (5, 15) 和 (10, 2) 之间的，为 |5 - 10| + |15 - 2| = 18 。
 * 在恰好移除一个点后，任意两点之间的最大距离可能的最小值是 12 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：points = [[1,1],[1,1],[1,1]]
 * 输出：0
 * 解释：移除任一点后，任意两点之间的最大距离都是 0 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 3 <= points.length <= 10^5
 * points[i].length == 2
 * 1 <= points[i][0], points[i][1] <= 10^8
 * 
 * 
 */

// 参考:
// - https://leetcode.cn/problems/minimize-manhattan-distances/solutions/2716752/xiang-xi-tui-dao-jing-dian-jie-lun-fu-li-3rkq
// - https://leetcode.cn/problems/minimize-manhattan-distances/solutions/2716755/tu-jie-man-ha-dun-ju-chi-heng-deng-shi-b-op84
// 1. 什么叫 "切比雪夫距离"
// 2. 如何通过切比雪夫距离把两个 (或多个) 绝对值符号拆开
//
// PS 看到 "闵可夫斯基距离" 的时候有点想起来了, 之前学过的都忘了!
public class Solution
{
    private record Distance(int D, int I);

    private IEnumerable<int> MaximumIndexes(int[][] points, int k)
    {
        var As = new List<Distance>();
        var Bs = new List<Distance>();
        foreach (var (p, i) in points.Select((p, i) => (p, i)))
        {
            if (i == k)
            {
                continue;
            }
            As.Add(new(p[0] - p[1], i));
            Bs.Add(new(p[0] + p[1], i));
        }
        As = As.OrderBy(x => x.D)
            .ToList();
        yield return As.First().I;
        yield return As.Last().I;
        Bs = Bs.OrderBy(x => x.D)
            .ToList();
        yield return Bs.First().I;
        yield return Bs.Last().I;
    }

    public int MinimumDistance(int[][] points) => MaximumIndexes(points, -1)
        .Select(i =>
        {
            var indexes = MaximumIndexes(points, i)
                .ToList();
            return indexes.SelectMany(i => indexes
                .Select(j => Math.Abs(points[i][0] - points[j][0])
                    + Math.Abs(points[i][1] - points[j][1]))).Max();
        }).Min();
}
