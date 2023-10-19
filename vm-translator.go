package main

import (
	"os"
	"unicode"
)

func main() {
	check_args()

	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		panic("Cannot open VM file")
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)
	for lineScanner.Scan() {
		line := strings.TrimSpace(lineScanner.Text())

		println(line)

		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
	}

	if err := lineScanner.Err(); err != nil {
		panic("Error while reading VM file")
	}
}

func check_args() {
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
