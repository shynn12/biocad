package utilites

func IsInSlice(a string, list []string) bool {
	for _, i := range list {
		if string(i) == a {
			return true
		}
	}
	return false
}
