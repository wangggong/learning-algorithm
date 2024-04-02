/*
 * @lc app=leetcode.cn id=shortest-subarray-to-be-removed-to-make-array-sorted lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1574] 删除最短的子数组使剩余数组有序
 *
 * https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/description/
 *
 * algorithms
 * Medium (37.30%)
 * Total Accepted:    13.6K
 * Total Submissions: 34.2K
 * Testcase Example:  '[1,2,3,10,4,2,3,5]'
 *
 * 给你一个整数数组 arr ，请你删除一个子数组（可以为空），使得 arr 中剩下的元素是 非递减 的。
 * 
 * 一个子数组指的是原数组中连续的一个子序列。
 * 
 * 请你返回满足题目要求的最短子数组的长度。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：arr = [1,2,3,10,4,2,3,5]
 * 输出：3
 * 解释：我们需要删除的最短子数组是 [10,4,2] ，长度为 3 。剩余元素形成非递减数组 [1,2,3,3,5] 。
 * 另一个正确的解为删除子数组 [3,10,4] 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：arr = [5,4,3,2,1]
 * 输出：4
 * 解释：由于数组是严格递减的，我们只能保留一个元素。所以我们需要删除长度为 4 的子数组，要么删除 [5,4,3,2]，要么删除 [4,3,2,1]。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：arr = [1,2,3]
 * 输出：0
 * 解释：数组已经是非递减的了，我们不需要删除任何元素。
 * 
 * 
 * 示例 4：
 * 
 * 
 * 输入：arr = [1]
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= arr.length <= 10^5
 * 0 <= arr[i] <= 10^9
 * 
 * 
 */

/*
 * // 二分, 写起来有点蛋疼. O(nlogn)
 * public class Solution
 * {
 *     public int FindLengthOfShortestSubarray(int[] arr)
 *     {
 *         var n = arr.Length;
 *         bool canMakeSorted(int k)
 *         {
 *             if (k >= n - 1)
 *             {
 *                 return true;
 *             }
 *             var left = new int[n - k];
 *             for (var i = 0; i + k < n; i++)
 *             {
 *                 left[i] = arr[i];
 *             }
 *             var m = n - k - 1;
 *             for (; m > 0 && left[m] >= left[m - 1]; m--) { }
 *             if (m == 0)
 *             {
 *                 return true;
 *             }
 *             for (var i = n - k - 1; i >= 0; i--)
 *             {
 *                 left[i] = arr[i + k];
 *                 if (i + 1 < n - k && left[i] > left[i + 1])
 *                 {
 *                     return false;
 *                 }
 *                 for (; m > 0 && left[m] >= left[m - 1]; m--) { }
 *                 if (m == 0)
 *                 {
 *                     return true;
 *                 }
 *             }
 *             return false;
 *         }
 *         var p = 0;
 *         var q = n + 1;
 *         while (p < q)
 *         {
 *             var mid = (p + q) >> 1;
 *             if (canMakeSorted(mid))
 *             {
 *                 q = mid;
 *             }
 *             else
 *             {
 *                 p = mid + 1;
 *             }
 *         }
 *         return p;
 *     }
 * }
 */

// 正解是同向双指针...
//
// 参考: https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/solution/dong-hua-yi-xie-jiu-cuo-liang-chong-xie-iijwz/
public class Solution
{
    public int FindLengthOfShortestSubarray(int[] arr)
    {
        var n = arr.Length;
        var right = n - 1;
        for (; right > 0 && arr[right] >= arr[right - 1]; right--) { }
        if (right == 0)
        {
            return 0;
        }
        var ans = right;
        for (var left = 0; left == 0 || arr[left] >= arr[left - 1]; left++)
        {
            for (; right < n && arr[right] < arr[left]; right++) { }
            ans = Math.Min(ans, right - left - 1);
        }
        return ans;
    }
}
