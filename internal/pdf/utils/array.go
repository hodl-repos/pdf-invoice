package utils

func IncludesStr(a []string, s string) bool {
	for _, el := range a {
		if el == s {
			return true
		}
	}
	return false
}

func IncludesInt(a []int, i int) bool {
	for _, el := range a {
		if el == i {
			return true
		}
	}
	return false
}
