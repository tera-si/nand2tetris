package main

// This file contains symbol processing logic.

import (
	"fmt"
	"strconv"
)

// Calculates the address that a jump label points to, then adds the label and
// its address to the variable data store.
func processLabel(buf *[]string, in string) {
	labelName := in[1 : len(in)-1] // Unwrap the parenthesises

	_, found := varStore.store[labelName]
	if !found {
		// The label itself is not added to the data store. If current
		// instruction count is n, then the label should points to the nth
		// instruction. (Array index starts with 0)
		adr := strconv.Itoa(len(*buf))
		// Adds label and corresponding address to data store
		varStore.store[labelName] = adr
	}
}

// Takes a symbol and returns an assembly A instruction with its address
// resolved.
// Example
// in = @R1
// Out = @1
func resolveSymbol(in string) string {
	symbol := in[1:] // Removes the "@"

	// Check if symbol is pure number, if it is, it is a memory location and not
	// a symbol, returns it as is.
	if _, err := strconv.Atoi(symbol); err == nil {
		return in
	}

	var adr string

	v, found := varStore.store[symbol]
	if !found {
		// Assign an address to this symbol
		adr = strconv.Itoa(varStore.malloc())
		// Adds the symbol and its address to the data store
		varStore.store[symbol] = adr

		v = adr
	}

	return fmt.Sprintf("@%s", v)
}
