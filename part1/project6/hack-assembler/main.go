package main

// This file contains the main application logic, data store initialisations,
// logger, and file I/O control.

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	varStore  = symbolStore{store: dataStore{}, used: 15} // variable memory starts at 16
	jmpStore  = dataStore{}
	compStore = dataStore{}
)

func init() {
	// initialises data stores
	varStore.populateBuiltinVars()
	jmpStore.populateJMP()
	compStore.populateCompPatterns()
}

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, helpMsg)
		os.Exit(0)
	}

	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "Enables verbosity")
	flag.BoolVar(&verbose, "v", false, "Enables verbosity")

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
	}

	inPath := flag.Arg(0)
	if !strings.HasSuffix(inPath, ".asm") {
		log.Fatalln("[!] Error: Hack assembly file (.asm) expected")
	}
	// Build the output file name
	outPath := strings.Split(inPath, ".asm")[0] + ".hack"

	// Using array allows convenient access to the "instruction - instruction
	// number" mapping, and the "total instructions" count. This array will be
	// processed and all comments will be removed. Once done so, the array
	// element index == instruction number, and array element value == assembly
	// instruction.
	var inLines []string
	if verbose {
		log.Printf("[i] Reading from assembly file %q\n", inPath)
	}
	readInput(&inLines, inPath)

	if verbose {
		log.Printf("[i] %d lines read after ignoring comments\n", len(inLines))
		log.Println("[i] Beginning assembly process")
	}

	var outLines []string
	assemble(&inLines, &outLines)

	if verbose {
		log.Printf("[i] %d/%d lines assembled\n", len(outLines), len(inLines))
		log.Printf("[i] Writing output to %q\n", outPath)
	}
	writeOutput(&outLines, outPath)

	log.Printf("[i] Assembly output to %q successful\n", outPath)
}

// Read the contents of inPath line-by-line and store each line as an array
// element of buffer
func readInput(buf *[]string, inPath string) {
	inFile, err := os.Open(inPath)
	if err != nil {
		log.Fatalf("[!] Error: Unable to open %q: %s", inPath, err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		line = removeInlineComment(line)
		line = strings.TrimSpace(line)

		// Entire line is comment or white space, don't add to buffer
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
			// Line is a jump label. Process it and don't add to instruction
			// list
			processLabel(buf, line)
			continue
		}

		*buf = append(*buf, line)
	}
}

// Write the contents of buf to outPath line-by-line
func writeOutput(buf *[]string, outPath string) {
	outFile, err := os.Create(outPath)
	if err != nil {
		log.Fatalf("[!] Error: Unable to create %q: %s", outPath, err)
	}
	defer outFile.Close()

	for _, code := range *buf {
		outFile.WriteString(code + "\n")
	}
}
