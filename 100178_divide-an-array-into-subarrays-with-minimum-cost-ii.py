#
# @lc app=leetcode.cn id=divide-an-array-into-subarrays-with-minimum-cost-ii lang=python3
#
# NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
#
# [100178] 将数组分成最小总代价的子数组 II
#
# https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-ii/description/
#
# algorithms
# Hard (30.15%)
# Total Accepted:    1.3K
# Total Submissions: 3.2K
# Testcase Example:  '[1,3,2,6,4,2]\n3\n3'
#
# 给你一个下标从 0 开始长度为 n 的整数数组 nums 和两个 正 整数 k 和 dist 。
# 
# 一个数组的 代价 是数组中的 第一个 元素。比方说，[1,2,3] 的代价为 1 ，[3,4,1] 的代价为 3 。
# 
# 你需要将 nums 分割成 k 个 连续且互不相交 的子数组，满足 第二 个子数组与第 k 个子数组中第一个元素的下标距离 不超过 dist
# 。换句话说，如果你将 nums 分割成子数组 nums[0..(i1 - 1)], nums[i1..(i2 - 1)], ...,
# nums[ik-1..(n - 1)] ，那么它需要满足 ik-1 - i1 <= dist 。
# 
# 请你返回这些子数组的 最小 总代价。
# 
# 
# 
# 示例 1：
# 
# 
# 输入：nums = [1,3,2,6,4,2], k = 3, dist = 3
# 输出：5
# 解释：将数组分割成 3 个子数组的最优方案是：[1,3] ，[2,6,4] 和 [2] 。这是一个合法分割，因为 ik-1 - i1 等于 5 - 2 =
# 3 ，等于 dist 。总代价为 nums[0] + nums[2] + nums[5] ，也就是 1 + 2 + 2 = 5 。
# 5 是分割成 3 个子数组的最小总代价。
# 
# 
# 示例 2：
# 
# 
# 输入：nums = [10,1,2,2,2,1], k = 4, dist = 3
# 输出：15
# 解释：将数组分割成 4 个子数组的最优方案是：[10] ，[1] ，[2] 和 [2,2,1] 。这是一个合法分割，因为 ik-1 - i1 等于 3 -
# 1 = 2 ，小于 dist 。总代价为 nums[0] + nums[1] + nums[2] + nums[3] ，也就是 10 + 1 + 2 +
# 2 = 15 。
# 分割 [10] ，[1] ，[2,2,2] 和 [1] 不是一个合法分割，因为 ik-1 和 i1 的差为 5 - 1 = 4 ，大于 dist 。
# 15 是分割成 4 个子数组的最小总代价。
# 
# 
# 示例 3：
# 
# 
# 输入：nums = [10,8,18,9], k = 3, dist = 1
# 输出：36
# 解释：将数组分割成 4 个子数组的最优方案是：[10] ，[8] 和 [18,9] 。这是一个合法分割，因为 ik-1 - i1 等于 2 - 1 = 1
# ，等于 dist 。总代价为 nums[0] + nums[1] + nums[2] ，也就是 10 + 8 + 18 = 36 。
# 分割 [10] ，[8,18] 和 [9] 不是一个合法分割，因为 ik-1 和 i1 的差为 3 - 1 = 2 ，大于 dist 。
# 36 是分割成 3 个子数组的最小总代价。
# 
# 
# 
# 
# 提示：
# 
# 
# 3 <= n <= 10^5
# 1 <= nums[i] <= 10^9
# 3 <= k <= n
# k - 2 <= dist <= n - 2
# 
# 
#

# 参考: https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-ii/solutions/2614067/liang-ge-you-xu-ji-he-wei-hu-qian-k-1-xi-zdzx/
#
# SortedList 确实比较方便地解决了有序数组的 CRUD 问题. 它是不是优先队列的超集?
# 相关题目: 480
from sortedcontainers import SortedList


class Solution:
    def minimumCost(self, nums: List[int], k: int, dist: int) -> int:
        k, dist, first, nums = k - 1, dist + 1, nums[0], nums[1:]
        L, R, cur = SortedList(), SortedList(), [0]
        def addL(v):
            L.add(v)
            cur[0] += v
        def popL():
            v = L.pop()
            cur[0] -= v
            return v
        def removeL(v):
            L.remove(v)
            cur[0] -= v
        for n in nums[:dist]:
            addL(n)
        while len(L) > k:
            R.add(popL())
        ans = cur[0]
        for i, n in enumerate(nums[dist:]):
            if nums[i] in L:
                removeL(nums[i])
            else:
                R.remove(nums[i])
            if n < L[-1]:
                addL(n)
            else:
                R.add(n)
            if len(L) == k - 1:
                addL(R.pop(0))
            if len(L) == k + 1:
                R.add(popL())
            ans = min(ans, cur[0])
        return first + ans
