/*
 * @lc app=leetcode.cn id=contains-duplicate-iii lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [220] 存在重复元素 III
 *
 * https://leetcode-cn.com/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Medium (28.73%)
 * Total Accepted:    70.8K
 * Total Submissions: 246K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * 给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j])
 * ，同时又满足 abs(i - j)  。
 *
 * 如果存在则返回 true，不存在返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3,1], k = 3, t = 0
 * 输出：true
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,0,1,1], k = 1, t = 2
 * 输出：true
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1,5,9,1,5,9], k = 2, t = 3
 * 输出：false
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 2 * 1e4
 * -2^31 <= nums[i] <= 2^31 - 1
 * 0 <= k <= 1e4
 * 0 <= t <= 2^31 - 1
 *
 *
 */
/*
 * import "sort"
 *
 * type IndVal struct {
 * 	index int
 * 	value int
 * }
 *
 * type IndValList []IndVal
 *
 * func (s IndValList) Len() int           { return len(s) }
 * func (s IndValList) Less(p, q int) bool { return s[p].value < s[q].value }
 * func (s IndValList) Swap(p, q int)      { s[p], s[q] = s[q], s[p] }
 *
 * func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
 * 	n := len(nums)
 * 	numIdxs := IndValList{}
 * 	for i, n := range nums {
 * 		numIdxs = append(numIdxs, IndVal{index: i, value: n})
 * 	}
 * 	sort.Sort(numIdxs)
 * 	for i, v := range numIdxs {
 *
 * 		for j := i + 1; j < n; j++ {
 * 			if numIdxs[j].value-v.value > t {
 * 				break
 * 			}
 * 			if d := numIdxs[j].index - v.index; -k <= d && d <= k {
 * 				return true
 * 			}
 * 		}
 * 	}
 * 	return false
 * }
 */

/*
 *
 * // 有序集合 / 树堆模版
 * //
 * // 参考: https://oi-wiki.org/ds/treap/#treap
 *
 * import "math/rand"
 *
 * type node struct {
 * 	ch       [2]*node
 * 	priority int
 * 	val      int
 * }
 *
 * func (n *node) cmp(b int) int {
 * 	switch {
 * 	case b < n.val:
 * 		return 0
 * 	case b > n.val:
 * 		return 1
 * 	default:
 * 		return -1
 * 	}
 * }
 *
 * func (n *node) rotate(d int) *node {
 * 	v := n.ch[d^1]
 * 	n.ch[d^1] = v.ch[d]
 * 	v.ch[d] = n
 * 	return v
 * }
 *
 * type treap struct {
 * 	root *node
 * }
 *
 * func (t *treap) _put(n *node, v int) *node {
 * 	if n == nil {
 * 		return &node{val: v, priority: rand.Int()}
 * 	}
 * 	d := n.cmp(v)
 * 	n.ch[d] = t._put(n.ch[d], v)
 * 	if n.ch[d].priority > n.priority {
 * 		n = n.rotate(d ^ 1)
 * 	}
 * 	return n
 * }
 *
 * func (t *treap) _delete(n *node, v int) *node {
 * 	if d := n.cmp(v); d >= 0 {
 * 		n.ch[d] = t._delete(n.ch[d], v)
 * 		return n
 * 	}
 * 	if n.ch[0] == nil {
 * 		return n.ch[1]
 * 	}
 * 	if n.ch[1] == nil {
 * 		return n.ch[0]
 * 	}
 * 	d := 0
 * 	if n.ch[0].priority > n.ch[1].priority {
 * 		d = 1
 * 	}
 * 	n = n.rotate(d)
 * 	n.ch[d] = t._delete(n.ch[d], v)
 * 	return n
 * }
 *
 * func (t *treap) put(v int)    { t.root = t._put(t.root, v) }
 * func (t *treap) delete(v int) { t.root = t._delete(t.root, v) }
 *
 * func (t *treap) lowerBound(v int) (lb *node) {
 * 	for n := t.root; n != nil; {
 * 		switch c := n.cmp(v); {
 * 		case c == 0:
 * 			lb = n
 * 			n = n.ch[0]
 * 		case c > 0:
 * 			n = n.ch[1]
 * 		default:
 * 			return n
 * 		}
 * 	}
 * 	return
 * }
 *
 * func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
 * 	set := &treap{}
 * 	for i, v := range nums {
 * 		if lb := set.lowerBound(v - t); lb != nil && lb.val <= v+t {
 * 			return true
 * 		}
 * 		set.put(v)
 * 		if i >= k {
 * 			set.delete(nums[i-k])
 * 		}
 * 	}
 * 	return false
 * }
 */

// 桶排序 -- 模版

func getID(n int, w int) int {
	n += 1 << 31
	return n / w
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	m := make(map[int]int)
	w := t + 1
	for i, n := range nums {
		id := getID(n, w)
		if _, ok := m[id]; ok {
			return true
		}
		if v, ok := m[id-1]; ok && abs(v-n) <= t {
			return true
		}
		if v, ok := m[id+1]; ok && abs(v-n) <= t {
			return true
		}
		m[id] = n
		if i >= k {
			delete(m, getID(nums[i-k], w))
		}
	}
	return false
}
