/*
 * Medium
 * 
 * 给你一个由 不同 正整数组成的数组 nums ，请你返回满足 a * b = c * d 的元组 (a, b, c, d) 的数量。其中 a、b、c 和 d 都是 nums 中的元素，且 a != b != c != d 。
 * 
 *  
 * 
 * 示例 1：
 * 
 * 输入：nums = [2,3,4,6]
 * 输出：8
 * 解释：存在 8 个满足题意的元组：
 * (2,6,3,4) , (2,6,4,3) , (6,2,3,4) , (6,2,4,3)
 * (3,4,2,6) , (4,3,2,6) , (3,4,6,2) , (4,3,6,2)
 * 示例 2：
 * 
 * 输入：nums = [1,2,4,5,10]
 * 输出：16
 * 解释：存在 16 个满足题意的元组：
 * (1,10,2,5) , (1,10,5,2) , (10,1,2,5) , (10,1,5,2)
 * (2,5,1,10) , (2,5,10,1) , (5,2,1,10) , (5,2,10,1)
 * (2,10,4,5) , (2,10,5,4) , (10,2,4,5) , (10,2,4,5)
 * (4,5,2,10) , (4,5,10,2) , (5,4,2,10) , (5,4,10,2)
 *  
 * 
 * 提示：
 * 
 * 1 <= nums.length <= 1000
 * 1 <= nums[i] <= 104
 * nums 中的所有元素 互不相同
 * 通过次数 8.9K
 * 提交次数 16.6K
 * 通过率 53.5%
 * 
 */
public class Solution
{
    public int TupleSameProduct(int[] nums)
    {
        var d = new Dictionary<int, int>();
        var ans = 0;
        for (var (i, n) = (0, nums.Length); i < n; i++)
        {
            for (var j = i + 1; j < n; j++)
            {
                var v = nums[i] * nums[j];
                d.TryGetValue(v, out var c);
                ans += c;
                d[v] = c + 1;
            }
        }
        return ans << 3;
    }
}
