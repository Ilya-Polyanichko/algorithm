package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	S1 := []int{1, 2, 3, 4, 5}
	S2 := []string{"c", "b", "a"}
	S3 := []float32{4.0, 2.0, 7.0}

	fmt.Println(Sort(S1))
	fmt.Println(Sort(S2))
	fmt.Println(Sort(S3))
}

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
