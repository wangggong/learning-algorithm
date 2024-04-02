/*
 * @lc app=leetcode.cn id=minimize-length-of-array-using-operations lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [100164] 通过操作使数组长度最小
 *
 * https://leetcode.cn/problems/minimize-length-of-array-using-operations/description/
 *
 * algorithms
 * Medium (26.60%)
 * Total Accepted:    1.3K
 * Total Submissions: 5K
 * Testcase Example:  '[1,4,3,1]'
 *
 * 给你一个下标从 0 开始的整数数组 nums ，它只包含 正 整数。
 * 
 * 你的任务是通过进行以下操作 任意次 （可以是 0 次） 最小化 nums 的长度：
 * 
 * 
 * 在 nums 中选择 两个不同 的下标 i 和 j ，满足 nums[i] > 0 且 nums[j] > 0 。
 * 将结果 nums[i] % nums[j] 插入 nums 的结尾。
 * 将 nums 中下标为 i 和 j 的元素删除。
 * 
 * 
 * 请你返回一个整数，它表示进行任意次操作以后 nums 的 最小长度 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,4,3,1]
 * 输出：1
 * 解释：使数组长度最小的一种方法是：
 * 操作 1 ：选择下标 2 和 1 ，插入 nums[2] % nums[1] 到数组末尾，得到 [1,4,3,1,3] ，然后删除下标为 2 和 1
 * 的元素。
 * nums 变为 [1,1,3] 。
 * 操作 2 ：选择下标 1 和 2 ，插入 nums[1] % nums[2] 到数组末尾，得到 [1,1,3,1] ，然后删除下标为 1 和 2
 * 的元素。
 * nums 变为 [1,1] 。
 * 操作 3 ：选择下标 1 和 0 ，插入 nums[1] % nums[0] 到数组末尾，得到 [1,1,0] ，然后删除下标为 1 和 0 的元素。
 * nums 变为 [0] 。
 * nums 的长度无法进一步减小，所以答案为 1 。
 * 1 是可以得到的最小长度。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,5,5,10,5]
 * 输出：2
 * 解释：使数组长度最小的一种方法是：
 * 操作 1 ：选择下标 0 和 3 ，插入 nums[0] % nums[3] 到数组末尾，得到 [5,5,5,10,5,5] ，然后删除下标为 0 和
 * 3 的元素。
 * nums 变为 [5,5,5,5] 。
 * 操作 2 ：选择下标 2 和 3 ，插入 nums[2] % nums[3] 到数组末尾，得到 [5,5,5,5,0] ，然后删除下标为 2 和 3
 * 的元素。
 * nums 变为 [5,5,0] 。
 * 操作 3 ：选择下标 0 和 1 ，插入 nums[0] % nums[1] 到数组末尾，得到 [5,5,0,0] ，然后删除下标为 0 和 1
 * 的元素。
 * nums 变为 [0,0] 。
 * nums 的长度无法进一步减小，所以答案为 2 。
 * 2 是可以得到的最小长度。
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [2,3,4]
 * 输出：1
 * 解释：使数组长度最小的一种方法是：
 * 操作 1 ：选择下标 1 和 2 ，插入 nums[1] % nums[2] 到数组末尾，得到 [2,3,4,3] ，然后删除下标为 1 和 2
 * 的元素。
 * nums 变为 [2,3] 。
 * 操作 2 ：选择下标 1 和 0 ，插入 nums[1] % nums[0] 到数组末尾，得到 [2,3,1] ，然后删除下标为 1 和 0 的元素。
 * nums 变为 [1] 。
 * nums 的长度无法进一步减小，所以答案为 1 。
 * 1 是可以得到的最小长度。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^9
 * 
 * 
 */
public class Solution
{
    public int MinimumArrayLength(int[] nums)
    {
        int gcd(int x, int y) => y is 0 ? x : gcd(y, x % y);
        var g = nums.Aggregate(gcd);
        return Math.Max((nums.Count(n => n == g) + 1) >> 1, 1);
    }
}

/*
 * // 参考: https://leetcode.cn/problems/minimize-length-of-array-using-operations/solutions/2613059/on-nao-jin-ji-zhuan-wan-pythonjavacgo-by-2lea/
 * public class Solution
 * {
 *     public int MinimumArrayLength(int[] nums)
 *     {
 *         var min = nums.Min();
 *         return nums.Any(n => (n % min) is not 0)
 *             ? 1
 *             : ((nums.Count(n => n == min) + 1) >> 1);
 *     }
 * }
 */
