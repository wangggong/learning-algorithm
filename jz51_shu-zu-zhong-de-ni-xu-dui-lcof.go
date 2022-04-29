/*
 * Hard
 *
 * 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
 *
 *
 *
 * 示例 1:
 *
 * 输入: [7,5,6,4]
 * 输出: 5
 *
 *
 * 限制：
 *
 * 0 <= 数组长度 <= 50000
 *
 * 通过次数120,254
 * 提交次数246,402
 *
 */

/*
 * import "sort"
 *
 * func lowBit(x int) int { return x & (-x) }
 *
 * type BIT []int
 *
 * func (b BIT) query(x int) int {
 * 	ans := 0
 * 	for ; x > 0; x -= lowBit(x) {
 * 		ans += b[x]
 * 	}
 * 	return ans
 * }
 *
 * func (b *BIT) add(x, v int) {
 * 	n := len(*b) - 1
 * 	for ; x <= n; x += lowBit(x) {
 * 		(*b)[x] += v
 * 	}
 * 	return
 * }
 *
 * func NewBIT(n int) *BIT {
 * 	arr := make([]int, n+1)
 * 	b := BIT(arr)
 * 	return &b
 * }
 *
 * // 树状数组, 有意思的解法
 * func reversePairs(nums []int) int {
 * 	n := len(nums)
 * 	m := make(map[int]int)
 * 	arr := make([]int, n)
 * 	copy(arr, nums)
 * 	sort.Ints(arr)
 * 	for i, a := range arr {
 * 		m[a] = i + 1
 * 	}
 * 	b := NewBIT(n)
 * 	ans := 0
 * 	for _, num := range nums {
 * 		k := m[num]
 * 		// 正着遍历的话需要避免把取等的情况算进去, 所以别写成这样:
 * 		//
 * 		// ans += b.query(n) - b.query(k-1)  // or ans += b.between(k, n)
 * 		//
 * 		// 上面算的都是 `[k,n]` 闭区间的前缀和, 也就是 "在当前数字前 >= k 的数字的总个数".
 * 		ans += b.query(n) - b.query(k)
 * 		b.add(k, 1)
 * 	}
 * 	return ans
 * }
 */

func reversePairs(nums []int) int {
	n := len(nums)
	return merge(nums, 0, n)
}

// 剑指 offer 原题解, 基于 mergeSort 实现.
func merge(arr []int, l, r int) int {
	if l+1 >= r {
		return 0
	}
	mid := (l + r) >> 1
	left := merge(arr, l, mid)
	right := merge(arr, mid, r)
	cpy := make([]int, r-l)
	p, q := mid-1, r-1
	k := r - l - 1
	cnt := 0
	for p >= l && q >= mid {
		// 这里同样要注意避免对取等计数, 所以是 `<=` 而不是 `<`.
		if arr[p] <= arr[q] {
			cpy[k] = arr[q]
			k--
			q--
		} else {
			cpy[k] = arr[p]
			cnt += q - mid + 1
			k--
			p--
		}
	}
	for p >= l {
		cpy[k] = arr[p]
		cnt += q - mid + 1
		k--
		p--
	}
	for q >= mid {
		cpy[k] = arr[q]
		k--
		q--
	}
	copy(arr[l:r], cpy)
	return left + right + cnt
}
