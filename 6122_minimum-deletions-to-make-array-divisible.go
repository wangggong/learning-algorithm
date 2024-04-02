import "sort"

/*
 * @lc app=leetcode.cn id=minimum-deletions-to-make-array-divisible lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [6122] 使数组可以被整除的最少删除次数
 *
 * https://leetcode.cn/problems/minimum-deletions-to-make-array-divisible/description/
 *
 * algorithms
 * Hard (49.56%)
 * Total Accepted:    4.6K
 * Total Submissions: 9.3K
 * Testcase Example:  '[2,3,2,4,3]\n[9,6,9,3,15]'
 *
 * 给你两个正整数数组 nums 和 numsDivide 。你可以从 nums 中删除任意数目的元素。
 *
 * 请你返回使 nums 中 最小 元素可以整除 numsDivide 中所有元素的 最少 删除次数。如果无法得到这样的元素，返回 -1 。
 *
 * 如果 y % x == 0 ，那么我们说整数 x 整除 y 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [2,3,2,4,3], numsDivide = [9,6,9,3,15]
 * 输出：2
 * 解释：
 * [2,3,2,4,3] 中最小元素是 2 ，它无法整除 numsDivide 中所有元素。
 * 我们从 nums 中删除 2 个大小为 2 的元素，得到 nums = [3,4,3] 。
 * [3,4,3] 中最小元素为 3 ，它可以整除 numsDivide 中所有元素。
 * 可以证明 2 是最少删除次数。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [4,3,6], numsDivide = [8,2,6,10]
 * 输出：-1
 * 解释：
 * 我们想 nums 中的最小元素可以整除 numsDivide 中的所有元素。
 * 没有任何办法可以达到这一目的。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length, numsDivide.length <= 10^5
 * 1 <= nums[i], numsDivide[i] <= 10^9
 *
 *
 */

/*
 * import "math"
 *
 * func minOperations(nums []int, numsDivide []int) int {
 * 	k := getGcd(numsDivide)
 * 	cand := math.MaxInt32
 * 	for _, n := range nums {
 * 		if k%n == 0 {
 * 			if cand < 0 {
 * 				cand = n
 * 				continue
 * 			}
 * 			cand = min(cand, n)
 * 		}
 * 	}
 * 	if cand == math.MaxInt32 {
 * 		return -1
 * 	}
 * 	ans := 0
 * 	for _, n := range nums {
 * 		if n < cand {
 * 			ans++
 * 		}
 * 	}
 * 	return ans
 * }
 */

func minOperations(nums []int, numsDivide []int) int {
	k := getGcd(numsDivide)
	sort.Ints(nums)
	ans := -1
	for i, n := range nums {
		if k%n == 0 {
			ans = i
			break
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getGcd(arr []int) int {
	ans := arr[0]
	for _, a := range arr {
		ans = gcd(ans, a)
	}
	return ans
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
