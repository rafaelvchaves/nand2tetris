package parse_test

import (
	"assembler/parse"
	"assembler/token"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		want  token.Token
	}{
		{
			input: "@1111",
			want:  token.A{Value: "1111"},
		},
		{
			input: "@R0",
			want:  token.A{Value: "R0"},
		},
		{
			input: "@loop",
			want:  token.A{Value: "loop"},
		},
		{
			input: "@17 //comment",
			want:  token.A{Value: "17"},
		},
		{
			input: "A=-1",
			want:  token.C{Dest: "A", Comp: "-1"},
		},
		{
			input: "M=0",
			want:  token.C{Dest: "M", Comp: "0"},
		},
		{
			input: "AMD=1;JGT",
			want:  token.C{Dest: "AMD", Comp: "1", Jump: "JGT"},
		},
		{
			input: "A=D",
			want:  token.C{Dest: "A", Comp: "D"},
		},
		{
			input: "D=A",
			want:  token.C{Dest: "D", Comp: "A"},
		},
		{
			input: "D=D+M",
			want:  token.C{Dest: "D", Comp: "D+M"},
		},
		{
			input: "D|A;JEQ",
			want:  token.C{Comp: "D|A", Jump: "JEQ"},
		},
		{
			input: "0;JMP",
			want:  token.C{Comp: "0", Jump: "JMP"},
		},
		{
			input: "AD=!D;JLE",
			want:  token.C{Dest: "AD", Comp: "!D", Jump: "JLE"},
		},
		{
			input: "AD=!D;JLE 	// comment ",
			want: token.C{Dest: "AD", Comp: "!D", Jump: "JLE"},
		},
		{
			input: "(LOOP)",
			want:  token.Label{Name: "LOOP"},
		},
		{
			input: "(LOOP) // comment",
			want:  token.Label{Name: "LOOP"},
		},
		{
			input: "(loop)",
			want:  token.Label{Name: "loop"},
		},
		{
			input: "",
			want:  token.Empty{},
		},
		{
			input: "     ",
			want:  token.Empty{},
		},
		{
			input: "// comment",
			want:  token.Empty{},
		},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got, err := parse.Line(test.input)
			if err != nil {
				t.Fatal(err)
			}
			if got != test.want {
				t.Errorf("parse.Line(%q) = %+v, want %+v", test.input, got, test.want)
			}
		})
	}
}
