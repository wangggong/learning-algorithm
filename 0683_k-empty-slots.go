/*
 * Hard
 *
 * n 个灯泡排成一行，编号从 1 到 n 。最初，所有灯泡都关闭。每天 只打开一个 灯泡，直到 n 天后所有灯泡都打开。
 *
 * 给你一个长度为 n 的灯泡数组 blubs ，其中 bulls[i] = x 意味着在第 (i+1) 天，我们会把在位置 x 的灯泡打开，其中 i 从 0 开始，x 从 1 开始。
 *
 * 给你一个整数 k ，请返回恰好有两个打开的灯泡，且它们中间 正好 有 k 个 全部关闭的 灯泡的 最小的天数 。如果不存在这种情况，返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：
 * bulbs = [1,3,2]，k = 1
 * 输出：2
 * 解释：
 * 第一天 bulbs[0] = 1，打开第一个灯泡 [1,0,0]
 * 第二天 bulbs[1] = 3，打开第三个灯泡 [1,0,1]
 * 第三天 bulbs[2] = 2，打开第二个灯泡 [1,1,1]
 * 返回2，因为在第二天，两个打开的灯泡之间恰好有一个关闭的灯泡。
 * 示例 2：
 *
 * 输入：bulbs = [1,2,3]，k = 1
 * 输出：-1
 *
 *
 * 提示：
 *
 * n == bulbs.length
 * 1 <= n <= 2 * 104
 * 1 <= bulbs[i] <= n
 * bulbs 是一个由从 1 到 n 的数字构成的排列
 * 0 <= k <= 2 * 104
 * 通过次数2,922
 * 提交次数6,573
 */

func lowBit(x int) int { return x & (-x) }

type BIT []int

func (b BIT) query(x int) int {
	ans := 0
	for ; x > 0; x -= lowBit(x) {
		ans += b[x]
	}
	return ans
}

// 区间求和, `[l, r]` 闭区间的总和为 `query(r) - query(l-1)`.
func (b BIT) between(l, r int) int {
	return b.query(r) - b.query(l-1)
}

func (b *BIT) add(x, v int) {
	n := len(*b)
	for ; x < n; x += lowBit(x) {
		(*b)[x] += v
	}
}

func NewBIT(n int) *BIT {
	arr := make([]int, n+1)
	b := BIT(arr)
	return &b
}

// 很好的一道 BIT 题. 除了发现每次开灯只需要验证前后两个区间外, 还要对边界的划分需要非常清楚.
func kEmptySlots(bulbs []int, k int) int {
	n := len(bulbs)
	bit := NewBIT(n)
	on := make([]bool, n+1)
	for i, b := range bulbs {
		// `b` 号灯前面 **间隔** `k` 个灯, 应该到 `b-k-1` 号灯.
		//
		// X . . . X
		// 1 2 3 4 5
		//
		// b = 5, k = 3
		// 因此 `bit.between(b-k, b-1) == 0`, 即 `[b-k, b-1]` 闭区间亮灯总数为 0.
		if b-k-1 > 0 && on[b-k-1] && bit.between(b-k, b-1) == 0 {
			return i + 1
		}
		// `b` 号灯后面 **间隔** `k` 个灯, 应该到 `b+k+1` 号灯.
		//
		// X . . . X
		// 1 2 3 4 5
		//
		// b = 1, k = 3
		// 因此 `bit.between(b+1, b+k) == 0`, 即 `[b+1, b+k]` 闭区间亮灯总数为 0.
		if b+k+1 <= n && on[b+k+1] && bit.between(b+1, b+k) == 0 {
			return i + 1
		}
		on[b] = true
		bit.add(b, 1)
	}
	return -1
}
