package backup

// TraverseAndDoEncryptedBackups will recursively traverse the source path
// and if it comes across a file during the process, it will encrypt the file
// and store its checksum in the meta file of its root directory.
func TraverseAndDoEncryptedBackup(source string, destination string) {

	if isDirectory(source) {

		dirName := extractName(source)
		makeDirectory(destination, dirName)
		createMetaIfNotFound(destination + dirName)
		contents := getDirectoryListing(source)

		for _, element := range contents {
			TraverseAndDoEncryptedBackup(source+element, destination+dirName)
		}

		return
	}

	cs := getCheckSum(source)
	metaData := newMetaData(cs, source)

	if !metaExists(metaData, destination) {
		addToMeta(metaData, destination)
		encrypted := encryptFile(source)
		writeToDisk(destination, encrypted)
	}
}

// TraverseAndDecryptBackup will recursively traverse the source path
// and if it comes across a file during the process, it will decrypt the file.
func TraverseAndDecryptBackup(source string, destination string) {

	if isDirectory(source) {

		dirName := extractName(source)
		contents := getDirectoryListing(source)

		for _, element := range contents {
			TraverseAndDecryptBackup(source+element, destination+dirName)
		}

		return
	}

	decrypted := decryptFile(source)
	writeToDisk(destination, decrypted)
}
