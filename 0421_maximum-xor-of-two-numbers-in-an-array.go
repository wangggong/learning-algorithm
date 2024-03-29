/*
 * @lc app=leetcode.cn id=maximum-xor-of-two-numbers-in-an-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [421] 数组中两个数的最大异或值
 *
 * https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/description/
 *
 * algorithms
 * Medium (62.13%)
 * Total Accepted:    35.3K
 * Total Submissions: 56.8K
 * Testcase Example:  '[3,10,5,25,2,8]'
 *
 * 给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
 *
 * 进阶：你可以在 O(n) 的时间解决这个问题吗？
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,10,5,25,2,8]
 * 输出：28
 * 解释：最大运算结果是 5 XOR 25 = 28.
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [2,4]
 * 输出：6
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums = [8,10,2]
 * 输出：10
 *
 *
 * 示例 5：
 *
 *
 * 输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
 * 输出：127
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * 0 <= nums[i] <= 2^31 - 1
 *
 *
 *
 *
 */
const HB = 30

func findMaximumXOR(N []int) int {
	x := 0
	for b := HB; b >= 0; b-- {
		M := make(map[int]bool)
		for _, n := range N {
			M[n>>b] = true
		}
		xn := (x << 1) + 1
		ok := false
		for _, n := range N {
			if M[n>>b^xn] {
				ok = true
				break
			}
		}
		if ok {
			x = xn
		} else {
			x = xn - 1
		}
	}
	return x
}
