package utils

//MapsEqual checks if two string maps are equal
func MapsEqual(m1, m2 map[string]string) bool {
	l1 := 0

	for k1, v1 := range m1 {
		l1++
		v2, e := m2[k1]
		if !e || v1 != v2 {
			return false
		}
	}

	l2 := len(m2)

	if l1 != l2 {
		return false
	}

	return true
}