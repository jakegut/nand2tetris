package parser

import (
	"bufio"
	"io"
	"strings"
)

type InstructionType string

const (
	NONE_INSTRUCTION InstructionType = "NONE"
	A_INSTRUCTION    InstructionType = "A"
	C_INSTRUCTION    InstructionType = "C"
	L_INSTRUCTION    InstructionType = "L"
)

type Parser struct {
	file            []string
	currentIndex    int
	Symbol          string
	Dest            string
	Comp            string
	Jump            string
	InstructionType InstructionType
}

func New(r io.Reader) *Parser {
	file := []string{}

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		s := scanner.Text()
		if idx := strings.Index(s, "//"); idx != -1 {
			s = s[:idx]
		}
		file = append(file, strings.TrimSpace(s))
	}

	return &Parser{
		currentIndex: 0,
		file:         file,
	}
}

func (p *Parser) HasMoreLines() bool {
	return p.currentIndex < len(p.file)
}

func (p *Parser) Advance() {
	p.Symbol = ""
	p.Dest = ""
	p.Comp = ""
	p.Jump = ""
	p.InstructionType = NONE_INSTRUCTION

	if p.HasMoreLines() {
		p.skipLines()

		line := p.file[p.currentIndex]

		p.currentIndex += 1

		switch line[0] {
		case '@':
			p.InstructionType = A_INSTRUCTION
			p.Symbol = line[1:]
		case '(':
			p.InstructionType = L_INSTRUCTION
			p.Symbol = line[1 : len(line)-1]
		default:
			p.InstructionType = C_INSTRUCTION
			l := line
			if idx := strings.Index(l, "="); idx != -1 {
				p.Dest = l[:idx]
				l = l[idx+1:]
			}
			if idx := strings.Index(l, ";"); idx != -1 {
				p.Jump = l[idx+1:]
				l = l[:idx]
			}
			p.Comp = l
		}
	}
}

func (p *Parser) skipLines() {
	for p.HasMoreLines() {
		line := p.file[p.currentIndex]
		if strings.HasPrefix(line, "//") || len(line) == 0 {
			p.currentIndex += 1
		} else {
			break
		}
	}
}

func (p *Parser) Reset() {
	p.currentIndex = 0
}
