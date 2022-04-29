/*
 * Medium
 *
 * 给你一个二维整数数组 rectangles ，其中 rectangles[i] = [li, hi] 表示第 i 个矩形长为 li 高为 hi 。给你一个二维整数数组 points ，其中 points[j] = [xj, yj] 是坐标为 (xj, yj) 的一个点。
 *
 * 第 i 个矩形的 左下角 在 (0, 0) 处，右上角 在 (li, hi) 。
 *
 * 请你返回一个整数数组 count ，长度为 points.length，其中 count[j]是 包含 第 j 个点的矩形数目。
 *
 * 如果 0 <= xj <= li 且 0 <= yj <= hi ，那么我们说第 i 个矩形包含第 j 个点。如果一个点刚好在矩形的 边上 ，这个点也被视为被矩形包含。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：rectangles = [[1,2],[2,3],[2,5]], points = [[2,1],[1,4]]
 * 输出：[2,1]
 * 解释：
 * 第一个矩形不包含任何点。
 * 第二个矩形只包含一个点 (2, 1) 。
 * 第三个矩形包含点 (2, 1) 和 (1, 4) 。
 * 包含点 (2, 1) 的矩形数目为 2 。
 * 包含点 (1, 4) 的矩形数目为 1 。
 * 所以，我们返回 [2, 1] 。
 * 示例 2：
 *
 *
 *
 * 输入：rectangles = [[1,1],[2,2],[3,3]], points = [[1,3],[1,1]]
 * 输出：[1,3]
 * 解释：
 * 第一个矩形只包含点 (1, 1) 。
 * 第二个矩形只包含点 (1, 1) 。
 * 第三个矩形包含点 (1, 3) 和 (1, 1) 。
 * 包含点 (1, 3) 的矩形数目为 1 。
 * 包含点 (1, 1) 的矩形数目为 3 。
 * 所以，我们返回 [1, 3] 。
 *
 *
 * 提示：
 *
 * 1 <= rectangles.length, points.length <= 5 * 104
 * rectangles[i].length == points[j].length == 2
 * 1 <= li, xj <= 109
 * 1 <= hi, yj <= 100
 * 所有 rectangles 互不相同 。
 * 所有 points 互不相同 。
 * 通过次数1,908
 * 提交次数8,827
 */

/*
 * const maxn int = 100
 *
 * var visit [maxn + 10][]int
 *
 * func countRectangles(rectangles [][]int, points [][]int) []int {
 * 	for i := range visit {
 * 		visit[i] = nil
 * 	}
 * 	for _, r := range rectangles {
 * 		x, y := r[0], r[1]
 * 		for j := 0; j <= y; j++ {
 * 			visit[j] = append(visit[j], x)
 * 		}
 * 	}
 * 	for i := range visit {
 * 		sort.Ints(visit[i])
 * 	}
 * 	var ans []int
 * 	for _, p := range points {
 * 		x, y := p[0], p[1]
 * 		p, q := 0, len(visit[y])
 * 		for p < q {
 * 			mid := (p + q) >> 1
 * 			if visit[y][mid] >= x {
 * 				q = mid
 * 			} else {
 * 				p = mid + 1
 * 			}
 * 		}
 * 		ans = append(ans, len(visit[y])-p)
 * 	}
 * 	return ans
 * }
 */

// 从高到低考虑, 点位是逐渐增加的. 每次只需要添加点位+排序即可.
//
// 注: 如果这题纵坐标的范围也是 `10^9`, 我们还可以用名次树来做出此题 (如 `Python` 的 `SortedList`).
func countRectangles(rectangles [][]int, points [][]int) []int {
	sort.Slice(rectangles, func(p, q int) bool { return rectangles[p][1] > rectangles[q][1] })
	n, m := len(rectangles), len(points)
	for i := range points {
		points[i] = append(points[i], i)
	}
	sort.Slice(points, func(p, q int) bool { return points[p][1] > points[q][1] })
	ans := make([]int, m)
	var xs []int
	ind := 0
	for _, p := range points {
		size := len(xs)
		for ; ind < n && rectangles[ind][1] >= p[1]; ind++ {
			xs = append(xs, rectangles[ind][0])
		}
		if len(xs) > size {
			sort.Ints(xs)
			size = len(xs)
		}
		ans[p[2]] = ind - sort.SearchInts(xs, p[0])
	}
	return ans
}
