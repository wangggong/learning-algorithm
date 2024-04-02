/*
 * @lc app=leetcode.cn id=prime-pairs-with-target-sum lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6916] 和等于目标值的质数对
 *
 * https://leetcode.cn/problems/prime-pairs-with-target-sum/description/
 *
 * algorithms
 * Medium (27.34%)
 * Total Accepted:    2.4K
 * Total Submissions: 8.6K
 * Testcase Example:  '10'
 *
 * 给你一个整数 n 。如果两个整数 x 和 y 满足下述条件，则认为二者形成一个质数对：
 * 
 * 
 * 1 <= x <= y <= n
 * x + y == n
 * x 和 y 都是质数
 * 
 * 
 * 请你以二维有序列表的形式返回符合题目要求的所有 [xi, yi] ，列表需要按 xi 的 非递减顺序
 * 排序。如果不存在符合要求的质数对，则返回一个空数组。
 * 
 * 注意：质数是大于 1 的自然数，并且只有两个因子，即它本身和 1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：n = 10
 * 输出：[[3,7],[5,5]]
 * 解释：在这个例子中，存在满足条件的两个质数对。 
 * 这两个质数对分别是 [3,7] 和 [5,5]，按照题面描述中的方式排序后返回。
 * 
 * 
 * 示例 2：
 * 
 * 输入：n = 2
 * 输出：[]
 * 解释：可以证明不存在和为 2 的质数对，所以返回一个空数组。 
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 10^6
 * 
 * 
 */

/*
 * public class Solution
 * {
 *     private const int N = (int)1e6;
 *     private bool[] isPrime;
 * 
 *     public Solution()
 *     {
 *         isPrime = new bool[N + 1];
 *         Array.Fill(isPrime, true);
 *         for (var i = 2; i <= N; i++)
 *         {
 *             if (isPrime[i])
 *             {
 *                 for (var j = 2; i * j <= N; j++)
 *                 {
 *                     isPrime[i * j] = false;
 *                 }
 *             }
 *         }
 *     }
 * 
 *     public IList<IList<int>> FindPrimePairs(int n) => Enumerable.Range(2, n)
 *         .Where(x => n - x >= x && isPrime[x] && isPrime[n - x])
 *         .Select(x => new List<int>{ x, n - x, } as IList<int>)
 *         .ToList();
 * }
 */

public class Solution
{
    const int N = (int)1e6;
    private static bool[] isPrimes;
    private static List<int> primes;

    // 正确的预处理姿势:
    // 利用 `static` 构造方法把 `static field` 赋值, 然后搞起.
    static Solution()
    {
        isPrimes = new bool[N + 1];
        primes = new List<int>();
        Array.Fill(isPrimes, true);
        for (var i = 2; i <= N; i++)
        {
            if (isPrimes[i])
            {
                primes.Add(i);
                for (var j = 2; i * j <= N; j++)
                {
                    isPrimes[i * j] = false;
                }
            }
        }
    }

    public IList<IList<int>> FindPrimePairs(int n) => primes
        .Where(x => n - x >= x && isPrimes[n - x])
        .Select(x => new List<int>{ x, n - x, } as IList<int>)
        .ToList();
}
