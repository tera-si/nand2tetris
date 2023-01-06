package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"nand2tetris/vm-translator/constants"
	"nand2tetris/vm-translator/helpers"
	"nand2tetris/vm-translator/translator"
	"os"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, constants.HelpMsg)
		os.Exit(0)
	}

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
	}

	inPath := flag.Arg(0)
	if !strings.HasSuffix(inPath, ".vm") {
		log.Fatalln("[!] Error: expected Hack VM (.vm) file")
	}
	staticLabel := helpers.GetStaticLabel(inPath)

	var bufIn []string
	readFile(&bufIn, inPath)

	var bufOut []string

	tr := translator.Translator{}
	tr.Setup(&bufIn, &bufOut, staticLabel)
	tr.TranslateAll()

	outPath := strings.Split(inPath, ".vm")[0] + ".asm"
	writeFile(&bufOut, outPath)

	log.Printf("[i] Translator output %q successful\n", outPath)
}

func readFile(buf *[]string, filePath string) {
	inFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("[!] Unable to open %q: %s", filePath, err)
	}

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		in := scanner.Text()
		in = strings.TrimSpace(in)
		in = strings.ToUpper(in)
		in = helpers.RemoveInlineComments(in)

		if len(in) == 0 {
			continue
		}

		*buf = append(*buf, in)
	}
}

func writeFile(buf *[]string, filePath string) {
	outFile, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("[!] Error: Unable to create %q: %s", filePath, err)
	}
	defer outFile.Close()

	for _, code := range *buf {
		outFile.WriteString(code + "\n")
	}
}
