/*
 * @lc app=leetcode.cn id=number-of-great-partitions lang=cpp
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6272] 好分区的数目
 *
 * https://leetcode.cn/problems/number-of-great-partitions/description/
 *
 * algorithms
 * Hard (33.93%)
 * Total Accepted:    723
 * Total Submissions: 1.9K
 * Testcase Example:  '[1,2,3,4]\n4'
 *
 * 给你一个正整数数组 nums 和一个整数 k 。
 * 
 * 分区 的定义是：将数组划分成两个有序的 组 ，并满足每个元素 恰好 存在于 某一个 组中。如果分区中每个组的元素和都大于等于 k
 * ，则认为分区是一个好分区。
 * 
 * 返回 不同 的好分区的数目。由于答案可能很大，请返回对 10^9 + 7 取余 后的结果。
 * 
 * 如果在两个分区中，存在某个元素 nums[i] 被分在不同的组中，则认为这两个分区不同。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3,4], k = 4
 * 输出：6
 * 解释：好分区的情况是 ([1,2,3], [4]), ([1,3], [2,4]), ([1,4], [2,3]), ([2,3], [1,4]),
 * ([2,4], [1,3]) 和 ([4], [1,2,3]) 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,3,3], k = 4
 * 输出：0
 * 解释：数组中不存在好分区。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [6,6], k = 2
 * 输出：2
 * 解释：可以将 nums[0] 放入第一个分区或第二个分区中。
 * 好分区的情况是 ([6], [6]) 和 ([6], [6]) 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length, k <= 1000
 * 1 <= nums[i] <= 10^9
 * 
 * 
 */

typedef long long ll;
const ll MOD = 1e9 + 7;
const ll N = 1000;
ll dp[N + 10];

ll POW(ll n, ll k) {
    ll ans = 1;
    for (; k; k >>= 1)
        ans = k & 1 ? (ans * n) % MOD : ans, n = (n * n) % MOD;
    return ans;
}

class Solution {
public:
    int countPartitions(vector<int>& nums, int k) {
        memset(dp, 0, sizeof dp);
        ll tot = 0;
        for (auto n : nums)
            tot += n;
        if (tot < k * 2)
            return 0;
        dp[0] = 1;
        for (auto n : nums)
            for (ll i = k - 1; i >= n; i--)
                dp[i] = (dp[i] + dp[i - n]) % MOD;
        tot = 0;
        for (ll i = 0; i < k; i++)
            tot = (tot + dp[i]) % MOD;
        return (POW(2, nums.size()) - tot * 2 % MOD + MOD) % MOD;
    }
};
