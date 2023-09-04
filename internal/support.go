package internal

func ContainsZeroes(s string) bool {
	if len(s) > 1 {
		for i := 0; i < len(s); i++ {
			if s[i] == '0' {
				return true
			} else {
				break
			}
		}
	} else {
		return false
	}
	return false
}
