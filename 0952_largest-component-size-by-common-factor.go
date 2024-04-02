/*
 * @lc app=leetcode.cn id=largest-component-size-by-common-factor lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [952] 按公因数计算最大组件大小
 *
 * https://leetcode.cn/problems/largest-component-size-by-common-factor/description/
 *
 * algorithms
 * Hard (37.60%)
 * Total Accepted:    11.9K
 * Total Submissions: 24.9K
 * Testcase Example:  '[4,6,15,35]'
 *
 * 给定一个由不同正整数的组成的非空数组 nums ，考虑下面的图：
 *
 *
 * 有 nums.length 个节点，按从 nums[0] 到 nums[nums.length - 1] 标记；
 * 只有当 nums[i] 和 nums[j] 共用一个大于 1 的公因数时，nums[i] 和 nums[j]之间才有一条边。
 *
 *
 * 返回 图中最大连通组件的大小 。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：nums = [4,6,15,35]
 * 输出：4
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入：nums = [20,50,9,63]
 * 输出：2
 *
 *
 * 示例 3：
 *
 *
 *
 *
 * 输入：nums = [2,3,6,7,4,12,21,39]
 * 输出：8
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 2 * 10^4
 * 1 <= nums[i] <= 10^5
 * nums 中所有值都 不同
 *
 *
 */

// 看了提示想起来可以先分解质因数.
//
// 然后发现分解质因数, 忘了!
var fa []int
var factors map[int][]int

func query(k int) int {
	if fa[k] != k {
		fa[k] = query(fa[k])
	}
	return fa[k]
}

func merge(p, q int) {
	fa[query(p)] = query(q)
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func factor(k, n int) {
	for i := 2; i <= n/i; i++ {
		ok := false
		for ; n%i == 0; n /= i {
			ok = true
		}
		if ok {
			factors[i] = append(factors[i], k)
		}
	}
	if n > 1 {
		factors[n] = append(factors[n], k)
	}
	return
}

func largestComponentSize(nums []int) int {
	n := len(nums)
	fa = make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	factors = make(map[int][]int)
	for i, n := range nums {
		factor(i, n)
	}
	for _, f := range factors {
		for i, m := 1, len(f); i < m; i++ {
			merge(f[0], f[i])
		}
	}
	size := make(map[int]int)
	for i := 0; i < n; i++ {
		size[query(i)]++
	}
	ans := 0
	for _, s := range size {
		ans = max(ans, s)
	}
	return ans
}
