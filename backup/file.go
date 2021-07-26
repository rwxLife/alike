package backup

import (
	"os"
)

type meta struct {
	fileName string
	checksum string
}

// extractName will get you the name of the file or the directory from the given
// variable 'path'
func extractName(path string) string {

	arr := make([]rune, 0)
	length := len(path)

	for i := length - 1; i >= 0; i-- {

		if path[i] == '/' && i != length-1 {
			break
		}

		arr = append(arr, rune(path[i]))
	}

	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}

	return string(arr)
}

// isDirectory will tell you if the given 'path' is a directory
func isDirectory(path string) bool {

	info, err := os.Stat(path)
	handleError(err)
	return info.IsDir()
}

// createMeta will create a .meta file in the root of the given 'path'
// if not found
func createMetaIfNotFound(path string) {

	slash := ""
	if path[len(path)-1] != '/' {
		slash = "/"
	}
	filePath := path + slash + ".meta"

	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {

		file, err := os.Create(filePath)
		handleError(err)
		file.Close()
	}
}

// makeDirectory will make a 'name' directory under the given 'path'
func makeDirectory(path string, name string) {

	slash := ""
	if path[len(path)-1] != '/' {
		slash = "/"
	}
	dirPath := path + slash + name
	err := os.Mkdir(dirPath, 0755)
	handleError(err)
}

// getDirectoryListing is basically your 'ls' command for the given directory
func getDirectoryListing(path string) []string {
	entries, err := os.ReadDir(path)
	handleError(err)
	var list []string = make([]string, 0)
	for _, entry := range entries {
		list = append(list, entry.Name())
	}
	return list
}