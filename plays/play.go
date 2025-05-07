package plays

import (
	"fmt"
	"sort"
)

type kv struct {
	key string
	val int
}

func FindLengthUniqueWords(input []string) int {
	var m = make(map[string]int)
	for _, v := range input {
		m[v]++
	}
	fmt.Println(m)
	fmt.Println(len(m))
	findTop5UniqueWords(m)
	return len(m)
}

func findTop5UniqueWords(input map[string]int) {
	var ss []kv
	for k, v := range input {
		ss = append(ss, kv{k, v})
	}
	// i and j are index
	//when the comparison function returns true, no swap happens.
	sort.Slice(ss, func(i, j int) bool { return ss[i].val > ss[j].val })

	for _, s := range ss[:5] {
		fmt.Println(s.key, " appears ", s.val, " times.")
	}
}
