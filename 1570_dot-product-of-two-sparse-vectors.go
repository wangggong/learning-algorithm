/*
 * Medium
 *
 * 给定两个稀疏向量，计算它们的点积（数量积）。
 *
 * 实现类 SparseVector：
 *
 * SparseVector(nums) 以向量 nums 初始化对象。
 * dotProduct(vec) 计算此向量与 vec 的点积。
 * 稀疏向量 是指绝大多数分量为 0 的向量。你需要 高效 地存储这个向量，并计算两个稀疏向量的点积。
 *
 * 进阶：当其中只有一个向量是稀疏向量时，你该如何解决此问题？
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [1,0,0,2,3], nums2 = [0,3,0,4,0]
 * 输出：8
 * 解释：v1 = SparseVector(nums1) , v2 = SparseVector(nums2)
 * v1.dotProduct(v2) = 1*0 + 0*3 + 0*0 + 2*4 + 3*0 = 8
 * 示例 2：
 *
 * 输入：nums1 = [0,1,0,0,0], nums2 = [0,0,0,0,2]
 * 输出：0
 * 解释：v1 = SparseVector(nums1) , v2 = SparseVector(nums2)
 * v1.dotProduct(v2) = 0*0 + 1*0 + 0*0 + 0*0 + 0*2 = 0
 * 示例 3：
 *
 * 输入：nums1 = [0,1,0,0,2,0,0], nums2 = [1,0,0,0,3,0,4]
 * 输出：6
 *
 *
 * 提示：
 *
 * n == nums1.length == nums2.length
 * 1 <= n <= 10^5
 * 0 <= nums1[i], nums2[i] <= 100
 *
 */
type SparseVector struct {
	index2Value map[int]int
}

func Constructor(nums []int) SparseVector {
	sv := SparseVector{index2Value: make(map[int]int)}
	for i, n := range nums {
		if n == 0 {
			continue
		}
		sv.index2Value[i] = n
	}
	return sv
}

// Return the dotProduct of two sparse vectors
func (sv *SparseVector) dotProduct(vec SparseVector) int {
	ans := 0
	if vec.index2Value == nil {
		return ans
	}
	for i, n := range sv.index2Value {
		ans += n * vec.index2Value[i]
	}
	return ans
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */
