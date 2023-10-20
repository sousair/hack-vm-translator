@256
D=A
@SP
M=D
// C_PUSH constant 20
@20
D=A
@SP
A=M
M=D
@SP
M=M+1
// C_PUSH constant 10
@10
D=A
@SP
A=M
M=D
@SP
M=M+1
// add
@SP
AM=M-1
D=M
A=A-1
M=M+D
// C_PUSH constant 30
@30
D=A
@SP
A=M
M=D
@SP
M=M+1
// eq
@SP
AM=M-1
D=M
A=A-1
D=M-D
@SET_TRUE_EQ_0
D;JEQ
D=0
@SET_DEFAULT_FALSE_EQ_0
0;JMP
(SET_TRUE_EQ_0)
D=-1
(SET_DEFAULT_FALSE_EQ_0)
@SP
A=M-1
M=D
// not
@SP
A=M-1
M=!M
