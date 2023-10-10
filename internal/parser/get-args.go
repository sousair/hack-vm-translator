package parser

import "strings"

func (p *HackVMParser) GetArgs(line string) (arg1, arg2 string) {
	words := strings.Split(line, " ")

	if len(words) == 1 {
		return words[0], ""
	}

	return words[1], words[2]
}
