/*
 * @lc app=leetcode.cn id=count-the-number-of-houses-at-a-certain-distance-ii lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [3017] 按距离统计房屋对数目 II
 *
 * https://leetcode.cn/problems/count-the-number-of-houses-at-a-certain-distance-ii/description/
 *
 * algorithms
 * Hard (31.23%)
 * Total Accepted:    1.6K
 * Total Submissions: 5.1K
 * Testcase Example:  '3\n1\n3'
 *
 * 给你三个 正整数 n 、x 和 y 。
 * 
 * 在城市中，存在编号从 1 到 n 的房屋，由 n 条街道相连。对所有 1 <= i < n ，都存在一条街道连接编号为 i 的房屋与编号为 i + 1
 * 的房屋。另存在一条街道连接编号为 x 的房屋与编号为 y 的房屋。
 * 
 * 对于每个 k（1 <= k <= n），你需要找出所有满足要求的 房屋对 [house1, house2] ，即从 house1 到 house2
 * 需要经过的 最少 街道数为 k 。
 * 
 * 返回一个下标从 1 开始且长度为 n 的数组 result ，其中 result[k]
 * 表示所有满足要求的房屋对的数量，即从一个房屋到另一个房屋需要经过的 最少 街道数为 k 。
 * 
 * 注意，x 与 y 可以 相等 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3, x = 1, y = 3
 * 输出：[6,0,0]
 * 解释：让我们检视每个房屋对
 * - 对于房屋对 (1, 2)，可以直接从房屋 1 到房屋 2。
 * - 对于房屋对 (2, 1)，可以直接从房屋 2 到房屋 1。
 * - 对于房屋对 (1, 3)，可以直接从房屋 1 到房屋 3。
 * - 对于房屋对 (3, 1)，可以直接从房屋 3 到房屋 1。
 * - 对于房屋对 (2, 3)，可以直接从房屋 2 到房屋 3。
 * - 对于房屋对 (3, 2)，可以直接从房屋 3 到房屋 2。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 5, x = 2, y = 4
 * 输出：[10,8,2,0,0]
 * 解释：对于每个距离 k ，满足要求的房屋对如下：
 * - 对于 k == 1，满足要求的房屋对有 (1, 2), (2, 1), (2, 3), (3, 2), (2, 4), (4, 2), (3,
 * 4), (4, 3), (4, 5), 以及 (5, 4)。
 * - 对于 k == 2，满足要求的房屋对有 (1, 3), (3, 1), (1, 4), (4, 1), (2, 5), (5, 2), (3,
 * 5), 以及 (5, 3)。
 * - 对于 k == 3，满足要求的房屋对有 (1, 5)，以及 (5, 1) 。
 * - 对于 k == 4 和 k == 5，不存在满足要求的房屋对。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：n = 4, x = 1, y = 1
 * 输出：[6,4,2,0]
 * 解释：对于每个距离 k ，满足要求的房屋对如下：
 * - 对于 k == 1，满足要求的房屋对有 (1, 2), (2, 1), (2, 3), (3, 2), (3, 4), 以及 (4, 3)。
 * - 对于 k == 2，满足要求的房屋对有 (1, 3), (3, 1), (2, 4), 以及 (4, 2)。
 * - 对于 k == 3，满足要求的房屋对有 (1, 4), 以及 (4, 1)。
 * - 对于 k == 4，不存在满足要求的房屋对。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 2 <= n <= 10^5
 * 1 <= x, y <= n
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/count-the-number-of-houses-at-a-certain-distance-ii/solutions/2613373/yong-che-xiao-de-fang-shi-si-kao-pythonj-o253/
//
// 史诗级带讨论, 麻了
public class Solution
{
    public long[] CountOfPairs(int n, int x, int y)
    {
        if (x > y) { (x, y) = (y, x); }
        if (x + 1 >= y)
        {
            return Enumerable.Range(0, n)
                .Select(i => (long)(n - i - 1) * 2)
                .ToArray();
        }
        var D = new long[n + 1];
        void add(int p, int q)
        {
            D[p]++;
            D[q + 1]--;
        }
        for (var i = 1; i < n; i++)
        {
            if (i <= x)
            {
                // k - i <= y - k + x - i + 1
                var k = (x + y + 1) / 2;
                add(1, k - i);
                add(x - i + 1, y - k + x - i);
                add(x - i + 2, n - y + x - i + 1);
            }
            else if (i < (x + y) / 2)
            {
                // k - i <= y - k + i - x + 1
                var k = i + (y - x + 1) / 2;
                add(1, k - i);
                add(i - x + 1, y - k + i - x);
                add(i - x + 2, n - y + i - x + 1);
            }
            else
            {
                add(1, n - i);
            }
        }
        for (var i = 1; i <= n; i++) { D[i] += D[i - 1]; }
        for (var i = 1; i <= n; i++) { D[i] <<= 1; }
        return D[1..];
    }
}
