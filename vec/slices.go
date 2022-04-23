package vec

func UnorderedDelete[T any](s []T, i int) []T {
	last := len(s) - 1
	if i < last {
		s[i], s[last] = s[last], s[i]
	}
	return s[:last]
}
