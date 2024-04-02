/*
 * @lc app=leetcode.cn id=3sum-closest lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode.cn/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (44.83%)
 * Total Accepted:    480.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
 * 
 * 返回这三个数的和。
 * 
 * 假定每组输入只存在恰好一个解。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [0,0,0], target = 1
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 3 <= nums.length <= 1000
 * -1000 <= nums[i] <= 1000
 * -10^4 <= target <= 10^4
 * 
 * 
 */
public class Solution
{
    private const int Maxn = 0x3f3f3f3f;

    public int ThreeSumClosest(int[] nums, int target)
    {
        var ans = Maxn;
        void relax(int k) => ans = Math.Abs(k - target) < Math.Abs(ans - target)
            ? k
            : ans;
        Array.Sort(nums);
        for (var (i, n) = (0, nums.Length); i + 2 < n; i++)
        {
            for (var j = i + 1; j + 1 < n; j++)
            {
                var (p, q) = (j + 1, n - 1);
                while (p < q)
                {
                    var mid = (p + q) >> 1;
                    if (nums[i] + nums[j] + nums[mid] >= target)
                    {
                        q = mid;
                    }
                    else
                    {
                        p = mid + 1;
                    }
                }
                relax(nums[i] + nums[j] + nums[p]);
                if (p - 1 > j)
                {
                    relax(nums[i] + nums[j] + nums[p - 1]);
                }
            }
        }
        return ans;
    }
}
