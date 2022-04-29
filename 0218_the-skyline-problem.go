/*
 * @lc app=leetcode.cn id=the-skyline-problem lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [218] 天际线问题
 *
 * https://leetcode-cn.com/problems/the-skyline-problem/description/
 *
 * algorithms
 * Hard (54.44%)
 * Total Accepted:    37.9K
 * Total Submissions: 69.5K
 * Testcase Example:  '[[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]'
 *
 * 城市的 天际线 是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。给你所有建筑物的位置和高度，请返回 由这些建筑物形成的 天际线 。
 *
 * 每个建筑物的几何信息由数组 buildings 表示，其中三元组 buildings[i] = [lefti, righti, heighti]
 * 表示：
 *
 *
 * lefti 是第 i 座建筑物左边缘的 x 坐标。
 * righti 是第 i 座建筑物右边缘的 x 坐标。
 * heighti 是第 i 座建筑物的高度。
 *
 *
 * 你可以假设所有的建筑都是完美的长方形，在高度为 0 的绝对平坦的表面上。
 *
 * 天际线 应该表示为由 “关键点” 组成的列表，格式 [[x1,y1],[x2,y2],...] ，并按 x 坐标 进行 排序
 * 。关键点是水平线段的左端点。列表中最后一个点是最右侧建筑物的终点，y 坐标始终为 0
 * ，仅用于标记天际线的终点。此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。
 *
 * 注意：输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...]
 * 是不正确的答案；三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
 * 输出：[[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
 * 解释：
 * 图 A 显示输入的所有建筑物的位置和高度，
 * 图 B 显示由这些建筑物形成的天际线。图 B 中的红点表示输出列表中的关键点。
 *
 * 示例 2：
 *
 *
 * 输入：buildings = [[0,2,3],[2,5,3]]
 * 输出：[[0,3],[5,0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= buildings.length <= 10^4
 * 0 <= lefti < righti <= 2^31 - 1
 * 1 <= heighti <= 2^31 - 1
 * buildings 按 lefti 非递减排序
 *
 *
 */
/*
 * func getSkyline(buildings [][]int) [][]int {
 * 	if len(buildings) == 0 {
 * 		return buildings
 * 	}
 * 	if len(buildings) == 1 {
 * 		ans := [][]int{{buildings[0][0], buildings[0][2]}, {buildings[0][1], 0}}
 * 		return ans
 * 	}
 * 	mid := len(buildings) >> 1
 * 	ans := mergeSkyline(buildings[:mid], buildings[mid:])
 * 	return ans
 * }
 *
 * func mergeSkyline(left, right [][]int) [][]int {
 * 	n, m := len(left), len(right)
 * 	if n == 0 {
 * 		return right
 * 	}
 * 	if m == 0 {
 * 		return left
 * 	}
 * 	var ans [][]int
 * 	left = getSkyline(left)
 * 	right = getSkyline(right)
 * 	n, m = len(left), len(right)
 * 	var positions []int
 * 	p, q := 0, 0
 * 	for p < n && q < m {
 * 		if left[p][0] < right[q][0] {
 * 			positions = append(positions, left[p][0])
 * 			p++
 * 		} else if left[p][0] > right[q][0] {
 * 			positions = append(positions, right[q][0])
 * 			q++
 * 		} else {
 * 			positions = append(positions, left[p][0])
 * 			p, q = p+1, q+1
 * 		}
 * 	}
 * 	for ; p < n; p++ {
 * 		positions = append(positions, left[p][0])
 * 	}
 * 	for ; q < m; q++ {
 * 		positions = append(positions, right[q][0])
 * 	}
 * 	p, q = 0, 0
 * 	for _, pos := range positions {
 * 		cur := 0
 * 		for get(left, p+1) <= pos {
 * 			p++
 * 		}
 * 		if left[p][0] <= pos {
 * 			cur = max(cur, left[p][1])
 * 		}
 * 		for get(right, q+1) <= pos {
 * 			q++
 * 		}
 * 		if right[q][0] <= pos {
 * 			cur = max(cur, right[q][1])
 * 		}
 * 		if len(ans) == 0 || ans[len(ans)-1][1] != cur {
 * 			ans = append(ans, []int{pos, cur})
 * 		}
 * 	}
 * 	return ans
 * }
 *
 * func get(arr [][]int, ind int) int {
 * 	if ind >= len(arr) {
 * 		return math.MaxInt64
 * 	}
 * 	return arr[ind][0]
 * }
 *
 * func max(x, y int) int {
 * 	if x > y {
 * 		return x
 * 	}
 * 	return y
 * }
 */

// 官解是优先队列, 我看不懂但我大受震撼...

const (
	BuildingLeft = iota
	BuildingRight
	BuildingHeight
)

const (
	SkylineRight = iota
	SkylineHeight
)

type Heap [][2]int

func (h Heap) Len() int              { return len(h) }
func (h Heap) Less(p, q int) bool    { return h[p][SkylineHeight] > h[q][SkylineHeight] }
func (h Heap) Swap(p, q int)         { h[p], h[q] = h[q], h[p] }
func (h *Heap) Push(v interface{})   { *h = append(*h, v.([2]int)) }
func (h *Heap) Pop() (v interface{}) { l := len(*h); v = (*h)[l-1]; *h = (*h)[:l-1]; return }
func (h *Heap) Front() [2]int        { return (*h)[0] }

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	positions := make([]int, 0, n<<1)
	for _, b := range buildings {
		positions = append(positions, b[BuildingLeft], b[BuildingRight])
	}
	sort.Ints(positions)

	var ans [][]int
	h := &Heap{}
	ind := 0
	for _, p := range positions {
		for ind < n && buildings[ind][BuildingLeft] <= p {
			heap.Push(h, [2]int{buildings[ind][BuildingRight], buildings[ind][BuildingHeight]})
			ind++
		}
		for h.Len() > 0 && h.Front()[SkylineRight] <= p {
			heap.Pop(h)
		}
		maxn := 0
		if h.Len() > 0 {
			maxn = h.Front()[SkylineHeight]
		}
		if m := len(ans); m == 0 || ans[m-1][SkylineHeight] != maxn {
			ans = append(ans, []int{p, maxn})
		}
	}
	return ans
}
