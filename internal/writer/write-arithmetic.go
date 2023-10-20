package writer

import (
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func (w *HackAssemblyWriter) WriteArithmetic(command string, comparisonsCount *int) {
	w.writeIntoAssemblyFile([]string{"// " + command})

	// TODO: Refactor and use make a function for each command using a map of [string](Function Interface)
	// ? [Maybe] Use some helpers functions (pop last value and store in D, get the address of the second value, etc.)
	unaryOp := []string{"neg", "not"}
	mathOp := []string{"add", "sub", "and", "or"}

	if slices.Contains(unaryOp, command) {
		w.writeUnaryOperation(command)
		return
	}

	w.writeIntoAssemblyFile([]string{
		// Pop the last value and store in D
		"@SP",
		"AM=M-1",
		"D=M",
		// Get the address of the second value (without popping it)
		"A=A-1",
	})

	if slices.Contains(mathOp, command) {
		w.writeMathOperation(command)
		return
	} else {
		w.writeComparisonOperation(command, comparisonsCount)
	}
}

func (w *HackAssemblyWriter) writeUnaryOperation(command string) {
	if command == "neg" {
		w.writeIntoAssemblyFile([]string{
			"@SP",
			"A=M-1",
			"M=-M",
		})
	} else if command == "not" {
		w.writeIntoAssemblyFile([]string{
			"@SP",
			"A=M-1",
			"M=!M",
		})
	}
}

func (w *HackAssemblyWriter) writeMathOperation(command string) {
	if command == "add" {
		w.writeIntoAssemblyFile([]string{"M=M+D"})
	} else if command == "sub" {
		w.writeIntoAssemblyFile([]string{"M=M-D"})
	} else if command == "and" {
		w.writeIntoAssemblyFile([]string{"M=M&D"})
	} else if command == "or" {
		w.writeIntoAssemblyFile([]string{"M=M|D"})
	}
}

func (w *HackAssemblyWriter) writeComparisonOperation(command string, comparisonsCount *int) {
	comparisons := map[string]string{
		"eq": "JEQ",
		"gt": "JGT",
		"lt": "JLT",
	}

	setBoolSuffix := strings.ToUpper(string(command)) + "_" + strconv.Itoa(*comparisonsCount)

	w.writeIntoAssemblyFile([]string{
		"D=M-D",
		"@SET_TRUE_" + setBoolSuffix,
		"D;" + comparisons[command],
		"D=0",
		"@SET_DEFAULT_FALSE_" + setBoolSuffix,
		"0;JMP",
		"(SET_TRUE_" + setBoolSuffix + ")",
		"D=-1",
		"(SET_DEFAULT_FALSE_" + setBoolSuffix + ")",
		"@SP",
		"A=M-1",
		"M=D",
	})

	*comparisonsCount += 1
}
