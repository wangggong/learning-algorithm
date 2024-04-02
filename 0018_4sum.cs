/*
 * @lc app=leetcode.cn id=4sum lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [18] 四数之和
 *
 * https://leetcode.cn/problems/4sum/description/
 *
 * algorithms
 * Medium (36.77%)
 * Total Accepted:    474.4K
 * Total Submissions: 1.3M
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a],
 * nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
 * 
 * 
 * 0 <= a, b, c, d < n
 * a、b、c 和 d 互不相同
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 * 
 * 
 * 你可以按 任意顺序 返回答案 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,0,-1,0,-2,2], target = 0
 * 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,2,2,2,2], target = 8
 * 输出：[[2,2,2,2]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 * 
 * 
 */
public class Solution
{
    public IList<IList<int>> FourSum(int[] nums, int target) => KSum(
        nums.OrderBy(x => x).Select(x => (long)x).ToArray(), (long)target, 0, 4).ToList();

    private IEnumerable<IList<int>> KSum(long[] arr, long target, int start, int k)
    {
        if (k is 2)
        {
            foreach (var list in TwoSum(arr, target, start, k))
            {
                yield return list;
            }
            yield break;
        }
        for (var (i, n) = (start, arr.Length); i < n; i++)
        {
            if (i > start && arr[i] == arr[i - 1]) { continue; }
            foreach (var list in KSum(arr, target - arr[i], i + 1, k - 1))
            {
                var ans = new List<int> { (int)arr[i], };
                ans.AddRange(list);
                yield return ans;
            }
        }
    }

    private IEnumerable<IList<int>> TwoSum(long[] arr, long target, int start, int k)
    {
        for (var (p, q) = (start, arr.Length - 1); p < q; )
        {
            if (arr[p] + arr[q] == target)
            {
                yield return new List<int> { (int)arr[p], (int)arr[q], };
                for (var last = p; p < q && arr[p] == arr[last]; p++) { }
                continue;
            }
            if (arr[p] + arr[q] > target) { q--; }
            else { p++; }
        }
    }
}
