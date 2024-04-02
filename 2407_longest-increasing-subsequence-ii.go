/*
 * @lc app=leetcode.cn id=longest-increasing-subsequence-ii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2407] 最长递增子序列 II
 *
 * https://leetcode.cn/problems/longest-increasing-subsequence-ii/description/
 *
 * algorithms
 * Hard (20.01%)
 * Total Accepted:    2.1K
 * Total Submissions: 10.2K
 * Testcase Example:  '[4,2,1,4,3,4,5,8,15]\n3'
 *
 * 给你一个整数数组 nums 和一个整数 k 。
 *
 * 找到 nums 中满足以下要求的最长子序列：
 *
 *
 * 子序列 严格递增
 * 子序列中相邻元素的差值 不超过 k 。
 *
 *
 * 请你返回满足上述要求的 最长子序列 的长度。
 *
 * 子序列 是从一个数组中删除部分元素后，剩余元素不改变顺序得到的数组。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [4,2,1,4,3,4,5,8,15], k = 3
 * 输出：5
 * 解释：
 * 满足要求的最长子序列是 [1,3,4,5,8] 。
 * 子序列长度为 5 ，所以我们返回 5 。
 * 注意子序列 [1,3,4,5,8,15] 不满足要求，因为 15 - 8 = 7 大于 3 。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [7,4,5,1,8,12,4,7], k = 5
 * 输出：4
 * 解释：
 * 满足要求的最长子序列是 [4,5,8,12] 。
 * 子序列长度为 4 ，所以我们返回 4 。
 *
 *
 * 示例 3：
 *
 * 输入：nums = [1,5], k = 1
 * 输出：1
 * 解释：
 * 满足要求的最长子序列是 [1] 。
 * 子序列长度为 1 ，所以我们返回 1 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i], k <= 10^5
 *
 *
 */

// 参考: https://leetcode.cn/problems/longest-increasing-subsequence-ii/solution/zhi-yu-xian-duan-shu-pythonjavacgo-by-en-p1gz/
type seg []struct{ l, r, max int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) modify(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].max = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.modify(o<<1, i, val)
	} else {
		t.modify(o<<1|1, i, val)
	}
	t[o].max = max(t[o<<1].max, t[o<<1|1].max)
}

// 返回区间 [l,r] 内的最大值
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].max
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return max(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func lengthOfLIS(nums []int, k int) int {
	mx := 0
	for _, x := range nums {
		mx = max(mx, x)
	}
	t := make(seg, mx*4)
	t.build(1, 1, mx)
	for _, x := range nums {
		if x == 1 {
			t.modify(1, 1, 1)
		} else {
			t.modify(1, x, 1+t.query(1, max(x-k, 1), x-1))
		}
	}
	return t[1].max
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
