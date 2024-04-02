/*
 * @lc app=leetcode.cn id=maximize-the-minimum-powered-city lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2528] 最大化城市的最小供电站数目
 *
 * https://leetcode.cn/problems/maximize-the-minimum-powered-city/description/
 *
 * algorithms
 * Hard (40.33%)
 * Total Accepted:    2.4K
 * Total Submissions: 5.8K
 * Testcase Example:  '[1,2,4,5,0]\n1\n2'
 *
 * 给你一个下标从 0 开始长度为 n 的整数数组 stations ，其中 stations[i] 表示第 i 座城市的供电站数目。
 * 
 * 每个供电站可以在一定 范围 内给所有城市提供电力。换句话说，如果给定的范围是 r ，在城市 i 处的供电站可以给所有满足 |i - j| <= r 且
 * 0 <= i, j <= n - 1 的城市 j 供电。
 * 
 * 
 * |x| 表示 x 的 绝对值 。比方说，|7 - 5| = 2 ，|3 - 10| = 7 。
 * 
 * 
 * 一座城市的 电量 是所有能给它供电的供电站数目。
 * 
 * 政府批准了可以额外建造 k 座供电站，你需要决定这些供电站分别应该建在哪里，这些供电站与已经存在的供电站有相同的供电范围。
 * 
 * 给你两个整数 r 和 k ，如果以最优策略建造额外的发电站，返回所有城市中，最小供电站数目的最大值是多少。
 * 
 * 这 k 座供电站可以建在多个城市。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：stations = [1,2,4,5,0], r = 1, k = 2
 * 输出：5
 * 解释：
 * 最优方案之一是把 2 座供电站都建在城市 1 。
 * 每座城市的供电站数目分别为 [1,4,4,5,0] 。
 * - 城市 0 的供电站数目为 1 + 4 = 5 。
 * - 城市 1 的供电站数目为 1 + 4 + 4 = 9 。
 * - 城市 2 的供电站数目为 4 + 4 + 5 = 13 。
 * - 城市 3 的供电站数目为 5 + 4 = 9 。
 * - 城市 4 的供电站数目为 5 + 0 = 5 。
 * 供电站数目最少是 5 。
 * 无法得到更优解，所以我们返回 5 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：stations = [4,4,4,4], r = 0, k = 3
 * 输出：4
 * 解释：
 * 无论如何安排，总有一座城市的供电站数目是 4 ，所以最优解是 4 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == stations.length
 * 1 <= n <= 10^5
 * 0 <= stations[i] <= 10^5
 * 0 <= r <= n - 1
 * 0 <= k <= 10^9
 * 
 * 
 */

public class Solution
{
    public long MaxPower(int[] stations, int r, int k)
    {
        var n = stations.Length;
        var ps = new long[n + 1];
        for (var i = 0; i < n; i++)
        {
            ps[i + 1] = ps[i] + stations[i];
        }
        var power = new long[n];
        for (var i = 0; i < n; i++)
        {
            power[i] = ps[Math.Min(i + r + 1, n)] - ps[Math.Max(i - r, 0)];
        }
        bool check(long target)
        {
            var diff = new long[n];
            long sumD = 0;
            long totalNeed = 0;
            for (var i = 0; i < n; i++)
            {
                sumD += diff[i];
                var need = target - power[i] - sumD;
                if (need > 0)
                {
                    totalNeed += need;
                    if (totalNeed > k)
                    {
                        return false;
                    }
                    sumD += need;
                    if (i + r * 2 + 1 < n)
                    {
                        diff[i + r * 2 + 1] -= need;
                    }
                }
            }
            return true;
        }
        long p = 0, q = long.MaxValue;
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (!check(mid))
            {
                q = mid;
            }
            else
            {
                p = mid + 1;
            }
        }
        return p - 1;
    }
}
