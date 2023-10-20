package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/sousair/hack-vm-translator/internal/parser"
	"github.com/sousair/hack-vm-translator/internal/writer"
)

func main() {
	checkArgs()

	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		panic("Cannot open VM file")
	}
	defer file.Close()

	lineScanner := bufio.NewScanner(file)

	vmParser := parser.NewHackVMParser()

	asmFileName := filename[:strings.LastIndex(filename, ".")]
	vmWriter := writer.NewHackAssemblyWriter(asmFileName)

	comparisonCount := 0
	for lineScanner.Scan() {
		line := strings.TrimSpace(lineScanner.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		commandType := vmParser.GetCommandType(line)
		arg1, arg2 := vmParser.GetArgs(line)

		if arg2 != "" {
			index, err := strconv.Atoi(arg2)
			if err != nil {
				panic("Error while parsing push/pop index")
			}

			vmWriter.WritePushPop(commandType, arg1, index)
			continue
		} else {
			vmWriter.WriteArithmetic(arg1, &comparisonCount)
		}

	}

	if err := lineScanner.Err(); err != nil {
		panic("Error while reading VM file")
	}
}

func checkArgs() {
	args := os.Args
	if len(args) != 2 {
		panic("Usage: vm-translator <File.vm>")
	}

	filename := args[1]

	if !unicode.IsUpper(rune(filename[0])) {
		panic("File name must start with an uppercase letter")
	}

	if !strings.HasSuffix(filename, ".vm") {
		panic("File extension must be .vm")
	}
}
