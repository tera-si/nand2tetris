package main

// This file contains helper functions that aid the other assembler components.

import (
	"strconv"
	"strings"
)

func removeInlineComment(in string) string {
	for strings.ContainsAny(in, markComment) {
		i := strings.IndexAny(in, markComment)
		in = in[:i]
	}

	return in
}

func getBinaryString(n int) string {
	return strconv.FormatInt(int64(n), 2)
}
