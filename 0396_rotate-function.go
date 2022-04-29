/*
 * @lc app=leetcode.cn id=rotate-function lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [396] 旋转函数
 *
 * https://leetcode-cn.com/problems/rotate-function/description/
 *
 * algorithms
 * Medium (44.86%)
 * Total Accepted:    17.5K
 * Total Submissions: 35.8K
 * Testcase Example:  '[4,3,2,6]'
 *
 * 给定一个长度为 n 的整数数组 nums 。
 *
 * 假设 arrk 是数组 nums 顺时针旋转 k 个位置后的数组，我们定义 nums 的 旋转函数  F 为：
 *
 *
 * F(k) = 0 * arrk[0] + 1 * arrk[1] + ... + (n - 1) * arrk[n - 1]
 *
 *
 * 返回 F(0), F(1), ..., F(n-1)中的最大值 。
 *
 * 生成的测试用例让答案符合 32 位 整数。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [4,3,2,6]
 * 输出: 26
 * 解释:
 * F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
 * F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
 * F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
 * F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26
 * 所以 F(0), F(1), F(2), F(3) 中的最大值是 F(3) = 26 。
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums = [100]
 * 输出: 0
 *
 *
 *
 *
 * 提示:
 *
 *
 * n == nums.length
 * 1 <= n <= 10^5
 * -100 <= nums[i] <= 100
 *
 *
 */
func maxRotateFunction(nums []int) int {
	// `i = 0 => i = 1`
	// `diff = \Sigma_{i=1}^{n}{i * arr[i]} - \Sigma{i=0}^{n-1}{i * arr[i]} = \Sigma_{i=0}^{n-1}{((i+1) % n) * arr[(i+1) % n] - i * arr[i]} = \Sigma_{i=0}^{n-1}{arr[i]} - n * arr[n-1]`
	n := len(nums)
	if n == 1 {
		return 0
	}
	tot := 0
	cur := 0
	for i, n := range nums {
		tot += n
		cur += i * n
	}
	ans := cur
	for i := 0; i < n; i++ {
		cur = cur + tot - n*nums[n-1-i]
		ans = max(ans, cur)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
