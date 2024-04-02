/*
 * @lc app=leetcode.cn id=prime-subtraction-operation lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6355] 质数减法运算
 *
 * https://leetcode.cn/problems/prime-subtraction-operation/description/
 *
 * algorithms
 * Medium (34.37%)
 * Total Accepted:    3.7K
 * Total Submissions: 10.7K
 * Testcase Example:  '[4,9,6,10]'
 *
 * 给你一个下标从 0 开始的整数数组 nums ，数组长度为 n 。
 * 
 * 你可以执行无限次下述运算：
 * 
 * 
 * 选择一个之前未选过的下标 i ，并选择一个 严格小于 nums[i] 的质数 p ，从 nums[i] 中减去 p 。
 * 
 * 
 * 如果你能通过上述运算使得 nums 成为严格递增数组，则返回 true ；否则返回 false 。
 * 
 * 严格递增数组 中的每个元素都严格大于其前面的元素。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [4,9,6,10]
 * 输出：true
 * 解释：
 * 在第一次运算中：选择 i = 0 和 p = 3 ，然后从 nums[0] 减去 3 ，nums 变为 [1,9,6,10] 。
 * 在第二次运算中：选择 i = 1 和 p = 7 ，然后从 nums[1] 减去 7 ，nums 变为 [1,2,6,10] 。
 * 第二次运算后，nums 按严格递增顺序排序，因此答案为 true 。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [6,8,11,12]
 * 输出：true
 * 解释：nums 从一开始就按严格递增顺序排序，因此不需要执行任何运算。
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [5,8,3]
 * 输出：false
 * 解释：可以证明，执行运算无法使 nums 按严格递增顺序排序，因此答案是 false 。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 1000
 * 1 <= nums[i] <= 1000
 * nums.length == n
 * 
 * 
 */
public class Solution
{
    public const int N = 1000;
    public int[] Primes;

    public Solution()
    {
        var isPrime = new bool[N + 1];
        for (var i = 0; i <= N; i++)
        {
            isPrime[i] = true;
        }
        for (var i = 2; i <= N; i++)
        {
            if (!isPrime[i])
            {
                continue;
            }
            for (var j = 2; i * j <= N; j++)
            {
                isPrime[i * j] = false;
            }
        }
        var primes = new List<int>();
        for (var i = 2; i <= N; i++)
        {
            if (isPrime[i])
            {
                primes.Add(i);
            }
        }
        Primes = primes.ToArray();
    }

    public bool PrimeSubOperation(int[] nums)
    {
        var n = nums.Length;
        var cur = n - 2;
        for (; cur >= 0; cur--)
        {
            if (nums[cur] < nums[cur + 1])
            {
                continue;
            }
            var i = 0;
            for (; i < Primes.Length && nums[cur] - Primes[i] >= nums[cur + 1]; i++) { }
            if (i == Primes.Length || Primes[i] >= nums[cur])
            {
                return false;
            }
            nums[cur] -= Primes[i];
        }
        return true;
    }
}
