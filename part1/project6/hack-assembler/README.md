# Hack Assembler (Project 6)

This assembler is implemented in GO. It takes a file containing Hack assembly
code and outputs its machine language equivalent.

# Run from Source

To run it from source, clone the project and use `go run .`

# Build from Source

To build the project, clone it and use `go build .`

# Usage
```
Nand2Tetris Hack Assembler
Usage:
        hack-assembler [-h/--help] [-v/--verbose] ASSEMBLY

Flags:
        -h/--help           Shows this help message and exits.
        -v/--verbose        Enables verbosity. (Default: off)

Positional Argument:
        ASSEMBLY            File containing Hack assembly code. File is expected to
                            have ".asm" extension. Required.

Description:
        The Hack computer assembler reads an Hack assembly (.asm) file and output its
        Hack machine language instructions (.hack). Both the assembly and machine codes
        follow the Hack computer and language specification defined in the Nand2Tetris
        courseware.

        This assembler is project #6 of the Nand2Tetris (https://www.nand2tetris.org)
        courseware and book "The Elements of Computing Systems" by Noam Nisan and
        Shimon Schocken. This implementation is written in GO by
        tera-si (https://github.com/tera-si).
```
