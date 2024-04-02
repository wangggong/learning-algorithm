/*
 * @lc app=leetcode.cn id=find-the-grid-of-region-average lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100189] 找出网格的区域平均强度
 *
 * https://leetcode.cn/problems/find-the-grid-of-region-average/description/
 *
 * algorithms
 * Medium (42.49%)
 * Total Accepted:    1.3K
 * Total Submissions: 2.9K
 * Testcase Example:  '[[5,6,7,10],[8,9,10,10],[11,12,13,10]]\n3'
 *
 * 给你一个下标从 0 开始、大小为 m x n 的网格 image ，表示一个灰度图像，其中 image[i][j] 表示在范围 [0..255]
 * 内的某个像素强度。另给你一个 非负 整数 threshold 。
 * 
 * 如果 image[a][b] 和 image[c][d] 满足 |a - c| + |b - d| == 1 ，则称这两个像素是 相邻像素 。
 * 
 * 区域 是一个 3 x 3 的子网格，且满足区域中任意两个 相邻 像素之间，像素强度的 绝对差  小于或等于 threshold 。
 * 
 * 区域 内的所有像素都认为属于该区域，而一个像素 可以 属于 多个 区域。
 * 
 * 你需要计算一个下标从 0 开始、大小为 m x n 的网格 result ，其中 result[i][j] 是 image[i][j] 所属区域的 平均
 * 强度，向下取整 到最接近的整数。如果 image[i][j] 属于多个区域，result[i][j] 是这些区域的 “取整后的平均强度” 的 平均值，也
 * 向下取整 到最接近的整数。如果 image[i][j] 不属于任何区域，则 result[i][j] 等于 image[i][j] 。
 * 
 * 返回网格 result 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：image = [[5,6,7,10],[8,9,10,10],[11,12,13,10]], threshold = 3
 * 输出：[[9,9,9,9],[9,9,9,9],[9,9,9,9]]
 * 解释：图像中存在两个区域，如图片中的阴影区域所示。第一个区域的平均强度为 9 ，而第二个区域的平均强度为 9.67 ，向下取整为 9
 * 。两个区域的平均强度为 (9 + 9) / 2 = 9 。由于所有像素都属于区域 1 、区域 2 或两者，因此 result 中每个像素的强度都为 9
 * 。
 * 注意，在计算多个区域的平均值时使用了向下取整的值，因此使用区域 2 的平均强度 9 来进行计算，而不是 9.67 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：image = [[10,20,30],[15,25,35],[20,30,40],[25,35,45]], threshold = 12
 * 输出：[[25,25,25],[27,27,27],[27,27,27],[30,30,30]]
 * 解释：图像中存在两个区域，如图片中的阴影区域所示。第一个区域的平均强度为 25 ，而第二个区域的平均强度为 30 。两个区域的平均强度为 (25 +
 * 30) / 2 = 27.5 ，向下取整为 27 。图像中第 0 行的所有像素属于区域 1 ，因此 result 中第 0 行的所有像素为 25
 * 。同理，result 中第 3 行的所有像素为 30 。图像中第 1 行和第 2 行的像素属于区域 1 和区域 2 ，因此它们在 result 中的值为
 * 27 。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：image = [[5,6,7],[8,9,10],[11,12,13]], threshold = 1
 * 输出：[[5,6,7],[8,9,10],[11,12,13]]
 * 解释：图像中不存在任何区域，因此对于所有像素，result[i][j] == image[i][j] 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 3 <= n, m <= 500
 * 0 <= image[i][j] <= 255
 * 0 <= threshold <= 255
 * 
 * 
 */
public class Solution
{
    public int[][] ResultGrid(int[][] I, int threshold)
    {
        const int N = 9;
        var (n, m) = (I.Length, I[0].Length);
        var S = new int[n][];
        for (var i = 0; i < n; i++)
        {
            S[i] = new int[m];
            Array.Fill(S[i], -1);
        }
        (int, int) get(int x, int y, int v) => (x + v / 3 - 1, y + v % 3 - 1);
        int dist(int x, int y) => Math.Abs(x - y);
        void fill(int x, int y)
        {
            var total = 0;
            for (var s = 0; s < N; s++)
            {
                var (xs, ys) = get(x, y, s);
                total += I[xs][ys];
                for (var t = 0; t < N; t++)
                {
                    var (xt, yt) = get(x, y, t);
                    if (dist(xs, xt) + dist(ys, yt) == 1
                        && dist(I[xs][ys], I[xt][yt]) > threshold)
                    { return; }
                }
            }
            S[x][y] = total / N;
        }
        for (var x = 1; x + 1 < n; x++)
        {
            for (var y = 1; y + 1 < m; y++) { fill(x, y); }
        }
        for (var x = 0; x < n; x++)
        {
            for (var y = 0; y < m; y++)
            {
                var (count, total) = (0, 0);
                for (var d = 0; d < N; d++)
                {
                    var (xc, yc) = get(x, y, d);
                    if (0 <= xc && xc < n
                        && 0 <= yc && yc < m
                        && S[xc][yc] >= 0)
                    {
                        count++;
                        total += S[xc][yc];
                    }
                }
                if (count > 0) { I[x][y] = total / count; }
            }
        }
        return I;
    }
}
