/*
 * @lc app=leetcode.cn id=find-k-pairs-with-smallest-sums lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [373] 查找和最小的 K 对数字
 *
 * https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums/description/
 *
 * algorithms
 * Medium (42.45%)
 * Total Accepted:    44.4K
 * Total Submissions: 104.7K
 * Testcase Example:  '[1,7,11]\n[2,4,6]\n3'
 *
 * 给定两个以 升序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
 *
 * 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
 *
 * 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
 * 输出: [1,2],[1,4],[1,6]
 * 解释: 返回序列中的前 3 对数：
 * ⁠    [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
 *
 *
 * 示例 2:
 *
 *
 * 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
 * 输出: [1,1],[1,1]
 * 解释: 返回序列中的前 2 对数：
 * [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
 *
 *
 * 示例 3:
 *
 *
 * 输入: nums1 = [1,2], nums2 = [3], k = 3
 * 输出: [1,3],[2,3]
 * 解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
 *
 *
 *
 *
 * 提示:
 *
 *
 * 1 <= nums1.length, nums2.length <= 10^5
 * -10^9 <= nums1[i], nums2[i] <= 10^9
 * nums1 和 nums2 均为升序排列
 * 1 <= k <= 1000
 *
 *
 */
import (
	"container/heap"
)

type P struct {
	x, y, sum int
}

type hp struct {
	Ps []P
}

func NewHp() *hp                 { return &hp{} }
func (h hp) Len() int            { return len(h.Ps) }
func (h hp) Less(p, q int) bool  { return h.Ps[p].sum < h.Ps[q].sum }
func (h hp) Swap(p, q int)       { h.Ps[p], h.Ps[q] = h.Ps[q], h.Ps[p] }
func (h *hp) Push(p interface{}) { h.Ps = append(h.Ps, p.(P)) }
func (h *hp) Pop() interface{}   { l := len(h.Ps); v := h.Ps[l-1]; h.Ps = h.Ps[:l-1]; return v }

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n1, n2 := len(nums1), len(nums2)
	ans := make([][]int, 0, k)
	h := NewHp()
	heap.Push(h, P{0, 0, nums1[0] + nums2[0]})

	m := make(map[P]struct{})
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(h).(P)
		// fmt.Printf("%v %v\n", p, h)
		x, y := p.x, p.y
		ans = append(ans, []int{nums1[x], nums2[y]})
		if x+1 < n1 {
			np := P{x + 1, y, nums1[x+1] + nums2[y]}
			// fmt.Printf("%v\n", np)
			if _, ok := m[np]; !ok {
				heap.Push(h, np)
				m[np] = struct{}{}
			}
		}
		if y+1 < n2 {
			np := P{x, y + 1, nums1[x] + nums2[y+1]}
			// fmt.Printf("%v\n", np)
			if _, ok := m[np]; !ok {
				heap.Push(h, np)
				m[np] = struct{}{}
			}
		}
	}
	return ans
}
