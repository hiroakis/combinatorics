package combinatorics

import (
	"sort"
)

func Combination(list []string, n, r int) [][]string {
	var combis [][]string

	var f func(list, data []string, start, end, index, r int)
	f = func(list, data []string, start, end, index, r int) {
		if index == r {
			s := make([]string, len(data))
			copy(s, data)
			combis = append(combis, s)
			return
		}
		for i := start; i <= end && end-i+1 >= r-index; i++ {
			data[index] = list[i]
			f(list, data, i+1, end, index+1, r)
		}
	}

	data := make([]string, r)
	f(list, data, 0, n-1, 0, r)
	return combis
}

func CombinationWithRepetition(list []string, n, r int) [][]string {
	var combis [][]string

	var f func(list, data []string, start, end, index, r int)
	f = func(list, data []string, start, end, index, r int) {
		if index == r {
			s := make([]string, len(data))
			copy(s, data)
			combis = append(combis, s)
			return
		}
		for i := start; i <= end && end-i+r >= r-index; i++ {
			data[index] = list[i]
			f(list, data, i, end, index+1, r)
		}
	}

	data := make([]string, r)
	f(list, data, 0, n-1, 0, r)
	return combis
}

func Unique(in []string) []string {
	m := map[string]struct{}{}
	for _, v := range in {
		m[v] = struct{}{}
	}
	var l []string
	for k := range m {
		l = append(l, k)
	}
	sort.Strings(l)
	return l
}
