package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	if len(args) == 0 {
		fmt.Println("Please enter a Google Drive link!")
		os.Exit(1)
	}

	final := ""
	enter := args[0]
	enter = strings.Split(enter, "/d/")[1]

	if strings.HasSuffix(enter, "/view") {
		enter = strings.Split(enter, "/view")[0]
	} else {
		enter = strings.Split(enter, "/edit")[0]
	}

	final = "https://drive.google.com/uc?id=" + enter
	fmt.Println(final)
}
