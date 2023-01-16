//go:build windows

package permission

// checkCurrentUserRoot on windows is not implemented
func checkCurrentUserRoot() (bool, error) {
	return false, ErrNotImplemented
}

// checkCurrentUserCapNetRaw on windows is not implemented
func checkCurrentUserCapNetRaw() (bool, error) {
	return false, ErrNotImplemented
}
