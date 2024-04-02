/*
 * @lc app=leetcode.cn id=number-of-different-subsequences-gcds lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1819] 序列中不同最大公约数的数目
 *
 * https://leetcode.cn/problems/number-of-different-subsequences-gcds/description/
 *
 * algorithms
 * Hard (41.98%)
 * Total Accepted:    11.8K
 * Total Submissions: 19.7K
 * Testcase Example:  '[6,10,3]'
 *
 * 给你一个由正整数组成的数组 nums 。
 * 
 * 数字序列的 最大公约数 定义为序列中所有整数的共有约数中的最大整数。
 * 
 * 
 * 例如，序列 [4,6,16] 的最大公约数是 2 。
 * 
 * 
 * 数组的一个 子序列 本质是一个序列，可以通过删除数组中的某些元素（或者不删除）得到。
 * 
 * 
 * 例如，[2,5,10] 是 [1,2,1,2,4,1,5,10] 的一个子序列。
 * 
 * 
 * 计算并返回 nums 的所有 非空 子序列中 不同 最大公约数的 数目 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [6,10,3]
 * 输出：5
 * 解释：上图显示了所有的非空子序列与各自的最大公约数。
 * 不同的最大公约数为 6 、10 、3 、2 和 1 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [5,15,40,5,6]
 * 输出：7
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 1 
 * 
 * 
 */

// 参考: https://leetcode.cn/problems/number-of-different-subsequences-gcds/solution/ji-bai-100mei-ju-gcdxun-huan-you-hua-pyt-get7/
//
// 枚举每个值, 判断是不是最大公约数.
// 对于每个值 `i`, 找到 `nums` 中 `i` 的所有倍数, 全部 `gcd` 掉. 如果这样都得不到 `i` 的话, `i` 就不可能是 `nums` 中任意一个序列的 `gcd`.
// 注意, 找到 `nums` 中 `i` 的所有倍数可以通过哈希表从 `O(n)` 优化为 `O(U / i)`.
// 这样整体复杂度为 `O(U) * O(1 / 1 + 1 / 2 + ... + 1 / U)`, 也即 `O(UlogU)`.
public class Solution
{
    public int CountDifferentSubsequenceGCDs(int[] nums)
    {
        int gcd(int x, int y) => y == 0 ? x : gcd(y, x % y);
        var S = nums.ToHashSet();
        // 考虑一个元素的序列, `gcd` 就是该元素本身.
        var ans = S.Count;
        // 如果有两个以上元素的序列, 如果序列的 `gcd` 为 `k`, 序列最大值至少为 `3*k` (序列为 `2*k, 3*k`). 这样, 枚举到 `n / 3` 后就不可能再出现两个以上元素序列的 `gcd` 了.
        for (int i = 1, n = nums.Max(); i <= n / 3; i++)
        {
            // 为了避免重复计算, 如果该值已经是数组元素了, 就跳过 (已经被作为一个元素的 `gcd` 计数了).
            if (!S.Contains(i))
            {
                for (int g = 0, j = i * 2; j <= n; j += i)
                {
                    if (S.Contains(j) && (g = gcd(g, j)) == i)
                    {
                        ans++;
                        break;
                    }
                }
            }
        }
        return ans;
    }
}
