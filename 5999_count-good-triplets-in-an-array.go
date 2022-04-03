/*
 * @lc app=leetcode.cn id=count-good-triplets-in-an-array lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [5999] 统计数组中好三元组数目
 *
 * https://leetcode-cn.com/problems/count-good-triplets-in-an-array/description/
 *
 * algorithms
 * Hard (13.41%)
 * Total Accepted:    556
 * Total Submissions: 4.2K
 * Testcase Example:  '[2,0,1,3]\n[0,1,2,3]'
 *
 * 给你两个下标从 0 开始且长度为 n 的整数数组 nums1 和 nums2 ，两者都是 [0, 1, ..., n - 1] 的 排列 。
 *
 * 好三元组 指的是 3 个 互不相同 的值，且它们在数组 nums1 和 nums2 中出现顺序保持一致。换句话说，如果我们将 pos1v 记为值 v 在
 * nums1 中出现的位置，pos2v 为值 v 在 nums2 中的位置，那么一个好三元组定义为 0 <= x, y, z <= n - 1 ，且
 * pos1x < pos1y < pos1z 和 pos2x < pos2y < pos2z 都成立的 (x, y, z) 。
 *
 * 请你返回好三元组的 总数目 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [2,0,1,3], nums2 = [0,1,2,3]
 * 输出：1
 * 解释：
 * 总共有 4 个三元组 (x,y,z) 满足 pos1x < pos1y < pos1z ，分别是 (2,0,1) ，(2,0,3) ，(2,1,3) 和
 * (0,1,3) 。
 * 这些三元组中，只有 (0,1,3) 满足 pos2x < pos2y < pos2z 。所以只有 1 个好三元组。
 *
 *
 * 示例 2：
 *
 * 输入：nums1 = [4,0,1,3,2], nums2 = [4,1,0,2,3]
 * 输出：4
 * 解释：总共有 4 个好三元组 (4,0,3) ，(4,0,2) ，(4,1,3) 和 (4,1,2) 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == nums1.length == nums2.length
 * 3 <= n <= 10^5
 * 0 <= nums1[i], nums2[i] <= n - 1
 * nums1 和 nums2 是 [0, 1, ..., n - 1] 的排列。
 *
 *
 */

// 当时卡在顺序三元组数量上了; 想到了枚举中间值, 但没想到用 BIT 的前缀和.
// 注意:
// 1. query 的中止条件是 index > 0, 不能取 "=", 会死循环;
// 2. 离散化后起始 val 是 0, 入数组的时候别忘了 "+1", 查询的时候应该查询 "val+1-1 = val" 的;
//
// 参考:
// - https://leetcode-cn.com/problems/count-good-triplets-in-an-array/solution/deng-jie-zhuan-huan-shu-zhuang-shu-zu-by-xmyd/
// - https://leetcode-cn.com/problems/count-good-triplets-in-an-array/solution/zhuan-hua-wei-jing-dian-wen-ti-ji-ke-by-b03sy/
// - https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof
// - https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self

func lowBit(x int) int { return x & -x }

// BIT -- Binary Indexed Tree, 树状数组.
type BIT struct{ Arr []int64 }

func (b BIT) query(index int) int64 {
	ans := int64(0)
	for ; index > 0; index -= lowBit(index) {
		ans += b.Arr[index]
	}
	return ans
}

func (b BIT) add(index int, val int64) {
	for ; index < len(b.Arr); index += lowBit(index) {
		b.Arr[index] += val
	}
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)
	P := make([]int, n)
	for i, v := range nums1 {
		P[v] = i
	}
	Q := make([]int, n)
	for i, v := range nums2 {
		Q[i] = P[v]
	}
	// 至此转为三元递增组数量 -- 已经离散化.
	left, right := make([]int64, n), make([]int64, n)
	tr := BIT{Arr: make([]int64, n+1)}
	ans := int64(0)
	for i := 0; i < n; i++ {
		left[i] = tr.query(Q[i])
		tr.add(Q[i]+1, 1)
	}
	tr = BIT{Arr: make([]int64, n+1)}
	for i := n - 1; i >= 0; i-- {
		right[i] = int64(n-1-i) - tr.query(Q[i])
		tr.add(Q[i]+1, 1)
	}
	for i := 0; i < n; i++ {
		ans += left[i] * right[i]
	}
	return ans
}
