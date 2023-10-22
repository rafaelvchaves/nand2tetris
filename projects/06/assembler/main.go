package main

import (
	"assembler/instruction"
	"assembler/parse"
	"assembler/token"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Expected 1 arg")
	}
	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	destPath := strings.TrimSuffix(os.Args[1], ".asm") + ".hack"
	dest, err := os.Create(destPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	table := map[string]int16{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"SCREEN": 16384,
		"KBD":    24576,
	}

	for i := int16(0); i < 16; i++ {
		table[fmt.Sprintf("R%d", i)] = i
	}

	scanner := bufio.NewScanner(src)
	var tokens []token.Token
	var currInstruction int16

	// First pass: resolve all of the labels.
	for scanner.Scan() {
		t, err := parse.Line(strings.Trim(scanner.Text(), " "))
		if err != nil {
			log.Fatal(err)
		}
		switch t := t.(type) {
		case token.A:
			currInstruction++
		case token.C:
			currInstruction++
		case token.Label:
			table[t.Name] = currInstruction
		}
		tokens = append(tokens, t)
	}

	writer := bufio.NewWriter(dest)

	// Second pass: resolve symbols and generate binary code for instructions.
	for _, insn := range instructions(table, tokens) {
		writer.WriteString(insn.ToBinary() + "\n")
	}
	writer.Flush()
}

func instructions(table map[string]int16, tokens []token.Token) []instruction.Instruction {
	var result []instruction.Instruction
	var nextSymbolAddress int16 = 16
	for _, t := range tokens {
		switch t := t.(type) {
		case token.A:
			if i, err := strconv.Atoi(t.Value); err == nil {
				// If the value is already an address, no lookup needs to be done.
				result = append(result, instruction.A{Value: int16(i)})
				continue
			}
			if _, ok := table[t.Value]; !ok {
				// Allocate a new address if the symbol is not in the table.
				table[t.Value] = nextSymbolAddress
				nextSymbolAddress++
			}
			result = append(result, instruction.A{Value: table[t.Value]})
		case token.C:
			result = append(result, instruction.C{Dest: t.Dest, Comp: t.Comp, Jump: t.Jump})
		}
	}
	return result
}
