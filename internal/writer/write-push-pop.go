package writer

import (
	"strconv"

	"github.com/sousair/hack-vm-translator/internal/parser"
)

func (w *HackAssemblyWriter) WritePushPop(command parser.CommandType, segment string, index int) {
	w.writeIntoAssemblyFile([]string{"// " + string(command) + " " + segment + " " + strconv.Itoa(index)})
	if command == parser.C_POP && segment == "constant" {
		panic("Cannot pop into constant segment")
	}

	if segment == "pointer" {
		if index > 1 || index < 0 {
			panic("Invalid index for pointer segment")
		}

		if index == 0 {
			segment = "this"
		} else if index == 1 {
			segment = "that"
		}
	}

	memorySegments := map[string]string{
		"local":    "LCL",
		"argument": "ARG",
		"this":     "THIS",
		"that":     "THAT",
		// !TODO: Check index for temp segment
		"temp":   "R" + strconv.Itoa(5+index),
		"static": w.Filename + "." + strconv.Itoa(index),
	}

	if segment != "constant" {
		w.writeIntoAssemblyFile([]string{
			// Calculate the address of the segment[index] and store in A
			"@" + memorySegments[segment],
			"D=M",
			"@" + strconv.Itoa(index),
			"A=D+A",
		})
	} else {
		w.writeIntoAssemblyFile([]string{
			"@" + strconv.Itoa(index),
		})
	}

	switch command {
	case parser.C_PUSH:
		{
			if segment == "constant" {
				w.writeIntoAssemblyFile([]string{
					"D=A",
				})
			} else {
				w.writeIntoAssemblyFile([]string{
					"D=M",
				})
			}

			w.writeIntoAssemblyFile([]string{
				// Push the value of D in the stack
				"@SP",
				"A=M",
				"M=D",
				// Increment the stack pointer
				"@SP",
				"M=M+1",
			})
			break
		}
	case parser.C_POP:
		{
			w.writeIntoAssemblyFile([]string{
				"D=A",
				// Store the address in R13
				"@R13",
				"M=D",
				// Pop the last value and store in D
				"@SP",
				"AM=M-1",
				"D=M",
				// Store the value in segment[index]
				"@R13",
				"A=M",
				"M=D",
			})
			break
		}
	}
}
