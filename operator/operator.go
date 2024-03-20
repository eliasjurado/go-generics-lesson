package operator

import "golang.org/x/exp/constraints"

func Includes[T comparable](values []T, value T) bool {
	for _, item := range values {
		if item == value {
			return true
		}

	}
	return false
}

func Filter[T constraints.Ordered](values []T, callback func(T) bool) []T {
	result := make([]T, 0, len(values))
	for _, v := range values {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}
