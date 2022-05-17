/*
 * @lc app=leetcode.cn id=top-k-frequent-words lang=golang
 *
 * NOTE (beewangruichao): I cannot use fid, thus I use slug instead.
 *
 * [692] 前K个高频单词
 *
 * https://leetcode.cn/problems/top-k-frequent-words/description/
 *
 * algorithms
 * Medium (56.94%)
 * Total Accepted:    81K
 * Total Submissions: 142.3K
 * Testcase Example:  '["i", "love", "leetcode", "i", "love", "coding"]\n2'
 *
 * 给定一个单词列表 words 和一个整数 k ，返回前 k 个出现次数最多的单词。
 *
 * 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率， 按字典顺序 排序。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入: words = ["i", "love", "leetcode", "i", "love", "coding"], k = 2
 * 输出: ["i", "love"]
 * 解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
 * ⁠   注意，按字母顺序 "i" 在 "love" 之前。
 *
 *
 * 示例 2：
 *
 *
 * 输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"],
 * k = 4
 * 输出: ["the", "is", "sunny", "day"]
 * 解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
 * ⁠   出现次数依次为 4, 3, 2 和 1 次。
 *
 *
 *
 *
 * 注意：
 *
 *
 * 1 <= words.length <= 500
 * 1 <= words[i] <= 10
 * words[i] 由小写英文字母组成。
 * k 的取值范围是 [1, 不同 words[i] 的数量]
 *
 *
 *
 *
 * 进阶：尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。
 *
 */
import (
	"container/heap"
)

type wordCnt struct {
	Word string
	Cnt  int
}

type Heap []wordCnt

func (h Heap) Len() int { return len(h) }
func (h Heap) Less(p, q int) bool {
	// return h[p].Cnt < h[q].Cnt || h[p].Cnt == h[q].Cnt && h[p].Word > h[q].Word
	if h[p].Cnt != h[q].Cnt {
		return h[p].Cnt < h[q].Cnt
	}
	return h[p].Word > h[q].Word
}
func (h Heap) Swap(p, q int)         { h[p], h[q] = h[q], h[p]; return }
func (h *Heap) Push(v interface{})   { *h = append(*h, v.(wordCnt)); return }
func (h *Heap) Pop() (v interface{}) { l := len(*h); v = (*h)[l-1]; *h = (*h)[:l-1]; return }

func topKFrequent(words []string, k int) []string {
	count := make(map[string]int)
	for _, w := range words {
		count[w]++
	}
	h := &Heap{}
	for w, c := range count {
		heap.Push(h, wordCnt{Word: w, Cnt: c})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	var ans []string
	for h.Len() > 0 {
		wc := heap.Pop(h).(wordCnt)
		ans = append(ans, wc.Word)
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}
