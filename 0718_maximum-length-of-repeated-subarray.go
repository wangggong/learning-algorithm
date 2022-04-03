/*
 * @lc app=leetcode.cn id=maximum-length-of-repeated-subarray lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [718] 最长重复子数组
 *
 * https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/description/
 *
 * algorithms
 * Medium (56.56%)
 * Total Accepted:    99K
 * Total Submissions: 174.9K
 * Testcase Example:  '[1,2,3,2,1]\n[3,2,1,4,7]'
 *
 * 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
 * 输出：3
 * 解释：长度最长的公共子数组是 [3,2,1] 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums1.length, nums2.length <= 1000
 * 0 <= nums1[i], nums2[i] <= 100
 *
 *
 */
const MAXN = 1e3 + 10

var dp [MAXN][MAXN]int

var n, m int

/*
 * // dp
 * func findLength(nums1 []int, nums2 []int) int {
 * 	n, m := len(nums1), len(nums2)
 * 	// assert n > 0 && n <= 1e3;
 * 	// assert m > 0 && m <= 1e3;
 * 	for i := 0; i < MAXN; i++ {
 * 		for j := 0; j < MAXN; j++ {
 * 			dp[i][j] = 0
 * 		}
 * 	}
 * 	ans := 0
 * 	for i := 0; i < n; i++ {
 * 		if nums1[i] == nums2[0] {
 * 			dp[i][0] = 1
 * 			ans = max(ans, dp[i][0])
 * 		}
 * 	}
 * 	for i := 0; i < m; i++ {
 * 		if nums1[0] == nums2[i] {
 * 			dp[0][i] = 1
 * 			ans = max(ans, dp[0][i])
 * 		}
 * 	}
 * 	for i := 1; i < n; i++ {
 * 		for j := 1; j < m; j++ {
 * 			if nums1[i] == nums2[j] {
 * 				dp[i][j] = max(dp[i-1][j-1]+1, 1)
 * 				ans = max(ans, dp[i][j])
 * 			}
 * 		}
 * 	}
 * 	return ans
 * }
 *
 */
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 滑动窗口
func findLength(nums1 []int, nums2 []int) int {
	n, m = len(nums1), len(nums2)
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, maxLength(nums1, nums2, i, 0))
	}
	for j := 0; j < m; j++ {
		ans = max(ans, maxLength(nums1, nums2, 0, j))
	}
	return ans
}

func maxLength(nums1, nums2 []int, p, q int) int {
	ans := 0
	cur := 0
	for i := 0; p+i < n && q+i < m; i++ {
		if nums1[p+i] == nums2[q+i] {
			cur++
			ans = max(ans, cur)
		} else {
			cur = 0
		}
	}
	return ans
}

