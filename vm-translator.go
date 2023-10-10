package main

import (
	"os"
	"unicode"
)

func main() {
	check_args()
}

func check_args() (valid bool) {
	args := os.Args

	if len(args) != 2 {
		panic("Usage: vm-translator <File.vm>")
	}

	filename := args[1]

	if !unicode.IsUpper(rune(filename[0])) {
		panic("File name must start with an uppercase letter")
	}

	valid = true
	return
}
