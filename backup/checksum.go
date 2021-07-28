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
	handleError("getCheckSum: file I/O ", err)
	defer file.Close()

	// Setup the hasher
	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	handleError("getCheckSum: hashing ", err)

	// Hash it!
	cs := hasher.Sum(nil)

	return hex.EncodeToString(cs)
}
