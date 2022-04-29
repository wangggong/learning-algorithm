/*
 * @lc app=leetcode.cn id=equal-sum-arrays-with-minimum-number-of-operations lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [1775] 通过最少操作次数使数组的和相等
 *
 * https://leetcode-cn.com/problems/equal-sum-arrays-with-minimum-number-of-operations/description/
 *
 * algorithms
 * Medium (48.13%)
 * Total Accepted:    4.9K
 * Total Submissions: 10.1K
 * Testcase Example:  '[1,2,3,4,5,6]\n[1,1,2,2,2,2]'
 *
 * 给你两个长度可能不等的整数数组 nums1 和 nums2 。两个数组中的所有值都在 1 到 6 之间（包含 1 和 6）。
 *
 * 每次操作中，你可以选择 任意 数组中的任意一个整数，将它变成 1 到 6 之间 任意 的值（包含 1 和 6）。
 *
 * 请你返回使 nums1 中所有数的和与 nums2 中所有数的和相等的最少操作次数。如果无法使两个数组的和相等，请返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [1,2,3,4,5,6], nums2 = [1,1,2,2,2,2]
 * 输出：3
 * 解释：你可以通过 3 次操作使 nums1 中所有数的和与 nums2 中所有数的和相等。以下数组下标都从 0 开始。
 * - 将 nums2[0] 变为 6 。 nums1 = [1,2,3,4,5,6], nums2 = [6,1,2,2,2,2] 。
 * - 将 nums1[5] 变为 1 。 nums1 = [1,2,3,4,5,1], nums2 = [6,1,2,2,2,2] 。
 * - 将 nums1[2] 变为 2 。 nums1 = [1,2,2,4,5,1], nums2 = [6,1,2,2,2,2] 。
 *
 *
 * 示例 2：
 *
 * 输入：nums1 = [1,1,1,1,1,1,1], nums2 = [6]
 * 输出：-1
 * 解释：没有办法减少 nums1 的和或者增加 nums2 的和使二者相等。
 *
 *
 * 示例 3：
 *
 * 输入：nums1 = [6,6], nums2 = [1]
 * 输出：3
 * 解释：你可以通过 3 次操作使 nums1 中所有数的和与 nums2 中所有数的和相等。以下数组下标都从 0 开始。
 * - 将 nums1[0] 变为 2 。 nums1 = [2,6], nums2 = [1] 。
 * - 将 nums1[1] 变为 2 。 nums1 = [2,2], nums2 = [1] 。
 * - 将 nums2[0] 变为 4 。 nums1 = [2,2], nums2 = [4] 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length, nums2.length <= 10^5
 * 1 <= nums1[i], nums2[i] <= 6
 *
 *
 */

/*
 * // 现场做的, 很狼狈.
 * import (
 * 	"math"
 * 	"sort"
 * )
 *
 * var dp [][2]int
 *
 * func minOperations(nums1 []int, nums2 []int) int {
 * 	n, m := len(nums1), len(nums2)
 * 	if (n > m && n > m*6) || (m > n && m > n*6) {
 * 		return -1
 * 	}
 * 	sort.Ints(nums1)
 * 	sort.Ints(nums2)
 * 	sum1, sum2 := 0, 0
 * 	for _, n := range nums1 {
 * 		sum1 += n
 * 	}
 * 	for _, n := range nums2 {
 * 		sum2 += n
 * 	}
 * 	if sum1 == sum2 {
 * 		return 0
 * 	}
 * 	if sum2 > sum1 {
 * 		sum1, sum2 = sum2, sum1
 * 		nums1, nums2 = nums2, nums1
 * 		n, m = len(nums1), len(nums2)
 * 	}
 * 	// fmt.Println(sum1, sum2)
 * 	dp = make([][2]int, sum1-sum2+1)
 * 	for i := range dp {
 * 		dp[i][0] = math.MaxInt32
 * 		dp[i][1] = math.MaxInt32
 * 	}
 * 	i := sum1 - sum2
 * 	dp[i][0] = 0
 * 	for j := n - 1; j >= 0; j-- {
 * 		for k := nums1[j]; k > 1; k-- {
 * 			i--
 * 			if i >= 0 {
 * 				dp[i][0] = n - j
 * 			}
 * 		}
 * 	}
 * 	i = 0
 * 	dp[i][1] = 0
 * 	for j := 0; j < m; j++ {
 * 		for k := nums2[j]; k < 6; k++ {
 * 			i++
 * 			if i <= sum1-sum2 {
 * 				dp[i][1] = j + 1
 * 			}
 * 		}
 * 	}
 * 	// fmt.Println(dp)
 * 	ans := math.MaxInt32
 * 	for i := 0; i <= sum1-sum2; i++ {
 * 		// fmt.Println(dp[i][0], dp[i][1])
 * 		ans = min(ans, dp[i][0]+dp[i][1])
 * 	}
 * 	return ans
 * }
 *
 * func min(x, y int) int {
 * 	if x < y {
 * 		return x
 * 	}
 * 	return y
 * }
 */

// 参考: https://leetcode-cn.com/problems/equal-sum-arrays-with-minimum-number-of-operations/solution/tong-guo-zui-shao-cao-zuo-ci-shu-shi-shu-o8no/
//
// 搞个哈希表, 直接贪心就完事了. 好理解还好调.
func minOperations(nums1 []int, nums2 []int) int {
	sum1, sum2 := 0, 0
	for _, n := range nums1 {
		sum1 += n
	}
	for _, n := range nums2 {
		sum2 += n
	}
	if sum1 == sum2 {
		return 0
	}
	if sum1 < sum2 {
		nums1, nums2 = nums2, nums1
		sum1, sum2 = sum2, sum1
	}
	// assert sum1 > sum2
	var freq [6]int
	for _, n := range nums1 {
		freq[n-1]++
	}
	for _, n := range nums2 {
		freq[6-n]++
	}
	ans := 0
	diff := sum1 - sum2
	for k := 5; k > 0 && diff >= 0; k-- {
		// 傻逼! 注意填窟窿时的处理方法!
		if diff > k*freq[k] {
			// 如果当前骰子数不够填窟窿的, 就梭哈.
			diff -= k * freq[k]
			ans += freq[k]
		} else {
			// 如果当前骰子数能把窟窿填满, 就 **向上取整**.
			f := (diff + k - 1) / k
			ans += f
			diff -= f * k
		}
	}
	if diff > 0 {
		return -1
	}
	return ans
}
