package parser

import (
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	Operator string
	Segment  string
	Dest     int
}

func ParseIn(in string) instruction {
	out := instruction{Dest: -1}

	s := strings.Split(in, " ")

	out.Operator = s[0]

	if len(s) == 3 {
		out.Segment = s[1]

		v, err := strconv.Atoi(s[2])
		if err != nil {
			log.Fatalf("[!] Error: unable to parse instruction %q: %s", in, err)
		}

		out.Dest = v
	}

	return out
}
