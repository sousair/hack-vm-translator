package writer

import (
	"os"

	"github.com/sousair/hack-vm-translator/internal/parser"
)

type AssemblyWriter interface {
	WriteArithmetic(command string, comparisonsCount *int)
	WritePushPop(command parser.CommandType, segment string, index int)
}

type HackAssemblyWriter struct {
	AssemblyFile os.File
}

func NewHackAssemblyWriter(filename string) AssemblyWriter {
	asmFile, err := os.Create(filename + ".asm")

	if err != nil {
		panic("Cannot create assembly file")
	}

	asmWriter := HackAssemblyWriter{
		AssemblyFile: *asmFile,
	}

	// Init SP to first stack address (256)
	asmWriter.writeIntoAssemblyFile([]string{
		"@256",
		"D=A",
		"@SP",
		"M=D",
	})

	return &asmWriter
}

func (w *HackAssemblyWriter) writeIntoAssemblyFile(lines []string) {
	for _, line := range lines {
		w.AssemblyFile.WriteString(line + "\n")
	}
}
