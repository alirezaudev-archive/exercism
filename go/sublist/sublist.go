package sublist

func Sublist(l1, l2 []int) Relation {
	switch {
	case slicesEqual(l1, l2):
		return RelationEqual
	case isSublist(l1, l2):
		return RelationSublist
	case isSublist(l2, l1):
		return RelationSuperlist
	default:
		return RelationUnequal
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isSublist(sub, full []int) bool {
	n := len(sub)
	if n == 0 {
		return true
	}
	for i := 0; i <= len(full)-n; i++ {
		if slicesEqual(full[i:i+n], sub) {
			return true
		}
	}
	return false
}
