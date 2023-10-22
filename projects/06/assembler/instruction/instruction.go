package instruction

import (
	"fmt"
	"log"
	"strings"
)

type A struct {
	Value int16
}

func (a A) ToBinary() string {
	return fmt.Sprintf("0%015b", a.Value)
}

type C struct {
	Dest string
	Comp string
	Jump string
}

func (c C) ToBinary() string {
	return "111" + aBit(c.Comp) + compBits(c.Comp) + destBits(c.Dest) + jumpBits(c.Jump)
}

func aBit(comp string) string {
	if strings.Contains(comp, "M") {
		return "1"
	}
	return "0"
}

func compBits(comp string) string {
	switch comp {
	case "0":
		return "101010"
	case "1":
		return "111111"
	case "-1":
		return "111010"
	case "D":
		return "001100"
	case "A", "M":
		return "110000"
	case "!D":
		return "001101"
	case "!A", "!M":
		return "110001"
	case "-D":
		return "001111"
	case "-A", "-M":
		return "110011"
	case "D+1":
		return "011111"
	case "A+1", "M+1":
		return "110111"
	case "D-1":
		return "001110"
	case "A-1", "M-1":
		return "110010"
	case "D+A", "D+M":
		return "000010"
	case "D-A", "D-M":
		return "010011"
	case "A-D", "M-D":
		return "000111"
	case "D&A", "D&M":
		return "000000"
	case "D|A", "D|M":
		return "010101"
	}
	log.Fatalf("op code not found for %q", comp)
	return ""
}

func destBits(dest string) string {
	var result string
	if strings.Contains(dest, "A") {
		result += "1"
	} else {
		result += "0"
	}
	if strings.Contains(dest, "D") {
		result += "1"
	} else {
		result += "0"
	}
	if strings.Contains(dest, "M") {
		result += "1"
	} else {
		result += "0"
	}
	return result
}

func jumpBits(jump string) string {
	switch jump {
	case "":
		return "000"
	case "JGT":
		return "001"
	case "JEQ":
		return "010"
	case "JGE":
		return "011"
	case "JLT":
		return "100"
	case "JNE":
		return "101"
	case "JLE":
		return "110"
	case "JMP":
		return "111"
	}
	log.Fatalf("op code not found for %q", jump)
	return ""
}

type Instruction interface {
	ToBinary() string
}
