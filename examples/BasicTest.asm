@256
D=A
@SP
M=D
// C_PUSH constant 10
@10
D=A
@SP
A=M
M=D
@SP
M=M+1
// C_POP local 0
@LCL
D=M
@0
A=D+A
D=A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// C_PUSH constant 21
@21
D=A
@SP
A=M
M=D
@SP
M=M+1
// C_PUSH constant 22
@22
D=A
@SP
A=M
M=D
@SP
M=M+1
// C_POP argument 2
@ARG
D=M
@2
A=D+A
D=A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// C_POP argument 1
@ARG
D=M
@1
A=D+A
D=A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// C_PUSH constant 36
@36
D=A
@SP
A=M
M=D
@SP
M=M+1
// C_POP this 6
@THIS
D=M
@6
A=D+A
D=A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// C_PUSH constant 42
@42
D=A
@SP
A=M
M=D
@SP
M=M+1
// C_PUSH constant 45
@45
D=A
@SP
A=M
M=D
@SP
M=M+1
// C_POP that 5
@THAT
D=M
@5
A=D+A
D=A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// C_POP that 2
@THAT
D=M
@2
A=D+A
D=A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// C_PUSH constant 510
@510
D=A
@SP
A=M
M=D
@SP
M=M+1
// C_POP temp 6
@R5
D=M
@6
A=D+A
D=A
@R13
M=D
@SP
AM=M-1
D=M
@R13
A=M
M=D
// C_PUSH local 0
@LCL
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// C_PUSH that 5
@THAT
D=M
@5
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// C_ARITHMETIC
@SP
AM=M-1
D=M
A=A-1
D=M-D
@SET_TRUE_C_ARITHMETIC_0
D;
@SET_FALSEC_ARITHMETIC_0
0;JMP
(SET_TRUE_C_ARITHMETIC_0)
D=-1
(SET_FALSEC_ARITHMETIC_0)
D=0
@SP
A=M-1
M=D
// C_PUSH argument 1
@ARG
D=M
@1
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// C_ARITHMETIC
@SP
AM=M-1
D=M
A=A-1
D=M-D
@SET_TRUE_C_ARITHMETIC_1
D;
@SET_FALSEC_ARITHMETIC_1
0;JMP
(SET_TRUE_C_ARITHMETIC_1)
D=-1
(SET_FALSEC_ARITHMETIC_1)
D=0
@SP
A=M-1
M=D
// C_PUSH this 6
@THIS
D=M
@6
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// C_PUSH this 6
@THIS
D=M
@6
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// C_ARITHMETIC
@SP
AM=M-1
D=M
A=A-1
D=M-D
@SET_TRUE_C_ARITHMETIC_2
D;
@SET_FALSEC_ARITHMETIC_2
0;JMP
(SET_TRUE_C_ARITHMETIC_2)
D=-1
(SET_FALSEC_ARITHMETIC_2)
D=0
@SP
A=M-1
M=D
// C_ARITHMETIC
@SP
AM=M-1
D=M
A=A-1
D=M-D
@SET_TRUE_C_ARITHMETIC_3
D;
@SET_FALSEC_ARITHMETIC_3
0;JMP
(SET_TRUE_C_ARITHMETIC_3)
D=-1
(SET_FALSEC_ARITHMETIC_3)
D=0
@SP
A=M-1
M=D
// C_PUSH temp 6
@R5
D=M
@6
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
// C_ARITHMETIC
@SP
AM=M-1
D=M
A=A-1
D=M-D
@SET_TRUE_C_ARITHMETIC_4
D;
@SET_FALSEC_ARITHMETIC_4
0;JMP
(SET_TRUE_C_ARITHMETIC_4)
D=-1
(SET_FALSEC_ARITHMETIC_4)
D=0
@SP
A=M-1
M=D
