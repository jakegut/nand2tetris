package parser

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	input := `(TEST)
@TEST
//skip please
D=D+M //ignore comments on same line
0;JMP
D=M;JGT`

	r := strings.NewReader(input)
	p := New(r)

	if len(p.file) != 6 {
		t.Fatalf("length of file not 6. got=%d", len(p.file))
	}

	tests := []struct {
		expectedSym  string
		expectedDest string
		expectedComp string
		expectedJump string
		expectedIns  InstructionType
	}{
		{
			expectedIns: L_INSTRUCTION,
			expectedSym: "TEST",
		},
		{
			expectedIns: A_INSTRUCTION,
			expectedSym: "TEST",
		},
		{
			expectedIns:  C_INSTRUCTION,
			expectedDest: "D",
			expectedComp: "D+M",
			expectedJump: "",
		},
		{
			expectedIns:  C_INSTRUCTION,
			expectedDest: "",
			expectedComp: "0",
			expectedJump: "JMP",
		},
		{
			expectedIns:  C_INSTRUCTION,
			expectedDest: "D",
			expectedComp: "M",
			expectedJump: "JGT",
		},
	}

	for i, tt := range tests {
		if !p.HasMoreLines() {
			t.Fatalf("expected parser to have more lines at index %d", i)
		}

		p.Advance()

		if p.InstructionType != tt.expectedIns {
			t.Errorf("got wrong instruction type, expected %q, got=%q", tt.expectedIns, p.InstructionType)
		}

		switch tt.expectedIns {
		case A_INSTRUCTION:
		case L_INSTRUCTION:
			if tt.expectedSym != p.Symbol {
				t.Errorf("got wrong symbol, expected %q, got=%q", tt.expectedSym, p.Symbol)
			}
		case C_INSTRUCTION:
			if tt.expectedDest != p.Dest {
				t.Errorf("got wrong dest, expected %q, got=%q", tt.expectedDest, p.Dest)
			}
			if tt.expectedComp != p.Comp {
				t.Errorf("got wrong comp, expected %q, got=%q", tt.expectedComp, p.Comp)
			}
			if tt.expectedJump != p.Jump {
				t.Errorf("got wrong jump, expected %q, got=%q", tt.expectedJump, p.Jump)
			}
		}
	}

	if p.HasMoreLines() {
		t.Fatalf("expected no more lines from parser")
	}

}
