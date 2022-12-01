package utils

import (
	"golang.org/x/exp/constraints"
)

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
