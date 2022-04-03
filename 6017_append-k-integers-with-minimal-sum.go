import "sort"

/*
 * Medium
 *
 * 给你一个整数数组 nums 和一个整数 k 。请你向 nums 中追加 k 个 未 出现在 nums 中的、互不相同 的 正 整数，并使结果数组的元素和 最小 。
 *
 * 返回追加到 nums 中的 k 个整数之和。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,4,25,10,25], k = 2
 * 输出：5
 * 解释：在该解法中，向数组中追加的两个互不相同且未出现的正整数是 2 和 3 。
 * nums 最终元素和为 1 + 4 + 25 + 10 + 25 + 2 + 3 = 70 ，这是所有情况中的最小值。
 * 所以追加到数组中的两个整数之和是 2 + 3 = 5 ，所以返回 5 。
 * 示例 2：
 *
 * 输入：nums = [5,6], k = 6
 * 输出：25
 * 解释：在该解法中，向数组中追加的两个互不相同且未出现的正整数是 1 、2 、3 、4 、7 和 8 。
 * nums 最终元素和为 5 + 6 + 1 + 2 + 3 + 4 + 7 + 8 = 36 ，这是所有情况中的最小值。
 * 所以追加到数组中的两个整数之和是 1 + 2 + 3 + 4 + 7 + 8 = 25 ，所以返回 25 。
 *
 *
 * 提示：
 *
 * 1 <= nums.length <= 105
 * 1 <= nums[i], k <= 109
 * 通过次数5,341
 * 提交次数29,915
 */

// 参考: https://leetcode-cn.com/problems/append-k-integers-with-minimal-sum/solution/shi-ma-qing-kuang-tan-xin-shu-xue-qiu-he-nr93/
// 贪心 + 求和公式
func minimalKSum(nums []int, k int) int64 {
	n := len(nums)
	// assert len(n) > 0;
	sort.Ints(nums)
	ans := int64(0)
	start := 1
	// 遍历数组, 计算当前值和数组中 "大于当前值的最小值" 的数据总和.
	//
	// 如果已经超过 `k` 个数, 就只算 `k` 个.
	for i := 0; i < n; i++ {
		if start < nums[i] {
			cnt := k
			if cnt > nums[i]-start {
				cnt = nums[i] - start
			}
			ans += int64((start + start + cnt - 1) * cnt / 2)
			k -= cnt
		}
		start = nums[i] + 1
	}
	// 如果还有多余的, 直接计算后面 `k` 个值的数据总和.
	if k > 0 {
		ans += int64((start + start + k - 1) * k / 2)
	}
	return ans
}
