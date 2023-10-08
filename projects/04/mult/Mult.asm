// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)
//
// This program only needs to handle arguments that satisfy
// R0 >= 0, R1 >= 0, and R0*R1 < 32768.

// Suppose a=RAM[R0],b=RAM[R1],sum=RAM[R2]
// Pseudocode:
// i=0, sum=0
// for i < b {
//   sum = sum + a  
// }

(INIT)
  @i
  M=0     // i=0
  @R2
  M=0     // sum=0

(LOOP)
  @i
  D=M     // D=i
  @R1
  D=D-M   // D=i-b
  @END
  D;JGE   // if (i-b)>=0 goto END
  @R0
  D=M     // D=a
  @R2
  D=M+D   // D=sum+a
  M=D     // sum=sum+a
  @i
  M=M+1   // i=i+1

  @LOOP
  0;JMP   // goto LOOP

(END)
  @END
  0;JMP   // infinite loop
