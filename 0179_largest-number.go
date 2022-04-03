import (
	// "fmt"
	"strconv"
	"sort"
	"strings"
)

func largestNumber(N []int) string {
	l := len(N)
	S := make([]string, 0, l)
	for _, n := range N {
		S = append(S, strconv.Itoa(n))
	}
	sort.Slice(S, func(p, q int) bool {
		return S[p]+S[q] > S[q]+S[p]
	})
	// fmt.Printf("%v", S)

	s := strings.TrimLeft(strings.Join(S, ""), "0")
	if len(s) == 0 {
		return "0"
	}
	return s
}
