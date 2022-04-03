/*
 * Medium
 *
 * 对于一棵深度小于 5 的树，可以用一组三位十进制整数来表示。对于每个整数：
 *
 * 百位上的数字表示这个节点的深度 d，1 <= d <= 4。
 * 十位上的数字表示这个节点在当前层所在的位置 P， 1 <= p <= 8。位置编号与一棵满二叉树的位置编号相同。
 * 个位上的数字表示这个节点的权值 v，0 <= v <= 9。
 * 给定一个包含三位整数的 升序 数组 nums ，表示一棵深度小于 5 的二叉树，请你返回 从根到所有叶子结点的路径之和 。
 *
 * 保证 给定的数组表示一个有效的连接二叉树。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入: nums = [113, 215, 221]
 * 输出: 12
 * 解释: 列表所表示的树如上所示。
 * 路径和 = (3 + 5) + (3 + 1) = 12.
 * 示例 2：
 *
 *
 *
 * 输入: nums = [113, 221]
 * 输出: 4
 * 解释: 列表所表示的树如上所示。
 * 路径和 = (3 + 1) = 4.
 *
 *
 * 提示:
 *
 * 1 <= nums.length <= 15
 * 110 <= nums[i] <= 489
 * nums 表示深度小于 5 的有效二叉树
 * 通过次数2,487
 * 提交次数4,028
 */

const MAXN int = 1 << 4

var val [MAXN]int

func pathSum(nums []int) int {
	for i := range val {
		val[i] = -1
	}
	for _, n := range nums {
		val[(1<<((n/100)-1))+((n/10)%10)-1] = n % 10
	}
	ans := 0
	for i := 0; i < MAXN; i++ {
		if valid(i<<1) || valid((i<<1)+1) || val[i] < 0 {
			continue
		}
		c := i
		for c > 0 {
			ans += val[c]
			c >>= 1
		}
	}
	return ans
}

func valid(k int) bool {
	if k < 0 || k >= MAXN {
		return false
	}
	return val[k] >= 0
}

