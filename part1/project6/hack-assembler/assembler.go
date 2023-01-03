package main

// This file contains the assembler/translator logic.

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Takes the assembly instructions in bufIn, translates to their machine language
// equivalent and stores them in bufOut, line-by-line
func assemble(bufIn *[]string, bufOut *[]string) {
	for _, s := range *bufIn {
		var translated string

		if strings.HasPrefix(s, markAInstruction) {
			s = resolveSymbol(s)
			translated = translateA(s)
		} else {
			translated = translateC(s)
		}

		*bufOut = append(*bufOut, translated)
	}
}

// Takes a single A instruction and returns its machine language equivalent
func translateA(in string) string {
	address, err := strconv.Atoi(in[1:])
	if err != nil {
		log.Fatalf("[!] Unable to parse A instruction %q: %s", in, err)
	}

	return fmt.Sprintf("%s%015s", opcodeA, getBinaryString(address))
}

// Takes a single C instruction and returns its machine language equivalent
func translateC(in string) string {
	asmC := parseC(in)
	binC := binaryInC{opcode: opcodeC, jump: "000"}

	// Handles destination bits: switch on the corresponding destination bit
	if strings.ContainsAny(asmC.dest, "M") {
		binC.dest[iM] = 1
	}
	if strings.ContainsAny(asmC.dest, "D") {
		binC.dest[iD] = 1
	}
	if strings.ContainsAny(asmC.dest, "A") {
		binC.dest[iA] = 1
	}

	// Handles jump bits
	if len(asmC.jump) != 0 {
		binC.jump = jmpStore[asmC.jump]
	}

	// Handles a bit
	if strings.ContainsAny(asmC.comp, "M") && !strings.ContainsAny(asmC.comp, "A") {
		binC.aOrM = 1
	}

	// Handles computation bits
	for pattern, binary := range compStore {
		if regexp.MustCompile(pattern).MatchString(asmC.comp) {
			binC.comp = binary
			break
		}
	}

	return binC.String()
}
