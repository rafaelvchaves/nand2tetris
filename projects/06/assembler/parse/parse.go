package parse

import (
	"assembler/token"
	"fmt"
	"regexp"
)

const (
	dest    = `(?:(A?M?D?)=)?`
	operand = `(?:[01DAM])`
	op      = `(?:[\-\+&\|!])`
	comp    = `(` + operand + `?` + op + `?` + operand + `)`
	jump    = `(?:;([A-Za-z]+))?`
)

var (
	aPattern     = regexp.MustCompile("^@(.+)")
	cPattern     = regexp.MustCompile(`^` + dest + comp + jump)
	labelPattern = regexp.MustCompile(`^\((.+)\)`)
)

func Line(line string) (token.Token, error) {
	switch {
	case aPattern.MatchString(line):
		fields := aPattern.FindStringSubmatch(line)
		if len(fields) != 2 {
			return nil, fmt.Errorf("expected 2 fields, got %d", len(fields))
		}
		return token.A{Value: fields[1]}, nil
	case cPattern.MatchString(line):
		fields := cPattern.FindStringSubmatch(line)
		if len(fields) != 4 {
			return nil, fmt.Errorf("expected 4 fields, got %d: %v", len(fields), fields)
		}
		return token.C{Dest: fields[1], Comp: fields[2], Jump: fields[3]}, nil
	case labelPattern.MatchString(line):
		fields := labelPattern.FindStringSubmatch(line)
		return token.Label{Name: fields[1]}, nil
	}
	return token.Empty{}, nil
}
