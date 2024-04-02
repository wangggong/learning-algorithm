/*
 * Medium
 *
 * 有个内含单词的超大文本文件，给定任意两个不同的单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?
 *
 * 示例：
 *
 * 输入：words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
 * 输出：1
 * 提示：
 *
 * words.length <= 100000
 * 通过次数26,622
 * 提交次数37,181
 */

import "math"

func findClosest(words []string, word1 string, word2 string) int {
	index := make(map[string][]int)
	for i, w := range words {
		index[w] = append(index[w], i)
	}
	ans := math.MaxInt32
	word2Indexes := index[word2]
	for _, i := range index[word1] {
		p, q := 0, len(word2Indexes)-1
		if i < word2Indexes[p] {
			ans = min(ans, word2Indexes[p]-i)
			continue
		}
		if i > word2Indexes[q] {
			ans = min(ans, i-word2Indexes[q])
			continue
		}
		for p < q {
			mid := (p + q) >> 1
			if word2Indexes[mid] >= i {
				q = mid
			} else {
				p = mid + 1
			}
		}
		ans = min(ans, word2Indexes[p]-i)
		ans = min(ans, i-word2Indexes[p-1])
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
 * func findClosest(words []string, word1 string, word2 string) int {
 * 	index := make(map[string][]int)
 * 	for i, w := range words {
 * 		index[w] = append(index[w], i)
 * 	}
 * 	ans := math.MaxInt32
 * 	for _, i := range index[word1] {
 * 		for _, j := range index[word2] {
 * 			ans = min(ans, abs(i-j))
 * 		}
 * 	}
 * 	return ans
 * }
 * 
 * func abs(x int) int {
 * 	if x >= 0 {
 * 		return x
 * 	}
 * 	return -x
 * }
 */ 
