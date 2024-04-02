/*
 * @lc app=leetcode.cn id=maximum-equal-frequency lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1224] 最大相等频率
 *
 * https://leetcode.cn/problems/maximum-equal-frequency/description/
 *
 * algorithms
 * Hard (33.46%)
 * Total Accepted:    18.9K
 * Total Submissions: 45.5K
 * Testcase Example:  '[2,2,1,1,5,3,3,5]'
 *
 * 给你一个正整数数组 nums，请你帮忙从该数组中找出能满足下面要求的 最长 前缀，并返回该前缀的长度：
 *
 *
 * 从前缀中 恰好删除一个 元素后，剩下每个数字的出现次数都相同。
 *
 *
 * 如果删除这个元素后没有剩余元素存在，仍可认为每个数字都具有相同的出现次数（也就是 0 次）。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,2,1,1,5,3,3,5]
 * 输出：7
 * 解释：对于长度为 7 的子数组 [2,2,1,1,5,3,3]，如果我们从中删去 nums[4] = 5，就可以得到
 * [2,2,1,1,3,3]，里面每个数字都出现了两次。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,1,1,2,2,2,3,3,3,4,4,4,5]
 * 输出：13
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^5
 *
 *
 */
func maxEqualFreq(nums []int) (ans int) {
	count, freq, mxFreq := make(map[int]int), make(map[int]int), 0
	for i, n := range nums {
		if count[n] > 0 {
			freq[count[n]]--
		}
		count[n]++
		freq[count[n]]++
		mxFreq = max(mxFreq, count[n])
		if mxFreq == 1 /* 所有数据均只出现了一次 */ || mxFreq*freq[mxFreq]+1 == i+1 && freq[1] == 1 /* 一组数据只出现了一次, 其余均出现了 `mxFreq` 次 */ || mxFreq*freq[mxFreq]+(mxFreq-1)*freq[mxFreq-1] == i+1 && freq[mxFreq] == 1 /* 一组数据出现了 `mxFreq` 次, 其余数据均出现了 `mxFreq-1` 次 */ {
			ans = i + 1
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
