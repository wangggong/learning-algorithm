/*
 * @lc app=leetcode.cn id=apply-operations-to-maximize-score lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [7023] 操作使得分最大
 *
 * https://leetcode.cn/problems/apply-operations-to-maximize-score/description/
 *
 * algorithms
 * Hard (35.51%)
 * Total Accepted:    1.4K
 * Total Submissions: 4K
 * Testcase Example:  '[8,3,9,3,8]\n2'
 *
 * 给你一个长度为 n 的正整数数组 nums 和一个整数 k 。
 * 
 * 一开始，你的分数为 1 。你可以进行以下操作至多 k 次，目标是使你的分数最大：
 * 
 * 
 * 选择一个之前没有选过的 非空 子数组 nums[l, ..., r] 。
 * 从 nums[l, ..., r] 里面选择一个 质数分数 最高的元素 x 。如果多个元素质数分数相同且最高，选择下标最小的一个。
 * 将你的分数乘以 x 。
 * 
 * 
 * nums[l, ..., r] 表示 nums 中起始下标为 l ，结束下标为 r 的子数组，两个端点都包含。
 * 
 * 一个整数的 质数分数 等于 x 不同质因子的数目。比方说， 300 的质数分数为 3 ，因为 300 = 2 * 2 * 3 * 5 * 5 。
 * 
 * 请你返回进行至多 k 次操作后，可以得到的 最大分数 。
 * 
 * 由于答案可能很大，请你将结果对 10^9 + 7 取余后返回。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [8,3,9,3,8], k = 2
 * 输出：81
 * 解释：进行以下操作可以得到分数 81 ：
 * - 选择子数组 nums[2, ..., 2] 。nums[2] 是子数组中唯一的元素。所以我们将分数乘以 nums[2] ，分数变为 1 * 9 =
 * 9 。
 * - 选择子数组 nums[2, ..., 3] 。nums[2] 和 nums[3] 质数分数都为 1 ，但是 nums[2]
 * 下标更小。所以我们将分数乘以 nums[2] ，分数变为 9 * 9 = 81 。
 * 81 是可以得到的最高得分。
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [19,12,14,6,10,18], k = 3
 * 输出：4788
 * 解释：进行以下操作可以得到分数 4788 ：
 * - 选择子数组 nums[0, ..., 0] 。nums[0] 是子数组中唯一的元素。所以我们将分数乘以 nums[0] ，分数变为 1 * 19 =
 * 19 。
 * - 选择子数组 nums[5, ..., 5] 。nums[5] 是子数组中唯一的元素。所以我们将分数乘以 nums[5] ，分数变为 19 * 18
 * = 342 。
 * - 选择子数组 nums[2, ..., 3] 。nums[2] 和 nums[3] 质数分数都为 2，但是 nums[2]
 * 下标更小。所以我们将分数乘以 nums[2] ，分数变为  342 * 14 = 4788 。
 * 4788 是可以得到的最高的分。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length == n <= 10^5
 * 1 <= nums[i] <= 10^5
 * 1 <= k <= min(n * (n + 1) / 2, 10^9)
 * 
 * 
 */
public class Solution
{
    private const long Mod = (long)1e9 + 7;
    private const int N = (int)1e5;
    private static int[] Scores;

    static Solution()
    {
        Scores = new int[N + 1];
        Array.Fill(Scores, -1);
        Scores[1] = 0;
        int score(int n)
        {
            if (Scores[n] >= 0) { return Scores[n]; }
            for (var i = 2; true; i++)
            {
                if (n % i == 0)
                {
                    for (; n % i == 0; n /= i) { }
                    return 1 + score(n);
                }
            }
        }
        for (var i = 2; i <= N; i++) { Scores[i] = score(i); }
    }

    private long Pow(long n, int k)
    {
        var ans = 1L;
        for (; k is not 0; k >>= 1)
        {
            if ((k & 1) is not 0) { ans = (ans * n) % Mod; }
            n = (n * n) % Mod;
        }
        return ans;
    }

    public int MaximumScore(IList<int> nums, int k)
    {
        var n = nums.Count();
        var scores = nums.Select(n => Scores[n]).ToArray();
        var lefts = new int[n];
        var S = new Stack<int>();
        for (var i = 0; i < n; i++)
        {
            for (; S.Any() && scores[S.Peek()] < scores[i]; S.Pop()) { }
            lefts[i] = S.Any() ? S.Peek() : -1;
            S.Push(i);
        }
        var rights = new int[n];
        S = new Stack<int>();
        for (var i = n - 1; i >= 0; i--)
        {
            for (; S.Any() && scores[S.Peek()] <= scores[i]; S.Pop()) { }
            rights[i] = S.Any() ? S.Peek() : n;
            S.Push(i);
        }
        var ans = 1L;
        foreach (var (x, c) in nums
            .Select((x, i) => (x: x, c: (i - lefts[i]) * (rights[i] - i)))
            .OrderByDescending(x => x.x))
        {
            if (k <= 0) { break; }
            var p = Math.Min(k, c);
            k -= p;
            ans = (ans * Pow((long)x, p)) % Mod;
        }
        return (int)ans;
    }
}
