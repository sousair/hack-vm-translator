package parser

type CommandType string

const (
	C_ARITHMETIC CommandType = "C_ARITHMETIC"
	C_PUSH       CommandType = "C_PUSH"
	C_POP        CommandType = "C_POP"
	C_LABEL      CommandType = "C_LABEL"
	C_GOTO       CommandType = "C_GOTO"
	C_IF         CommandType = "C_IF"
	C_FUNCTION   CommandType = "C_FUNCTION"
	C_RETURN     CommandType = "C_RETURN"
	C_CALL       CommandType = "C_CALL"
)

type VMParser interface {
	GetCommandType(line string) CommandType
	GetArgs(line string) (arg1, arg2 string)
}

type HackVMParser struct{}

func NewHackVMParser() *HackVMParser {
	return &HackVMParser{}
}
