package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func scan(folder string) {
	println("Found folders: \n")
	repositories := recursiveScanFolder(folder)
	filepath := getDotFilePath()
	addNewSliceElementsToFile(filepath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")

}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func scanGitFolders(folders []string, folder string) []string {
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			println(path)
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}
			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}

func addNewSliceElementsToFile(filepath, newRepos []string) {
	panic("unimplemented")
}

func getDotFilePath() []string {
	panic("unimplemented")
}

func stats(email string) {
	println("Stats")
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
