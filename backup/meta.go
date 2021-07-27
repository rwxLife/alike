package backup

import (
	"fmt"
	"os"
)

type meta struct {
	checksum string
	fileName string
}

// addToMeta will add given 'metaData' to the .meta file of the given 'path'
func addToMeta(metaData *meta, path string) {

	// Manage file
	file, err := os.OpenFile(path+".meta", os.O_APPEND|os.O_WRONLY, permissionRWRR)
	handleError(err)
	defer file.Close()

	// Append
	_, err = fmt.Fprintf(file, metaReadWriteFormat, metaData.checksum, metaData.fileName)
	handleError(err)
}

// createMetaData will return a reference to a new meta object
func createMetaData(cs string, name string) *meta {
	return &meta{checksum: cs, fileName: name}
}

// checkMetaExistence will check the .meta file in the given 'path' to see
// if the given metadata exists in the file
func checkMetaExistence(metaData *meta, path string) (bool, bool) {

	// Manage file
	file, err := os.OpenFile(path+".meta", os.O_RDONLY, permissionRWRR)
	handleError(err)
	defer file.Close()

	// Result vars
	var fileExists bool = false
	var checkSumExists bool = false

	// For storing a line's content
	var one string
	var two string

	// Read line-by-line and check
	for true {
		n, err := fmt.Fscanf(file, metaReadWriteFormat, &one, &two)

		if (err != nil && err.Error() == "EOF") || n == 0 {
			break
		}
		handleError(err)

		if one == metaData.checksum {
			checkSumExists = true
		}

		if two == metaData.fileName {
			fileExists = true
		}

	}

	return fileExists, checkSumExists
}

// updateMeta will update given 'metaData' to the .meta file of the given 'path' using fileName as the key
func updateMeta(metaData *meta, path string) {

	reader, err := os.OpenFile(path+".meta", os.O_RDONLY, permissionRWRR)
	handleError(err)

	// For storing a line's content
	var one string
	var two string

	// Store all content in memory
	arr := make([]meta, 0)

	// Read line-by-line and check
	for true {
		n, err := fmt.Fscanf(reader, metaReadWriteFormat, &one, &two)

		if (err != nil && err.Error() == "EOF") || n == 0 {
			break
		}
		handleError(err)

		if two == metaData.fileName {
			one = metaData.checksum
		}
		arr = append(arr, meta{checksum: one, fileName: two})
	}

	reader.Close()

	// Manage writer
	writer, err := os.OpenFile(path+".metacopy", os.O_CREATE|os.O_WRONLY, permissionRWRR)
	handleError(err)
	defer writer.Close()

	// Write all content from memory to disk
	for _, element := range arr {
		_, err = fmt.Fprintf(writer, metaReadWriteFormat, element.checksum, element.fileName)
		handleError(err)
	}

	// Kill the outdated .meta
	removeFile(path + ".meta")

	// Rename the copy
	err = os.Rename(path+".metacopy", path+".meta")
	handleError(err)
}
