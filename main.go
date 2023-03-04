package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

func main() {
	args, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(args)

	final := ""
	enter := args
	enter = strings.Split(enter, "/d/")[1]
	if strings.HasSuffix(enter, "/view") {
		enter = strings.Split(enter, "/view")[0]
	} else {
		enter = strings.Split(enter, "/edit")[0]
	}
	final = "https://drive.google.com/uc?id=" + enter
	fmt.Println(final)
	clipboard.WriteAll(final)
}
