package main

// This file contains the data models

import "fmt"

// Represents a single assembly C instruction
type asmInC struct {
	dest string
	comp string
	jump string
}

// Represents a single machine language C instruction
type binaryInC struct {
	opcode string // We never need to change this, so it can stay as string
	aOrM   int
	// We'll fetch the corresponding values from the data store and replace the
	// whole string
	comp string
	// Using int array so we can flip individual bits without having to replace
	// the whole (immutable) string
	dest [3]int
	jump string // Lookup from data store
}

// Overrides the ToString method
func (in binaryInC) String() string {
	out := ""

	out += in.opcode
	out += fmt.Sprintf("%d", in.aOrM)
	out += in.comp

	for _, v := range in.dest {
		out += fmt.Sprintf("%d", v)
	}

	out += in.jump

	return out
}

type dataStore map[string]string

// Populate the data store with jump instruction and their binary values
func (store dataStore) populateJMP() {
	store["JGT"] = jgtBin
	store["JEQ"] = jeqBin
	store["JGE"] = jgeBin
	store["JLT"] = jltBin
	store["JNE"] = jneBin
	store["JLE"] = jleBin
	store["JMP"] = jmpBin
}

// Populate the data store with regex patterns that matches the computation part
// of an C instruction, and their binary values
func (store dataStore) populateCompPatterns() {
	// Map key: Regex patterns that greps C instruction computations
	// Map value: Their machine language equivalents
	store[`^0$`] = comp0                  // 0
	store[`^1$`] = comp1                  // 1
	store[`^-1$`] = comp2                 // -1
	store[`^D$`] = comp3                  // D
	store[`^[AM]$`] = comp4               // A or M
	store[`^!D$`] = comp5                 // !D
	store[`^![AM]$`] = comp6              // !A or !M
	store[`^-D$`] = comp7                 // -D
	store[`^-[AM]$`] = comp8              // -A or -M
	store[`^(D\+1|1\+D)$`] = comp9        // D+1 or 1+D
	store[`^([AM]\+1|1\+[AM])$`] = comp10 // A+1, 1+A, M+1, 1+M
	store[`^(D-1|-1\+D)$`] = comp11       // D-1, -1+D
	store[`^([AM]-1|-1\+[AM])$`] = comp12 // A-1, -1+A, M-1, -1+M
	store[`^(D\+[AM]|[AM]\+D)$`] = comp13 // D+A, A+D, D+M, M+D
	store[`^(D-[AM]|-[AM]\+D)$`] = comp14 // D-A, -A+D, D-M, -M+D
	store[`^([AM]-D|-D\+[AM])$`] = comp15 // A-D, -D+A, M-D, -D+M
	store[`^(D\&[AM]|[AM]\&D)$`] = comp16 // D&A, A&D, D&M, M&D
	store[`^(D\|[AM]|[AM]\|D)$`] = comp17 // D|A, A|D, D|M, M|D
}

type symbolStore struct {
	store dataStore
	used  int // Keeps tracks of the number of memory assigned to variables
}

// Populate the data store with builtin variables (e.g. R0, SCREEN)
func (symStore symbolStore) populateBuiltinVars() {
	// Populate R0-15 registers
	for i := 0; i < 16; i++ {
		reg := fmt.Sprintf("R%d", i)
		v := fmt.Sprintf("%d", i)

		symStore.store[reg] = v
	}

	symStore.store["SCREEN"] = screenAdr
	symStore.store["KBD"] = kbdAdr
	symStore.store["SP"] = spAdr
	symStore.store["LCL"] = lclAdr
	symStore.store["ARG"] = argAdr
	symStore.store["THIS"] = thisAdr
	symStore.store["THAT"] = thatAdr
}

// Returns the next free memory address, and increment the used memory count by
// 1. Has to use pointer to refer to the object instance itself, otherwise the
// "used" field will not change.
// btw, I could've call this function "getMemroyAdr" or something like that, but
// let's call it "malloc" for old times sake :)
func (symStore *symbolStore) malloc() int {
	symStore.used++
	return symStore.used
}
