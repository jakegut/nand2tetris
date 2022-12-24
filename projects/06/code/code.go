package code

import (
	"fmt"
	"log"
	"strings"
)

func Dest(s string) string {
	dest := 0
	if strings.Contains(s, "M") {
		dest |= 0b001
	}
	if strings.Contains(s, "D") {
		dest |= 0b010
	}
	if strings.Contains(s, "A") {
		dest |= 0b100
	}

	return fmt.Sprintf("%03b", dest)
}

var compAlias = map[string]string{
	"1+D": "D+1",
	"1+M": "M+1",
	"1+A": "A+1",

	"A+D": "D+A",
	"M+D": "D+M",

	"A&D": "D&A",
	"M&D": "D&M",
	"A|D": "D|A",
	"M|D": "D|M",
}

var compMap = map[string]string{
	"0":   "101010",
	"1":   "111111",
	"-1":  "111010",
	"D":   "001100",
	"A":   "110000",
	"!D":  "001101",
	"!A":  "110001",
	"-D":  "001111",
	"-A":  "110011",
	"D+1": "011111",
	"A+1": "110111",
	"D-1": "001110",
	"A-1": "110010",
	"D+A": "000010",
	"D-A": "010011",
	"A-D": "000111",
	"D&A": "000000",
	"D|A": "010101",
}

func Comp(s string) string {
	aBit := "0"
	if idx := strings.IndexByte(s, 'M'); idx != -1 {
		aBit = "1"
		s = s[:idx] + "A" + s[idx+1:]
	}

	if alias, ok := compAlias[s]; ok {
		s = alias
	}

	compBits, ok := compMap[s]
	if !ok {
		log.Fatalf("couldn't find comp: %q", s)
	}

	return aBit + compBits
}

var jumps = []string{"", "JGT", "JEQ", "JGE", "JLT", "JNE", "JLE", "JMP"}

func Jump(s string) string {
	for i, j := range jumps {
		if s == j {
			return fmt.Sprintf("%03b", i)
		}
	}
	return "000"
}
