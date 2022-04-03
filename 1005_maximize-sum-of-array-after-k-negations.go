import "sort"

/*
 * @lc app=leetcode.cn id=maximize-sum-of-array-after-k-negations lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1005] K 次取反后最大化的数组和
 *
 * https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations/description/
 *
 * algorithms
 * Easy (52.01%)
 * Total Accepted:    73.2K
 * Total Submissions: 141.3K
 * Testcase Example:  '[4,2,3]\n1'
 *
 * 给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
 *
 *
 * 选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
 *
 *
 * 重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
 *
 * 以这种方式修改数组后，返回数组 可能的最大和 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [4,2,3], k = 1
 * 输出：5
 * 解释：选择下标 1 ，nums 变为 [4,-2,3] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [3,-1,0,2], k = 3
 * 输出：6
 * 解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [2,-3,-1,5,-4], k = 2
 * 输出：13
 * 解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^4
 * -100 <= nums[i] <= 100
 * 1 <= k <= 10^4
 *
 *
 */
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	neg := 0
	for _, n := range nums {
		if n < 0 {
			neg++
		}
	}
	if neg >= k {
		ans := 0
		for i, n := range nums {
			if i < k {
				ans += abs(n)
			} else {
				ans += n
			}
		}
		return ans
	}
	// assert neg < k;
	ans := 0
	minAbs := abs(nums[0])
	for _, n := range nums {
		an := abs(n)
		ans += an
		if an < minAbs {
			minAbs = an
		}
		// fmt.Println(ans, n, minAbs)
	}
	if (k-neg)%2 > 0 {
		ans -= 2 * minAbs
	}
	return ans
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
