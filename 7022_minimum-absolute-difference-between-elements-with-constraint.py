#
# @lc app=leetcode.cn id=minimum-absolute-difference-between-elements-with-constraint lang=python3
#
# NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
#
# [7022] 限制条件下元素之间的最小绝对差
#
# https://leetcode.cn/problems/minimum-absolute-difference-between-elements-with-constraint/description/
#
# algorithms
# Medium (24.76%)
# Total Accepted:    2.7K
# Total Submissions: 10.8K
# Testcase Example:  '[4,3,2,4]\n2'
#
# 给你一个下标从 0 开始的整数数组 nums 和一个整数 x 。
# 
# 请你找到数组中下标距离至少为 x 的两个元素的 差值绝对值 的 最小值 。
# 
# 换言之，请你找到两个下标 i 和 j ，满足 abs(i - j) >= x 且 abs(nums[i] - nums[j]) 的值最小。
# 
# 请你返回一个整数，表示下标距离至少为 x 的两个元素之间的差值绝对值的 最小值 。
# 
# 
# 
# 示例 1：
# 
# 
# 输入：nums = [4,3,2,4], x = 2
# 输出：0
# 解释：我们选择 nums[0] = 4 和 nums[3] = 4 。
# 它们下标距离满足至少为 2 ，差值绝对值为最小值 0 。
# 0 是最优解。
# 
# 
# 示例 2：
# 
# 
# 输入：nums = [5,3,2,10,15], x = 1
# 输出：1
# 解释：我们选择 nums[1] = 3 和 nums[2] = 2 。
# 它们下标距离满足至少为 1 ，差值绝对值为最小值 1 。
# 1 是最优解。
# 
# 
# 示例 3：
# 
# 
# 输入：nums = [1,2,3,4], x = 3
# 输出：3
# 解释：我们选择 nums[0] = 1 和 nums[3] = 4 。
# 它们下标距离满足至少为 3 ，差值绝对值为最小值 3 。
# 3 是最优解。
# 
# 
# 
# 
# 提示：
# 
# 
# 1 <= nums.length <= 10^5
# 1 <= nums[i] <= 10^9
# 0 <= x < nums.length
# 
# 
#

# 参考: https://leetcode.cn/problems/minimum-absolute-difference-between-elements-with-constraint/solutions/2386043/ping-heng-shu-shuang-zhi-zhen-pythonjava-50eh/

from sortedcontainers import SortedList


class Solution:
    def minAbsoluteDifference(self, nums: List[int], k: int) -> int:
        ans = inf
        sl = SortedList([-inf, inf])
        for x, y in zip(nums, nums[k:]):
            sl.add(x)
            i = sl.bisect_left(y)
            ans = min(ans, sl[i] - y, y - sl[i - 1])
        return ans
