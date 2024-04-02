/*
 * @lc app=leetcode.cn id=points-that-intersect-with-cars lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [8029] 与车相交的点
 *
 * https://leetcode.cn/problems/points-that-intersect-with-cars/description/
 *
 * algorithms
 * Easy (73.50%)
 * Total Accepted:    4.7K
 * Total Submissions: 6.4K
 * Testcase Example:  '[[3,6],[1,5],[4,7]]'
 *
 * 给你一个下标从 0 开始的二维整数数组 nums 表示汽车停放在数轴上的坐标。对于任意下标 i，nums[i] = [starti, endi] ，其中
 * starti 是第 i 辆车的起点，endi 是第 i 辆车的终点。
 * 
 * 返回数轴上被车 任意部分 覆盖的整数点的数目。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [[3,6],[1,5],[4,7]]
 * 输出：7
 * 解释：从 1 到 7 的所有点都至少与一辆车相交，因此答案为 7 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [[1,3],[5,8]]
 * 输出：7
 * 解释：1、2、3、5、6、7、8 共计 7 个点满足至少与一辆车相交，因此答案为 7 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 100
 * nums[i].length == 2
 * 1 <= starti <= endi <= 100
 * 
 * 
 */
public class Solution
{
    public int NumberOfPoints(IList<IList<int>> nums)
    {
        const int N = 100;
        var arr = new int[N + 2];
        foreach (var num in nums)
        {
            arr[num[0]]++;
            arr[num[1] + 1]--;
        }
        for (var i = 0; i < N; i++)
        {
            arr[i + 1] += arr[i];
        }
        return arr.Where(x => x > 0)
            .Count();
    }
}
