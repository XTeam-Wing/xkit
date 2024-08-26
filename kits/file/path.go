package file

import "os"

// FileExists checks if the specified file exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
