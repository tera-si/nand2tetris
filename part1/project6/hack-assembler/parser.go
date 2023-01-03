package main

// This file contains the C instruction parser logic.

import (
	"strings"
)

// Takes a single C instruction, parses it, and returns an asmInC object.
func parseC(in string) asmInC {
	out := asmInC{"", "", ""}
	var (
		i int // index of "="
		j int // index of ";"
	)

	i = strings.IndexAny(in, markDestComp)
	if i != -1 {
		// Some instructions don't have dest, e.g. `D;JGT`
		out.dest = strings.ToUpper(in[:i]) // Left hand side of "="
	}

	j = strings.IndexAny(in, markCompJmp)
	if j != -1 {
		out.comp = strings.ToUpper(in[i+1 : j]) // Left hand side of ";"
		out.jump = strings.ToUpper(in[j+1:])    // Right hand side of ";"
	} else {
		// No jump bits. Entire part is computation
		out.comp = strings.ToUpper(in[i+1:])
	}

	return out
}
