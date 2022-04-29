/*
 * @lc app=leetcode.cn id=sort-array-by-parity lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [905] 按奇偶排序数组
 *
 * https://leetcode-cn.com/problems/sort-array-by-parity/description/
 *
 * algorithms
 * Easy (69.99%)
 * Total Accepted:    76.1K
 * Total Submissions: 107.7K
 * Testcase Example:  '[3,1,2,4]'
 *
 * 给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。
 * 
 * 返回满足此条件的 任一数组 作为答案。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [3,1,2,4]
 * 输出：[2,4,3,1]
 * 解释：[4,2,3,1]、[2,4,1,3] 和 [4,2,1,3] 也会被视作正确答案。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [0]
 * 输出：[0]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 5000
 * 0 <= nums[i] <= 5000
 * 
 * 
 */
class Solution {
public:
    vector<int> sortArrayByParity(vector<int>& nums) {
        auto n = nums.size();
        int p = 0, q = n - 1;
        while (p < q) {
            while (p < q && nums[p] % 2 == 0) p++;
            while (p < q && nums[q] % 2 == 1) q--;
            if (p < q) swap(nums[p], nums[q]), p++, q--;
        }
        return nums;
    }
};
