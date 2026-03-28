// Package sublist checks for sublists.
package sublist

func startswith(a, b []int) bool {
	for i, v := range b {
		if v != a[i] {
			return false
		}
	}
	return true
}

// Sublist returns a Relation: equal, unequal, sublist or superlist.
func Sublist(a, b []int) Relation {
	if len(a) == len(b) && startswith(a, b) {
		return RelationEqual
	}

	// a is the longer list. Check that b is in a.
	swapped := len(a) < len(b)
	r := RelationSuperlist
	if swapped {
		a, b = b, a
		r = RelationSublist
	}

	for i := 0; i <= len(a)-len(b); i++ {
		if startswith(a[i:], b) {
			return r
		}
	}
	return RelationUnequal
}
