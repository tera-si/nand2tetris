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

(FOREVER)           // Keep listening for keyboard events until forever
    @i
        M=0         // Loop iteration counter

    @SCREEN
        D=A
    @screen_addr
        M=D         // Base address of screen map

    (LOOP_SCREEN)
        @8192
            D=A         // Store 8192 in data register
        @i
            D=D-M       // Calculate loop iteration left
        @FOREVER
            D;JLE       // 0 iteration left, loop back to FOREVER

        @KBD
            D=M         // Fetch key pressed from map
        @SET_PIXEL_BLACK
            D;JNE       // Key press detected, set next pixel to black
        @SET_PIXEL_WHITE
            D;JEQ       // No key press, set next pixel to white

        (SET_PIXEL_BLACK)
            @i
                D=M     // Get iteration count
            @screen_addr
                A=M+D   // Calculate address of SCREEN[i]
                M=-1    // Set SCREEN[i] to -1, which is black
            @i
                M=M+1   // Increment loop counter
            @LOOP_SCREEN
                0;JMP

        (SET_PIXEL_WHITE)
            @i
                D=M
            @screen_addr
                A=M+D
                M=0
            @i
                M=M+1
            @LOOP_SCREEN
                0;JMP
