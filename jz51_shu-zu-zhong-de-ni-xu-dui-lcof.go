import "sort"

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

func lowBit(x int) int { return x & (-x) }

type BIT struct {
	Arr []int
}

func (b BIT) query(index int) int {
	ans := 0
	for ; index > 0; index -= lowBit(index) {
		ans += b.Arr[index]
	}
	return ans
}

func (b BIT) add(index int, val int) {
	for ind := index; ind < len(b.Arr); ind += lowBit(ind) {
		b.Arr[ind] += val
	}
}

func reversePairs(nums []int) int {
	n := len(nums)
	m := make(map[int]int)
	var vals []int
	for _, v := range nums {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = 0
		vals = append(vals, v)
	}
	sort.Ints(vals)
	for i, v := range vals {
		m[v] = i
	}
	for i := 0; i < n; i++ {
		nums[i] = m[nums[i]] + 1
	}
	tr := BIT{Arr: make([]int, n)}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans += tr.query(nums[i] - 1)
		tr.add(nums[i], 1)
	}
	return ans
}
