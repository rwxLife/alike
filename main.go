package main

import (
	"fmt"
	"os"

	"github.com/rwxLife/alike/backup"
)

// usage will print the information about the 'alike' command
func usage() {

	fmt.Println("Usage: alike OPERATION SOURCE DESTINATION")
	fmt.Println("Possible values for, \"OPERATION\": backup, restore")
	fmt.Println("Flags")
	fmt.Println("--help: display this help and exit")
}

func main() {

	argLength := len(os.Args)

	// To deal with typos and/or ignorance
	if argLength == 1 ||
		os.Args[1] == "--help" ||
		argLength != 4 ||
		!(os.Args[1] == "backup" || os.Args[1] == "restore") {
		usage()
		return
	}

	operation := os.Args[1]
	source := os.Args[2]
	destination := os.Args[3]

	if operation == "backup" {
		backup.TraverseAndDoEncryptedBackup(source, destination)
		return
	}

	backup.TraverseAndDecryptBackup(source, destination)
}
