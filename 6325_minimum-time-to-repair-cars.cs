/*
 * @lc app=leetcode.cn id=minimum-time-to-repair-cars lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6325] 修车的最少时间
 *
 * https://leetcode.cn/problems/minimum-time-to-repair-cars/description/
 *
 * algorithms
 * Medium (39.60%)
 * Total Accepted:    1.5K
 * Total Submissions: 3.7K
 * Testcase Example:  '[4,2,3,1]\n10'
 *
 * 给你一个整数数组 ranks ，表示一些机械工的 能力值 。ranksi 是第 i 位机械工的能力值。能力值为 r 的机械工可以在 r * n^2
 * 分钟内修好 n 辆车。
 * 
 * 同时给你一个整数 cars ，表示总共需要修理的汽车数目。
 * 
 * 请你返回修理所有汽车 最少 需要多少时间。
 * 
 * 注意：所有机械工可以同时修理汽车。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：ranks = [4,2,3,1], cars = 10
 * 输出：16
 * 解释：
 * - 第一位机械工修 2 辆车，需要 4 * 2 * 2 = 16 分钟。
 * - 第二位机械工修 2 辆车，需要 2 * 2 * 2 = 8 分钟。
 * - 第三位机械工修 2 辆车，需要 3 * 2 * 2 = 12 分钟。
 * - 第四位机械工修 4 辆车，需要 1 * 4 * 4 = 16 分钟。
 * 16 分钟是修理完所有车需要的最少时间。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：ranks = [5,1,8], cars = 6
 * 输出：16
 * 解释：
 * - 第一位机械工修 1 辆车，需要 5 * 1 * 1 = 5 分钟。
 * - 第二位机械工修 4 辆车，需要 1 * 4 * 4 = 16 分钟。
 * - 第三位机械工修 1 辆车，需要 8 * 1 * 1 = 8 分钟。
 * 16 分钟时修理完所有车需要的最少时间。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= ranks.length <= 10^5
 * 1 <= ranks[i] <= 100
 * 1 <= cars <= 10^6
 * 
 * 
 */
public class Solution
{
    private static long Maxn = (long)1e14;

    public long RepairCars(int[] ranks, int cars)
    {
        bool canFixCars(long k) => ranks.Select(r => (long)(Math.Sqrt(k / (long)r))).Sum() >= (long)cars;
        long p = 0;
        long q = Maxn + 10;
        while (p < q)
        {
            var mid = (p + q) >> 1;
            if (canFixCars(mid))
            {
                q = mid;
            }
            else
            {
                p = mid + 1;
            }
        }
        return p;
    }
}
