/*
 * @lc app=leetcode.cn id=smallest-rotation-with-highest-score lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [798] 得分最高的最小轮调
 *
 * https://leetcode-cn.com/problems/smallest-rotation-with-highest-score/description/
 *
 * algorithms
 * Hard (48.57%)
 * Total Accepted:    5.9K
 * Total Submissions: 10.4K
 * Testcase Example:  '[2,3,1,4,0]'
 *
 * 给你一个数组 nums，我们可以将它按一个非负整数 k 进行轮调，这样可以使数组变为 [nums[k], nums[k + 1], ...
 * nums[nums.length - 1], nums[0], nums[1], ..., nums[k-1]]
 * 的形式。此后，任何值小于或等于其索引的项都可以记作一分。
 *
 *
 * 例如，数组为 nums = [2,4,1,3,0]，我们按 k = 2 进行轮调后，它将变成 [1,3,0,2,4]。这将记为 3 分，因为 1 > 0
 * [不计分]、3 > 1 [不计分]、0 <= 2 [计 1 分]、2 <= 3 [计 1 分]，4 <= 4 [计 1 分]。
 *
 *
 * 在所有可能的轮调中，返回我们所能得到的最高分数对应的轮调下标 k 。如果有多个答案，返回满足条件的最小的下标 k 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,3,1,4,0]
 * 输出：3
 * 解释：
 * 下面列出了每个 k 的得分：
 * k = 0,  nums = [2,3,1,4,0],    score 2
 * k = 1,  nums = [3,1,4,0,2],    score 3
 * k = 2,  nums = [1,4,0,2,3],    score 3
 * k = 3,  nums = [4,0,2,3,1],    score 4
 * k = 4,  nums = [0,2,3,1,4],    score 3
 * 所以我们应当选择 k = 3，得分最高。
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,3,0,2,4]
 * 输出：0
 * 解释：
 * nums 无论怎么变化总是有 3 分。
 * 所以我们将选择最小的 k，即 0。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 0 <= nums[i] < nums.length
 *
 *
 */

const MAXN int = 1e5

var count [MAXN + 5]int

func add(p, q int) {
	count[p]++
	count[q+1]--
}

func bestRotation(nums []int) int {
	n := len(nums)
	for i := range count {
		count[i] = 0
	}
	for i, v := range nums {
		// k = (i+1) % n
		p := (i + 1) % n
		// (i+k-n) % n == v => k = (n-v+i) % n
		q := (n - v + i) % n
		// e.g.: [2,3,1,4,0]
		// for 2 (index 0): p = 1, q = 3
		// p <= q => count[p] = count[p-1]+1 (一定成立), count[q+1] = count[q] - 1 (超过 q 就不成立了)
		if p <= q {
			add(p, q)
		} else {
			add(0, q)
			add(p, n-1)
		}
	}
	// fmt.Printf("%v\n", count[:10])
	for i := 1; i < n; i++ {
		count[i] += count[i-1]
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if count[i] > count[ans] {
			ans = i
		}
	}
	return ans
}
