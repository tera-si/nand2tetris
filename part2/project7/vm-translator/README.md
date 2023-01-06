# VM Translator (Project 7)

This VM translator is implemented in GO. It takes a file containing Hack VM
bytecode and output its equivalent assembly instructions.

# Run from source

Clone the repository and run `go run .`

# Build from source

Clone the repository and run `go build .`

# Usage
```
Hack VM Translator
Usage:
        vm-translator [-h/--help] BYTECODE

Flags:
        -h/--help            Shows this help message and exits.

Positional Argument:
        BYTECODE             File containing byte code for the Hack virtual machine.
                                                        File is expected to have ".vm" extension. Required.

Description:
        The Hack virtual machine translator reads an Hack virtual machine bytecode
        file (.vm) and outputs its equivalent Hack assembly code (.asm). Both the
        bytecode and assembly code follow the Hack computer architecture and language
        specification defined in the Nand2Tetris courseware.

        This translator is project #7 of the Nand2Tetris (https://www.nand2tetris.org)
        courseware and book "The Elements of Computing Systems" by Noam Nisan and
        Shimon Schocken. This implementation is written in GO by
        tera-si (https://github.com/tera-si).
```
