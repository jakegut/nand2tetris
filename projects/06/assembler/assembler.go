package assembler

import (
	"fmt"
	"hack-assembler/code"
	"hack-assembler/parser"
	"hack-assembler/symbols"
	"strconv"
)

type Assembler struct {
	symbols *symbols.Symbols
	parser  *parser.Parser
}

func New(p *parser.Parser) *Assembler {
	return &Assembler{
		symbols: symbols.New(),
		parser:  p,
	}
}

func (a *Assembler) symbolize() {
	currentLine := 1

	for a.parser.HasMoreLines() {
		a.parser.Advance()

		if a.parser.InstructionType == parser.L_INSTRUCTION {
			a.symbols.Set(a.parser.Symbol, currentLine-1)
		} else {
			currentLine += 1
		}
	}

	a.parser.Reset()
}

func (a *Assembler) Assemble() []string {
	a.symbolize()
	currentAddress := 16

	res := []string{}

	for a.parser.HasMoreLines() {
		a.parser.Advance()

		var s string
		switch a.parser.InstructionType {
		case parser.C_INSTRUCTION:
			s = fmt.Sprintf("111%s%s%s", code.Comp(a.parser.Comp), code.Dest(a.parser.Dest), code.Jump(a.parser.Jump))
		case parser.A_INSTRUCTION:
			address, err := strconv.ParseInt(a.parser.Symbol, 10, 64)
			if err != nil {
				if !a.symbols.Contains(a.parser.Symbol) {
					a.symbols.Set(a.parser.Symbol, currentAddress)
					currentAddress += 1
				}
				address = int64(a.symbols.Get(a.parser.Symbol))
			}
			s = fmt.Sprintf("0%015b", address)
		case parser.L_INSTRUCTION:
			fallthrough
		case parser.NONE_INSTRUCTION:
			continue
		}
		res = append(res, s)
	}

	return res
}
