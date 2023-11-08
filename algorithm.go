package algorithm

import (
	"golang.org/x/exp/constraints"
)

func Sort[T constraints.Ordered](s []T) ([]T, error) {

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
	return s, nil
}
