/*
 * @lc app=leetcode.cn id=burst-balloons lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [312] 戳气球
 *
 * https://leetcode-cn.com/problems/burst-balloons/description/
 *
 * algorithms
 * Hard (69.18%)
 * Total Accepted:    73.2K
 * Total Submissions: 105.8K
 * Testcase Example:  '[3,1,5,8]'
 *
 * 有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
 * 
 * 现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i
 * - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
 * 
 * 求所能获得硬币的最大数量。
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [3,1,5,8]
 * 输出：167
 * 解释：
 * nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
 * coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,5]
 * 输出：10
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * n == nums.length
 * 1 <= n <= 300
 * 0 <= nums[i] <= 100
 * 
 * 
 */
class Solution {
public:
    // 本题注意是开区间, 左右两侧都是开区间. 因此需要在前后把那个 1 加上.
    // 
    // `dp[i][j]` 代表 `(i,j)` 开区间能得到的最大分值. 因此枚举中间的气球 `k`, 就有 `dp[i][j] = max(dp[i][j], dp[i][k] + (nums[i] * nums[k] * nums[j]) + dp[k][j])`.
    int maxCoins(vector<int>& nums) {
        int n = nums.size();
        auto get = [&nums, n](int i) {
            return i < 0 || i >= n ? 1 : nums[i];
        };
        vector<vector<int>> dp(n + 2, vector<int>(n + 2, 0));
        for (int i = n - 1; i >= 0; i--)
            for (int j = i + 2; j <= n + 1; j++)
                for (int k = i + 1; k < j; k++) {
                    int curr = get(i) * get(k) * get(j);
                    dp[i][j] = max(dp[i][j], dp[i][k] + curr + dp[k][j]);
                }
        return dp[0][n + 1];
    }
};
