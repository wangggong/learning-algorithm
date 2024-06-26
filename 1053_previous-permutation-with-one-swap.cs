/*
 * @lc app=leetcode.cn id=previous-permutation-with-one-swap lang=csharp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1053] 交换一次的先前排列
 *
 * https://leetcode.cn/problems/previous-permutation-with-one-swap/description/
 *
 * algorithms
 * Medium (45.91%)
 * Total Accepted:    11.6K
 * Total Submissions: 24.6K
 * Testcase Example:  '[3,2,1]'
 *
 * 给你一个正整数数组 arr（可能存在重复的元素），请你返回可在 一次交换（交换两数字 arr[i] 和 arr[j] 的位置）后得到的、按字典序排列小于
 * arr 的最大排列。
 * 
 * 如果无法这么操作，就请返回原数组。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：arr = [3,2,1]
 * 输出：[3,1,2]
 * 解释：交换 2 和 1
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：arr = [1,1,5]
 * 输出：[1,1,5]
 * 解释：已经是最小排列
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：arr = [1,9,4,6,7]
 * 输出：[1,7,4,6,9]
 * 解释：交换 9 和 7
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= arr.length <= 10^4
 * 1 <= arr[i] <= 10^4
 * 
 * 
 */
public class Solution
{
    public int[] PrevPermOpt1(int[] arr)
    {
        for (var i = arr.Length - 2; i >= 0; i--)
        {
            if (arr[i] > arr[i + 1])
            {
                var next = i + 1;
                for (var j = arr.Length - 1; j > i; j--)
                {
                    if (arr[j] < arr[i] && arr[next] <= arr[j])
                    {
                        next = j;
                    }
                }
                (arr[i], arr[next]) = (arr[next], arr[i]);
                break;
            }
        }
        return arr;
    }
}
