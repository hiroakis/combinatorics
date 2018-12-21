package combinatorics

import (
	"reflect"
	"testing"
)

func TestCombination(t *testing.T) {
	tests := []struct {
		name   string
		list   []string
		n      int
		r      int
		expect [][]string
	}{
		{
			name: "3C2",
			list: []string{"a", "b", "c"},
			n:    3,
			r:    2,
			expect: [][]string{
				[]string{"a", "b"},
				[]string{"a", "c"},
				[]string{"b", "c"},
			},
		},
		{
			name: "3C3",
			list: []string{"a", "b", "c"},
			n:    3,
			r:    3,
			expect: [][]string{
				[]string{"a", "b", "c"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retval := Combination(tt.list, tt.n, tt.r)
			if !reflect.DeepEqual(retval, tt.expect) {
				t.Errorf(`got(%v) != expect(%v)`, retval, tt.expect)
			}
		})
	}
}

func TestCombinationWithRepetition(t *testing.T) {
	tests := []struct {
		name   string
		list   []string
		n      int
		r      int
		expect [][]string
	}{
		{
			name: "3C2",
			list: []string{"a", "b", "c"},
			n:    3,
			r:    2,
			expect: [][]string{
				[]string{"a", "a"},
				[]string{"a", "b"},
				[]string{"a", "c"},
				[]string{"b", "b"},
				[]string{"b", "c"},
				[]string{"c", "c"},
			},
		},
		{
			name: "3C3",
			list: []string{"a", "b", "c"},
			n:    3,
			r:    3,
			expect: [][]string{
				[]string{"a", "a", "a"},
				[]string{"a", "a", "b"},
				[]string{"a", "a", "c"},
				[]string{"a", "b", "b"},
				[]string{"a", "b", "c"},
				[]string{"a", "c", "c"},
				[]string{"b", "b", "b"},
				[]string{"b", "b", "c"},
				[]string{"b", "c", "c"},
				[]string{"c", "c", "c"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retval := CombinationWithRepetition(tt.list, tt.n, tt.r)
			if !reflect.DeepEqual(retval, tt.expect) {
				t.Errorf(`got(%v) != expect(%v)`, retval, tt.expect)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name   string
		in     []string
		expect []string
	}{
		{
			name:   "a should be removed",
			in:     []string{"a", "a", "a"},
			expect: []string{"a"},
		},
		{
			name:   "b should be removed",
			in:     []string{"a", "b", "c", "b"},
			expect: []string{"a", "b", "c"},
		},
		{
			name:   "already unique",
			in:     []string{"a", "b", "c"},
			expect: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retval := Unique(tt.in)
			if !reflect.DeepEqual(retval, tt.expect) {
				t.Errorf(`got(%v) != expect(%v)`, retval, tt.expect)
			}
		})
	}
}
