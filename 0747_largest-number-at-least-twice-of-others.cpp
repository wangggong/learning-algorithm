/*
 * @lc app=leetcode.cn id=largest-number-at-least-twice-of-others lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [747] 至少是其他数字两倍的最大数
 *
 * https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others/description/
 *
 * algorithms
 * Easy (45.90%)
 * Total Accepted:    79.6K
 * Total Submissions: 173.4K
 * Testcase Example:  '[3,6,1,0]'
 *
 * 给你一个整数数组 nums ，其中总是存在 唯一的 一个最大整数 。
 * 
 * 请你找出数组中的最大元素并检查它是否 至少是数组中每个其他数字的两倍 。如果是，则返回 最大元素的下标 ，否则返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [3,6,1,0]
 * 输出：1
 * 解释：6 是最大的整数，对于数组中的其他整数，6 至少是数组中其他元素的两倍。6 的下标是 1 ，所以返回 1 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3,4]
 * 输出：-1
 * 解释：4 没有超过 3 的两倍大，所以返回 -1 。
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1]
 * 输出：0
 * 解释：因为不存在其他数字，所以认为现有数字 1 至少是其他数字的两倍。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 50
 * 0 <= nums[i] <= 100
 * nums 中的最大元素是唯一的
 * 
 * 
 */
class Solution {
public:
  /*
   * int dominantIndex(vector<int>& nums) {
   *   auto F = -1, S = -1;
   *   auto n = nums.size();
   *   if (n == 1) return 0;
   *   for (auto i = 0; i < n; i++) {
   *     if (F == -1 || nums[i] > nums[F]) F = i;
   *
   *   }
   *   for (auto i = 0; i < n; i++) {
   *     if (i == F) continue;
   *     if (S == -1 || nums[i] > nums[S]) S = i;
   *   }
   *   return nums[F] >= 2*nums[S] ? F : -1;
   * }
  */

  int dominantIndex(vector<int>& nums) {
    auto F = -1, S = -1, index = 0;
    auto n = nums.size();
    for (auto i = 0; i < n; i++) {
      if (nums[i] > F) {
        S = F;
        F = nums[i];
        index = i;
      } else if (nums[i] > S) {
        S = nums[i];
      }
    }
    return F >= S * 2 ? index : -1;
  }
};
