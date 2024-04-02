/*
 * @lc app=leetcode.cn id=contain-virus lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [749] 隔离病毒
 *
 * https://leetcode.cn/problems/contain-virus/description/
 *
 * algorithms
 * Hard (50.13%)
 * Total Accepted:    2.6K
 * Total Submissions: 4.9K
 * Testcase Example:  '[[0,1,0,0,0,0,0,1],[0,1,0,0,0,0,0,1],[0,0,0,0,0,0,0,1],[0,0,0,0,0,0,0,0]]'
 *
 * 病毒扩散得很快，现在你的任务是尽可能地通过安装防火墙来隔离病毒。
 *
 * 假设世界由 m x n 的二维矩阵 isInfected 组成， isInfected[i][j] == 0 表示该区域未感染病毒，而
 * isInfected[i][j] == 1 表示该区域已感染病毒。可以在任意 2 个相邻单元之间的共享边界上安装一个防火墙（并且只有一个防火墙）。
 *
 *
 * 每天晚上，病毒会从被感染区域向相邻未感染区域扩散，除非被防火墙隔离。现由于资源有限，每天你只能安装一系列防火墙来隔离其中一个被病毒感染的区域（一个区域或连续的一片区域），且该感染区域对未感染区域的威胁最大且
 * 保证唯一 。
 *
 * 你需要努力使得最后有部分区域不被病毒感染，如果可以成功，那么返回需要使用的防火墙个数;
 * 如果无法实现，则返回在世界被病毒全部感染时已安装的防火墙个数。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入: isInfected =
 * [[0,1,0,0,0,0,0,1],[0,1,0,0,0,0,0,1],[0,0,0,0,0,0,0,1],[0,0,0,0,0,0,0,0]]
 * 输出: 10
 * 解释:一共有两块被病毒感染的区域。
 * 在第一天，添加 5 墙隔离病毒区域的左侧。病毒传播后的状态是:
 *
 * 第二天，在右侧添加 5 个墙来隔离病毒区域。此时病毒已经被完全控制住了。
 *
 *
 *
 * 示例 2：
 *
 *
 *
 *
 * 输入: isInfected = [[1,1,1],[1,0,1],[1,1,1]]
 * 输出: 4
 * 解释: 虽然只保存了一个小区域，但却有四面墙。
 * 注意，防火墙只建立在两个不同区域的共享边界上。
 *
 *
 * 示例 3:
 *
 *
 * 输入: isInfected =
 * [[1,1,1,0,0,0,0,0,0],[1,0,1,0,1,1,1,1,1],[1,1,1,0,0,0,0,0,0]]
 * 输出: 13
 * 解释: 在隔离右边感染区域后，隔离左边病毒区域只需要 2 个防火墙。
 *
 *
 *
 *
 * 提示:
 *
 *
 * m == isInfected.length
 * n == isInfected[i].length
 * 1 <= m, n <= 50
 * isInfected[i][j] is either 0 or 1
 * 在整个描述的过程中，总有一个相邻的病毒区域，它将在下一轮 严格地感染更多未受污染的方块
 *
 *
 *
 *
 */

// 超烦的一道题, 难受
var directions = []int{0, 1, 0, -1, 0}

func containVirus(G [][]int) int {
	n, m := len(G), len(G[0])
	ans := 0
	for true {
		var neighbors []map[[2]int]struct{}
		var firewalls []int
		idx := 0
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if G[i][j] != 1 {
					continue
				}
				idx--
				neighbor := make(map[[2]int]struct{})
				firewall := 0
				var Q [][2]int
				Q = append(Q, [2]int{i, j})
				for len(Q) > 0 {
					q := Q[0]
					Q = Q[1:]
					x, y := q[0], q[1]
					switch G[x][y] {
					case 1:
						G[x][y] = idx
						for i := 0; i < 4; i++ {
							nx, ny := x+directions[i], y+directions[i+1]
							if 0 <= nx && nx < n && 0 <= ny && ny < m {
								Q = append(Q, [2]int{nx, ny})
							}
						}
					case 0:
						neighbor[q] = struct{}{}
						firewall++
					default:
					}
				}
				neighbors = append(neighbors, neighbor)
				firewalls = append(firewalls, firewall)
			}
		}
		k := len(neighbors)
		if k == 0 {
			break
		}
		idx = 0
		for i := 0; i < k; i++ {
			if len(neighbors[i]) > len(neighbors[idx]) {
				idx = i
			}
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if G[i][j] >= 0 {
					continue
				}
				if G[i][j] == -1-idx {
					G[i][j] = 2
				} else {
					G[i][j] = 1
				}
			}
		}
		for i := 0; i < k; i++ {
			if i == idx {
				ans += firewalls[i]
				continue
			}
			for n := range neighbors[i] {
				x, y := n[0], n[1]
				G[x][y] = 1
			}
		}
		if k == 1 {
			break
		}
	}
	return ans
}
