package sublist

// Relation type is defined in relations.go file.

func equal(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func contains(l1, l2 []int) bool {
	len1 := len(l1)
	len2 := len(l2)
	if len2 == 0 {
		return true
	}
	if len1 < len2 {
		return false
	}
	if len1 == len2 {
		return equal(l1, l2)
	}
	for i1 := 0; i1+len2 <= len1; i1++ {
		if equal(l1[i1:i1+len2], l2) {
			return true
		}
	}
	return false
}

func Sublist(l1, l2 []int) Relation {
	len1 := len(l1)
	len2 := len(l2)
	switch {
	case len1 == len2:
		if equal(l1, l2) {
			return RelationEqual
		}
		return RelationUnequal
	case len1 > len2:
		if contains(l1, l2) {
			return RelationSuperlist
		}
		return RelationUnequal
	default: // len1 < len2
		if contains(l2, l1) {
			return RelationSublist
		}
		return RelationUnequal
	}
}
