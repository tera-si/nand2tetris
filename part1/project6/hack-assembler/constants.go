package main

// This file defines all the hard-coded data/values

const (
	helpMsg = `Nand2Tetris Hack Assembler
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
`
	markComment      = "//" // Afaik, there aren't multi-line comment in Hack asm
	markAInstruction = "@"
	opcodeA          = "0"

	markDestComp = "="   // Separates the dest and comp in a C instruction
	markCompJmp  = ";"   // Separates the comp and jump in a C instruction
	opcodeC      = "111" // Technically it's just "1" with two unused "11" bits

	// Builtin variables. R0-15 will be populated when assembler initialises
	screenAdr = "16384"
	kbdAdr    = "24576"
	spAdr     = "0"
	lclAdr    = "1"
	argAdr    = "2"
	thisAdr   = "3"
	thatAdr   = "4"

	// Jump bits for C instructions
	jgtBin = "001"
	jeqBin = "010"
	jgeBin = "011"
	jltBin = "100"
	jneBin = "101"
	jleBin = "110"
	jmpBin = "111"

	// Destination index for C instructions
	// For example, if destination is M (001), the 3rd bit should be switched on
	iM = 2
	iD = 1
	iA = 0

	// Computation bits for C instructions
	comp0  = "101010" // 0
	comp1  = "111111" // 1
	comp2  = "111010" // -1
	comp3  = "001100" // D
	comp4  = "110000" // A or M
	comp5  = "001101" // !D
	comp6  = "110001" // !A or !M
	comp7  = "001111" // -D
	comp8  = "110011" // -A or -M
	comp9  = "011111" // D+1
	comp10 = "110111" // A+1 or M+1
	comp11 = "001110" // D-1
	comp12 = "110010" // A-1 or M-1
	comp13 = "000010" // D+A or D+M
	comp14 = "010011" // D-A or D-M
	comp15 = "000111" // A-D or M-D
	comp16 = "000000" // D&A or D&M
	comp17 = "010101" // D|A or D|M
)
