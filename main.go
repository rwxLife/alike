package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"syscall"

	"github.com/rwxLife/alike/backup"
	"golang.org/x/term"
)

const (
	asciiLogo = `
     ###    ##       #### ##    ## ########
    ## ##   ##        ##  ##   ##  ##
   ##   ##  ##        ##  ##  ##   ##
  ##     ## ##        ##  #####    ######
  ######### ##        ##  ##  ##   ##
  ##     ## ##        ##  ##   ##  ##
  ##     ## ######## #### ##    ## ########`
	colorReset = "\033[0m"
	colorCyan  = "\033[36m"
)

// usage will print the information about the 'alike' command
func usage() {

	fmt.Println("Usage: alike OPERATION SOURCE DESTINATION")
	fmt.Println("Possible values for, \"OPERATION\": backup, restore")
	fmt.Println("Flags")
	fmt.Println("--help: display this help and exit")
}

// acquirePassword will prompt the user to enter a password
func acquirePassword() {

	fmt.Print("Enter password: ")
	binput, err := term.ReadPassword(int(syscall.Stdin))
	input := string(binput)

	if err != nil {
		log.Fatal(err)
	}

	if runtime.GOOS != "windows" {
		input = strings.Replace(input, "\n", "", -1)
	} else {
		input = strings.Replace(input, "\r\n", "", -1)
	}

	os.Setenv("ALIKE_PASSWORD", input)
	fmt.Println()
}

func printColor(colorName string) {

	os := runtime.GOOS
	if os != "windows" {
		fmt.Print(colorName)
	}
}

func main() {

	printColor(colorCyan)
	fmt.Println(asciiLogo)
	printColor(colorReset)

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

	acquirePassword()
	backup.CalculateCredentials()

	if operation == "backup" {
		backup.TraverseAndDoEncryptedBackup(source, destination)
		return
	}

	backup.TraverseAndDecryptBackup(source, destination)
}
