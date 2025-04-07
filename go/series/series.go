package series

func All(n int, s string) []string {
	if n > len(s) || n <= 0 {
		return []string{}
	}

	res := make([]string, 0, len(s)-n+1)
	for i := 0; i <= len(s)-n; i++ {
		res = append(res, s[i:i+n])
	}
	return res
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) || n <= 0 {
		return ""
	}
	return s[:n]
}
