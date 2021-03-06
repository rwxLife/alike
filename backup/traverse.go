package backup

// TraverseAndDoEncryptedBackup will recursively traverse the source path
// and if it comes across a file during the process, it will encrypt the file
// and store its checksum in the meta file of its root directory.
func TraverseAndDoEncryptedBackup(source string, destination string) {

	if isDirectory(source) {

		source = appendSeparator(source)
		dirName := extractName(source)
		setupDirectory(destination, dirName)
		contents := getDirectoryListing(source)

		for _, element := range contents {
			TraverseAndDoEncryptedBackup(source+element, destination+dirName)
		}

		return
	}

	createMetaIfNotFound(destination)
	encryptAndBackup(source, destination)
}

// TraverseAndDecryptBackup will recursively traverse the source path
// and if it comes across a file during the process, it will decrypt the file.
func TraverseAndDecryptBackup(source string, destination string) {

	if isDirectory(source) {

		source = appendSeparator(source)
		dirName := extractName(source)
		makeDirectory(destination, dirName)
		contents := getDirectoryListing(source)

		for _, element := range contents {
			TraverseAndDecryptBackup(source+element, destination+dirName)
		}

		return
	}

	createMetaIfNotFound(destination)
	decryptAndRestore(source, destination)
}

// appendSeparator will add '/' to the end of the 'source' if it does not exist
func appendSeparator(source string) string {

	if len(source) != 0 && source[len(source)-1] != '/' {
		source += "/"
	}
	return source
}

// setupDirectory will do create the directory along with its
// .meta file
func setupDirectory(destination string, dirName string) {

	makeDirectory(destination, dirName)
	createMetaIfNotFound(destination + dirName)
}

// encryptAndBackup will encrypt the given file and write to disk
func encryptAndBackup(source string, destination string) {

	fileName := extractName(source)
	cs := getCheckSum(source)
	metaData := createMetaData(cs, fileName)
	fileExists, checkSumExists := checkMetaExistence(metaData, destination)

	if fileExists && checkSumExists {
		return
	}

	if fileExists && !checkSumExists {
		updateMeta(metaData, destination)
		removeFile(destination + fileName)
	}

	if !fileExists {
		addToMeta(metaData, destination)
	}

	encrypted := encryptFile(source)
	writeToDisk(destination+fileName, encrypted)
}

// decryptAndRestore will decrypt the given file and write to disk
func decryptAndRestore(source string, destination string) {

	fileName := extractName(source)
	if fileName != ".meta" {
		decrypted := decryptFile(source)
		writeToDisk(destination+fileName, decrypted)
	}
}
