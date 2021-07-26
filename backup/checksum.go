package backup

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// getCheckSum will get you the CheckSum for the given file
func getCheckSum(path string) string {

	// Manage the file
	file, err := os.Open(path)
	handleError(err)
	defer file.Close()

	// Setup the hasher
	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	handleError(err)

	// Hash it!
	cs := hasher.Sum(nil)

	return hex.EncodeToString(cs)
}
