/*
 * @lc app=leetcode.cn id=count-nodes-with-the-highest-score lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [2049] 统计最高分的节点数目
 *
 * https://leetcode-cn.com/problems/count-nodes-with-the-highest-score/description/
 *
 * algorithms
 * Medium (40.01%)
 * Total Accepted:    9.2K
 * Total Submissions: 18.4K
 * Testcase Example:  '[-1,2,0,2,0]'
 *
 * 给你一棵根节点为 0 的 二叉树 ，它总共有 n 个节点，节点编号为 0 到 n - 1 。同时给你一个下标从 0 开始的整数数组 parents
 * 表示这棵树，其中 parents[i] 是节点 i 的父节点。由于节点 0 是根，所以 parents[0] == -1 。
 *
 * 一个子树的 大小 为这个子树内节点的数目。每个节点都有一个与之关联的 分数 。求出某个节点分数的方法是，将这个节点和与它相连的边全部 删除
 * ，剩余部分是若干个 非空 子树，这个节点的 分数 为所有这些子树 大小的乘积 。
 *
 * 请你返回有 最高得分 节点的 数目 。
 *
 *
 *
 * 示例 1:
 *
 *
 *
 * 输入：parents = [-1,2,0,2,0]
 * 输出：3
 * 解释：
 * - 节点 0 的分数为：3 * 1 = 3
 * - 节点 1 的分数为：4 = 4
 * - 节点 2 的分数为：1 * 1 * 2 = 2
 * - 节点 3 的分数为：4 = 4
 * - 节点 4 的分数为：4 = 4
 * 最高得分为 4 ，有三个节点得分为 4 （分别是节点 1，3 和 4 ）。
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：parents = [-1,2,0]
 * 输出：2
 * 解释：
 * - 节点 0 的分数为：2 = 2
 * - 节点 1 的分数为：2 = 2
 * - 节点 2 的分数为：1 * 1 = 1
 * 最高分数为 2 ，有两个节点分数为 2 （分别为节点 0 和 1 ）。
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == parents.length
 * 2 <= n <= 10^5
 * parents[0] == -1
 * 对于 i != 0 ，有 0 <= parents[i] <= n - 1
 * parents 表示一棵二叉树。
 *
 *
 */

// const MAXN int = 5
const MAXN int = 1e5

var left, right, father, size, mul [MAXN + 5]int
var n int

func countHighestScoreNodes(parents []int) int {
	n = len(parents)
	for i := 0; i < MAXN+5; i++ {
		left[i] = -1
		right[i] = -1
		size[i] = 0
		mul[i] = 0
	}
	for i, p := range parents {
		father[i] = p
		if p < 0 {
			continue
		}
		if left[p] == -1 {
			left[p] = i
		} else {
			right[p] = i
		}
	}
	dfs(0)
	maxVal := 0
	for i := 0; i < n; i++ {
		leftVal, rightVal, fatherVal := 1, 1, 1
		if left[i] >= 0 {
			leftVal = size[left[i]]
		}
		if right[i] >= 0 {
			rightVal = size[right[i]]
		}
		if i != 0 {
			fatherVal = size[0] - size[i]
		}
		mul[i] = leftVal * rightVal * fatherVal
		if mul[i] > maxVal {
			maxVal = mul[i]
		}
	}
	// fmt.Println(left, right, father, size, mul)
	ans := 0
	for i := 0; i < n; i++ {
		if mul[i] == maxVal {
			ans++
		}
	}
	return ans
}

func dfs(k int) int {
	if left[k] < 0 && right[k] < 0 {
		size[k] = 1
		return size[k]
	}
	size[k] = 1
	if left[k] >= 0 {
		size[k] += dfs(left[k])
	}
	if right[k] >= 0 {
		size[k] += dfs(right[k])
	}
	return size[k]
}

