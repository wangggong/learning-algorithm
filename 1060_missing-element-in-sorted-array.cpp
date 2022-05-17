/*
 * Medium
 * 
 * 现有一个按 升序 排列的整数数组 nums ，其中每个数字都 互不相同 。
 * 
 * 给你一个整数 k ，请你找出并返回从数组最左边开始的第 k 个缺失数字。
 * 
 *  
 * 
 * 示例 1：
 * 
 * 输入：nums = [4,7,9,10], k = 1
 * 输出：5
 * 解释：第一个缺失数字为 5 。
 * 示例 2：
 * 
 * 输入：nums = [4,7,9,10], k = 3
 * 输出：8
 * 解释：缺失数字有 [5,6,8,...]，因此第三个缺失数字为 8 。
 * 示例 3：
 * 
 * 输入：nums = [1,2,4], k = 3
 * 输出：6
 * 解释：缺失数字有 [3,5,6,7,...]，因此第三个缺失数字为 6 。
 *  
 * 
 * 提示：
 * 
 * 1 <= nums.length <= 5 * 104
 * 1 <= nums[i] <= 107
 * nums 按 升序 排列，其中所有元素 互不相同 。
 * 1 <= k <= 108
 *  
 * 
 * 进阶：你可以设计一个对数时间复杂度（即，O(log(n))）的解决方案吗？
 * 
 * 通过次数7,898
 * 提交次数14,799
 */ 

class Solution {
    int lower_bound(vector<int> &nums, int k) {
        int p = -1, q = nums.size() - 1;
        while (p < q) {
            int mid = (p + q + 1) >> 1;
            if (nums[mid] <= k) p = mid;
            else q = mid - 1;
        }
        return p;
    }
public:
    /*
     * int missingElement(vector<int>& nums, int k) {
     *     int n = nums.size();
     *     int tot = (nums[n - 1] - nums[0]) - n + 1;
     *     if (k > tot) return nums[n - 1] + (k - tot);
     *     int p = nums[0], q = nums[n - 1];
     *     while (p < q) {
     *         int mid = (p + q) >> 1;
     *         int ind = lower_bound(nums, mid);
     *         if (mid - nums[0] >= ind + k) q = mid;
     *         else p = mid + 1;
     *     }
     *     return p;
     * }
     */

    int missingElement(vector<int> &nums, int k) {
        int n = nums.size(), p = 0, q = n;
        while (p < q) {
            int mid = p + q >> 1;
            if (nums[mid] - nums[0] >= mid + k)
                q = mid;
            else
                p = mid + 1;
        }
        return nums[0] + k + p - 1;
    }
};