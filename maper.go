package main

// auth checks access (validate client)
func auth(path string) (access bool) {
	if len(path) < 26 {
		return false
	}

	return true
}
