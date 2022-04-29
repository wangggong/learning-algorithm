/*
 * Hard
 *
 * 给你一个下标从 0 开始的二维整数数组 flowers ，其中 flowers[i] = [starti, endi] 表示第 i 朵花的 花期 从 starti 到 endi （都 包含）。同时给你一个下标从 0 开始大小为 n 的整数数组 persons ，persons[i] 是第 i 个人来看花的时间。
 *
 * 请你返回一个大小为 n 的整数数组 answer ，其中 answer[i]是第 i 个人到达时在花期内花的 数目 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：flowers = [[1,6],[3,7],[9,12],[4,13]], persons = [2,3,7,11]
 * 输出：[1,2,2,2]
 * 解释：上图展示了每朵花的花期时间，和每个人的到达时间。
 * 对每个人，我们返回他们到达时在花期内花的数目。
 * 示例 2：
 *
 *
 *
 * 输入：flowers = [[1,10],[3,3]], persons = [3,3,2]
 * 输出：[2,2,1]
 * 解释：上图展示了每朵花的花期时间，和每个人的到达时间。
 * 对每个人，我们返回他们到达时在花期内花的数目。
 *
 *
 * 提示：
 *
 * 1 <= flowers.length <= 5 * 104
 * flowers[i].length == 2
 * 1 <= starti <= endi <= 109
 * 1 <= persons.length <= 5 * 104
 * 1 <= persons[i] <= 109
 * 通过次数1,755
 * 提交次数4,416
 */

/*
 * var diff []int
 * var prefixSum []int
 *
 * func fullBloomFlowers(flowers [][]int, persons []int) []int {
 * 	var indexes []int
 * 	m := make(map[int]int)
 * 	for _, p := range persons {
 * 		if _, ok := m[p]; !ok {
 * 			indexes = append(indexes, p)
 * 			m[p] = 0
 * 		}
 * 	}
 * 	for _, f := range flowers {
 * 		s, e := f[0], f[1]
 * 		if _, ok := m[s]; !ok {
 * 			indexes = append(indexes, s)
 * 			m[s] = 0
 * 		}
 * 		if _, ok := m[e]; !ok {
 * 			indexes = append(indexes, e)
 * 			m[e] = 0
 * 		}
 *
 * 	}
 * 	sort.Ints(indexes)
 * 	for i, ind := range indexes {
 * 		m[ind] = i
 * 	}
 * 	diff = make([]int, len(indexes)+1)
 * 	for _, f := range flowers {
 * 		s, e := f[0], f[1]
 * 		diff[m[s]]++
 * 		diff[m[e]+1]--
 * 	}
 * 	prefixSum = make([]int, len(diff)+1)
 * 	for i, d := range diff {
 * 		prefixSum[i+1] = prefixSum[i] + d
 * 	}
 * 	// fmt.Println(m, diff, prefixSum)
 * 	var ans []int
 * 	for _, p := range persons {
 * 		ans = append(ans, prefixSum[m[p]+1])
 * 	}
 * 	return ans
 * }
 */

/*
 * // 差分+排序+双指针
 * func fullBloomFlowers(flowers [][]int, persons []int) []int {
 * 	diff := make(map[int]int)
 * 	for _, f := range flowers {
 * 		s, e := f[0], f[1]
 * 		diff[s]++
 * 		diff[e+1]--
 * 	}
 * 	times := make([]int, 0, len(diff))
 * 	for d := range diff {
 * 		times = append(times, d)
 * 	}
 * 	sort.Ints(times)
 * 	var personInfos [][2]int
 * 	for i, p := range persons {
 * 		personInfos = append(personInfos, [2]int{i, p})
 * 	}
 * 	sort.Slice(personInfos, func(p, q int) bool { return personInfos[p][1] < personInfos[q][1] })
 * 	ans := make([]int, len(persons))
 * 	sum := 0
 * 	ind := 0
 * 	for _, pi := range personInfos {
 * 		for ; ind < len(times) && times[ind] <= pi[1]; ind++ {
 * 			sum += diff[times[ind]]
 * 		}
 * 		ans[pi[0]] = sum
 * 	}
 * 	return ans
 * }
 */

// 树状数组 (哈希版), 你猜怎么着? 和数组版一摸一样!
func lowBit(x int) int { return x & -x }

type BIT struct {
	m map[int]int
	n int
}

func (b BIT) query(x int) int {
	ans := 0
	for ; x > 0; x -= lowBit(x) {
		ans += b.m[x]
	}
	return ans
}

func (b *BIT) add(x, v int) {
	for ; x <= b.n; x += lowBit(x) {
		b.m[x] += v
	}
}

func NewBIT(n int) *BIT {
	return &BIT{
		m: make(map[int]int),
		n: n,
	}
}

const maxn int = 1e9

func fullBloomFlowers(flowers [][]int, persons []int) []int {
	b := NewBIT(maxn + 10)
	for _, f := range flowers {
		s, e := f[0], f[1]
		b.add(s, 1)
		b.add(e+1, -1)
	}
	var ans []int
	for _, p := range persons {
		ans = append(ans, b.query(p))
	}
	return ans
}
