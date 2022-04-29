/*
 * @lc app=leetcode.cn id=hSRGyL lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [LCP 26] 导航装置
 *
 * https://leetcode-cn.com/problems/hSRGyL/description/
 *
 * algorithms
 * Hard (35.52%)
 * Total Accepted:    671
 * Total Submissions: 1.9K
 * Testcase Example:  '[1,2,null,3,4]'
 *
 * 小扣参加的秋日市集景区共有 $N$ 个景点，景点编号为 $1$~$N$。景点内设有 $N-1$ 条双向道路，使所有景点形成了一个二叉树结构，根结点记为
 * `root`，景点编号即为节点值。
 *
 * 由于秋日市集景区的结构特殊，游客很容易迷路，主办方决定在景区的若干个景点设置导航装置，按照所在景点编号升序排列后定义装置编号为 1 ~
 * M。导航装置向游客发送数据，数据内容为列表 `[游客与装置 1 的相对距离,游客与装置 2 的相对距离,...,游客与装置 M
 * 的相对距离]`。由于游客根据导航装置发送的信息来确认位置，因此主办方需保证游客在每个景点接收的数据信息皆不相同。请返回主办方最少需要设置多少个导航装置。
 *
 * **示例 1：**
 * >输入：`root = [1,2,null,3,4]`
 * >
 * >输出：`2`
 * >
 * >解释：在景点 1、3 或景点 1、4 或景点 3、4 设置导航装置。
 * >
 *
 * >![image.png](https://pic.leetcode-cn.com/1597996812-tqrgwu-image.png){:height="250px"}
 *
 *
 *
 * **示例 2：**
 * >输入：`root = [1,2,3,4]`
 * >
 * >输出：`1`
 * >
 * >解释：在景点 3、4 设置导航装置皆可。
 * >
 *
 * >![image.png](https://pic.leetcode-cn.com/1597996826-EUQRyz-image.png){:height="200px"}
 *
 *
 *
 * **提示：**
 * - `2 <= N <= 50000`
 * - 二叉树的非空节点值为 `1~N` 的一个排列。
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 树形贪心, 参考 `监控二叉树`. 诶看着像看着像就是想不到, 气不气? 所幸实现的时候还有好几个坑, 活该做不出来.
//
// 参考: https://leetcode-cn.com/problems/hSRGyL/solution/tan-xin-suan-fa-zheng-ming-by-newhar/
const (
	// noCamera 指当前位置不能被导航 cover 到.
	noCamera = iota
	// camera 指当前位置能被导航 cover 到.
	camera
)

var ans int
var status int

func navigation(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans, status = 0, noCamera
	left, right := dfs(root.Left), dfs(root.Right)
	// 如果左右子树都能被导航 cover, 直接返回.
	if left == camera && right == camera {
		return ans
	}
	// 考虑第一个三岔口, 如果这里没导航, 得安一个.
	if status == noCamera {
		ans++
	}
	return ans
}

func dfs(root *TreeNode) int {
	// 叶子结点默认 **先** 不安导航, 也不能被 cover 到. 如果出现了 "三岔口", 可能会在叶子结点安一个. 但 **安哪都行**.
	if root == nil {
		return noCamera
	}
	left, right := dfs(root.Left), dfs(root.Right)
	// "三岔口" 的特殊判定:
	if root.Left != nil && root.Right != nil {
		// 如果左右子都没安导航, 随便选一个安上. 总数加 1.
		if left == noCamera && right == noCamera {
			ans++
		}
		// 后序遍历, 最终 `status` 指的是最后一次 "三岔口" 的导航状态. 如果左右子都能被导航 cover 到, 其实也就不用安了. 否则必安.
		status = noCamera
		if left == camera && right == camera {
			status = camera
		}
		// "三岔口" 均视为能被导航 cover 到 (结论是一个 "三岔口" 至少得有 2 个导航)
		return camera
	}
	if left == camera || right == camera {
		return camera
	}
	return noCamera
}

// 二刷也是, 一点思路没有.
