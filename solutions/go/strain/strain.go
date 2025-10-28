// Package strain does straining.
package strain

func Keep[T any](i []T, f func(T) bool) []T {
	var ret []T
	for _, v := range i {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func Discard[T any](i []T, f func(T) bool) []T {
	return Keep(i, func(v T) bool { return ! f(v) })
}

