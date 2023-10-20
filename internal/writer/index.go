package writer

import (
	"os"

	"github.com/sousair/hack-vm-translator/internal/parser"
)

type AssemblyWriter interface {
	WriteArithmetic(command string, comparisonsCount *int)
	WritePushPop(command parser.CommandType, segment string, index int)
	CloseAssemblyFile()
}

type HackAssemblyWriter struct {
	Filename     string
	AssemblyFile os.File
}

func NewHackAssemblyWriter(filename string) AssemblyWriter {
	asmFile, err := os.Create(filename + ".asm")

	if err != nil {
		panic("Cannot create assembly file")
	}

	asmWriter := HackAssemblyWriter{
		AssemblyFile: *asmFile,
		Filename:     filename,
	}

	return &asmWriter
}

func (w *HackAssemblyWriter) CloseAssemblyFile() {
	w.AssemblyFile.Close()
}

func (w *HackAssemblyWriter) writeIntoAssemblyFile(lines []string) {
	for _, line := range lines {
		w.AssemblyFile.WriteString(line + "\n")
	}
}
