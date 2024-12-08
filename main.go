package main

import (
	"flag"
)

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
