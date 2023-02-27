package utils

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func SortSlice[T constraints.Ordered](s []T, max bool) {
	sort.Slice(s, func(i, j int) bool {
		if max {
			return s[i] > s[j]
		}
		return s[i] < s[j]
	})
}

func Max[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if m < v {
			m = v
		}
	}

	return m
}

func SliceMin[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := s[0]
	for _, v := range s {
		if v < m {
			m = v
		}
	}

	return m
}

func Min[T constraints.Ordered](x, y T) T {
	if x > y {
		return y
	}
	return x
}

func Sum[T constraints.Integer](s []T) T {
	var zero T
	if len(s) == 0 {
		var zero T
		return zero
	}
	m := zero
	for _, v := range s {
		m += v
	}
	return m
}

func AbsVal[T constraints.Signed](x T) T {
	if x < 0 {
		return -1 * x
	}
	return x
}
