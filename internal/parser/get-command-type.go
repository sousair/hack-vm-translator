package parser

import "strings"

func (p *HackVMParser) GetCommandType(line string) CommandType {
	firstWord := strings.Split(line, " ")[0]

	switch firstWord {
	case "push":
		return C_PUSH
	case "pop":
		return C_POP
	case "label":
		return C_LABEL
	case "goto":
		return C_GOTO
	case "if-goto":
		return C_IF
	case "function":
		return C_FUNCTION
	case "return":
		return C_RETURN
	case "call":
		return C_CALL
	default:
		return C_ARITHMETIC
	}
}
