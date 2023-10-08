// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.

(INIT)
  @color
  M=0       // color=white

(LOOP)
  @24576
  D=M       // D=current key
  @DRAWWHITE
  D;JEQ     // if current key = 0, draw white
  @DRAWBLACK
  0;JMP     // otherwise, draw black

(DRAWWHITE)
  @color
  M=0       // color=white
  @DRAW
  0;JMP     // goto DRAW

(DRAWBLACK)
  @color
  M=-1      // color=black
  @DRAW
  0;JMP     // goto DRAW

(DRAW)
  @SCREEN
  D=A       // D=offset
  @address
  M=D       // address=offset

(DRAWLOOP)
  @address
  D=M
  @24576
  D=D-A     // D=address-24576
  @LOOP
  D;JGE     // if (address-24576)>=0 goto LOOP
  @color
  D=M       // D=color
  @address
  A=M       // A=address
  M=D       // word at address=color
  @address
  M=M+1     // address=address+1
  @DRAWLOOP
  0;JMP     // goto DRAWLOOP
