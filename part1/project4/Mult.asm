// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
//
// This program only needs to handle arguments that satisfy
// R0 >= 0, R1 >= 0, and R0*R1 < 32768.

// --------------------------------------------------------------------
// in0 = RAM[0]
// in1 = RAM[1]
// sum = 0
// for i in range(in1):
//     sum += in0
// RAM[2] = sum
// --------------------------------------------------------------------

// Declarations:
// Fetch value of R0 and store in variable in0
@R0
    D=M
@in0
    M=D

// Fetch value of R1 and store in variable in1
@R1
    D=M
@in1
    M=D

@R2
    M=0 // Prevent garbage value. The test script actually tests for this.

@i
    M=0 // Loop iteration counter
@sum
    M=0
// End of declarations

// Main body:
(ADD_LOOP)
    // Check if need to terminate loop
    @i
        D=M    // Get current iteration
    @in1
        D=M-D  // Check iteration left
    @ADD_LOOP_END
        D;JLE  // End ADD_LOOP if (in1 - i) <= 0

    @in0
        D=M
    @sum
        M=D+M   // sum += in0
    @i
        M=M+1   // Increment loop counter
    @ADD_LOOP
    0;JMP

(ADD_LOOP_END)
    // Fetch computation result and store in target register
    @sum
        D=M
    @R2
        M=D

(PROGRAM_END)
    @PROGRAM_END
    0;JMP
