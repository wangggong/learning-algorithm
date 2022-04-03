/*
 * @lc app=leetcode.cn id=zui-xiao-tiao-yue-ci-shu lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [LCP 09] 最小跳跃次数
 *
 * https://leetcode-cn.com/problems/zui-xiao-tiao-yue-ci-shu/description/
 *
 * algorithms
 * Hard (28.45%)
 * Total Accepted:    7.2K
 * Total Submissions: 25.4K
 * Testcase Example:  '[2,5,1,1,1,1]'
 *
 * 为了给刷题的同学一些奖励，力扣团队引入了一个弹簧游戏机。游戏机由 N 个特殊弹簧排成一排，编号为 0 到 N-1。初始有一个小球在编号 0
 * 的弹簧处。若小球在编号为 i 的弹簧处，通过按动弹簧，可以选择把小球向右弹射 jump[i]
 * 的距离，或者向左弹射到任意左侧弹簧的位置。也就是说，在编号为 i 弹簧处按动弹簧，小球可以弹向 0 到 i-1 中任意弹簧或者 i+jump[i]
 * 的弹簧（若 i+jump[i]>=N ，则表示小球弹出了机器）。小球位于编号 0 处的弹簧时不能再向左弹。
 *
 * 为了获得奖励，你需要将小球弹出机器。请求出最少需要按动多少次弹簧，可以将小球从编号 0 弹簧弹出整个机器，即向右越过编号 N-1 的弹簧。
 *
 * 示例 1：
 *
 *
 * 输入：jump = [2, 5, 1, 1, 1, 1]
 *
 * 输出：3
 *
 * 解释：小 Z 最少需要按动 3 次弹簧，小球依次到达的顺序为 0 -> 2 -> 1 -> 6，最终小球弹出了机器。
 *
 *
 * 限制：
 *
 *
 * 1 <= jump.length <= 10^6
 * 1 <= jump[i] <= 10000
 *
 *
 */

func minJump(jump []int) int {
	n := len(jump)
	ans := n
	preIndex := 0
	var Q [][2]int
	visit := make(map[int]struct{})
	Q = append(Q, [2]int{0, 0})
	for len(Q) > 0 {
		q := Q[0]
		visit[q[0]] = struct{}{}
		Q = Q[1:]
		if q[1] > ans {
			continue
		}
		if q[0] >= n {
			if ans > q[1] {
				ans = q[1]
			}
			continue
		}
		newPos := q[0] + jump[q[0]]
		if _, ok := visit[newPos]; !ok {
			Q = append(Q, [2]int{newPos, q[1] + 1})
			visit[newPos] = struct{}{}
		}
		/*
			// 这么写会超时的, `O(n^2)`. 所以维护一个已经走过的 index, 让每个 index 只走一次.
			for i := 0; i < q[0]; i++ {
				newPos := i
				if _, ok := visit[newPos]; !ok {
					Q = append(Q, [2]int{newPos, q[1] + 1})
					visit[newPos] = struct{}{}
				}
			}
		*/
		for preIndex < q[0] {
			if _, ok := visit[preIndex]; !ok {
				Q = append(Q, [2]int{preIndex, q[1] + 1})
				visit[preIndex] = struct{}{}
			}
			preIndex++
		}
	}
	return ans
}
