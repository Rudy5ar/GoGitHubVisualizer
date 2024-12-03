package main

import (
	"flag"
	"fmt"
)

func scan(folder string) {
	println("Found folders: \n")
	repositories := recursiveScanFolder(folder)
	filepath := getDotFilePath()
	addNewSliceElementsToFile(filepath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")

}

func stats(email string) {
	println("Stats")
}

func recursiveScanFolder(folder string) {
	return
}

func addNewSliceElementsToFile(filepath, repositories) {
	panic("unimplemented")
}

func getDotFilePath() {
	panic("unimplemented")
}

func main() {
	var email string
	var folder string
	flag.StringVar(&email, "email", "your@email.com", "The email to scan")
	flag.StringVar(&folder, "add", "", "Add a new folder to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)

}
